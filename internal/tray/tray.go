package tray

import (
	"context"
	"embed"
	"fmt"
	"io"
	"log"
	"runtime"

	"fyne.io/systray"
	"github.com/sonjek/mouse-stay-up/internal/config"
	"github.com/sonjek/mouse-stay-up/internal/keyboard"
	"github.com/sonjek/mouse-stay-up/internal/mouse"
	"github.com/sonjek/mouse-stay-up/internal/utils"
)

//go:embed icon.svg
var iconFile embed.FS

const tooltipText = "Enable or Disable periodic mouse movements"

func loadIcon() ([]byte, error) {
	file, err := iconFile.Open("icon.svg")
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	return io.ReadAll(file)
}

// Tray manages the system tray icon and menu.
type Tray struct {
	ctx    context.Context
	cancel context.CancelFunc

	mouseController    *mouse.Controller
	keyboardController *keyboard.Controller
	conf               *config.Config

	// Menu items — grouped in display order.
	mEnable         *systray.MenuItem
	mDisable        *systray.MenuItem
	mWorkingHours   *systray.MenuItem
	mLockKeyboard   *systray.MenuItem
	mUnlockKeyboard *systray.MenuItem
	mAbout          *systray.MenuItem
	mQuit           *systray.MenuItem

	workingHoursMenuItems map[string]*systray.MenuItem
}

// NewTray creates a new Tray instance.
func NewTray(
	ctx context.Context,
	mouseController *mouse.Controller,
	keyboardController *keyboard.Controller,
	conf *config.Config,
) *Tray {
	return &Tray{
		ctx:                   ctx,
		mouseController:       mouseController,
		keyboardController:    keyboardController,
		conf:                  conf,
		workingHoursMenuItems: make(map[string]*systray.MenuItem),
	}
}

// Run starts the system tray. It blocks until the user quits.
func (t *Tray) Run() {
	t.ctx, t.cancel = context.WithCancel(t.ctx)
	defer t.cancel()
	systray.Run(t.onReady, t.onExit)
}

func (t *Tray) onReady() {
	if err := t.setupIcon(); err != nil {
		log.Printf("tray: %v", err)
		systray.Quit()
		return
	}
	t.setupMenu()
	t.applyInitialState()
	go t.runEventLoop()
	go t.mouseController.MoveMouse(t.ctx)
}

func (t *Tray) setupIcon() error {
	icon, err := loadIcon()
	if err != nil {
		return fmt.Errorf("failed to load icon: %w", err)
	}
	systray.SetTemplateIcon(icon, icon)
	systray.SetTooltip(tooltipText)
	return nil
}

// setupMenu creates all menu items in display order.
func (t *Tray) setupMenu() {
	t.mEnable = systray.AddMenuItem("Enable movement", "Enable mouse movement")
	t.mDisable = systray.AddMenuItem("Disable movement", "Disable mouse movement")
	t.mWorkingHours = systray.AddMenuItem("Working hours", "Select a range of working hours")
	for _, hours := range t.conf.WorkingHours {
		t.workingHoursMenuItems[hours] = t.mWorkingHours.AddSubMenuItem(hours, hours)
	}

	systray.AddSeparator()

	t.mLockKeyboard = systray.AddMenuItem("Lock keyboard", "Lock keyboard")
	t.mUnlockKeyboard = systray.AddMenuItem("Unlock keyboard", "Unlock keyboard")

	systray.AddSeparator()

	t.mAbout = systray.AddMenuItem("About", "Open GitHub repo")
	t.mQuit = systray.AddMenuItem("Quit", "Quit the application")
}

// applyInitialState sets menu visibility and checks based on the persisted config.
func (t *Tray) applyInitialState() {
	t.updateMovementMenuState()
	t.updateKeyboardMenuState()

	if item, ok := t.workingHoursMenuItems[t.conf.WorkingHoursInterval]; ok {
		item.Check()
	}
}

// runEventLoop handles all menu click events. Must be called in a goroutine.
func (t *Tray) runEventLoop() {
	workingHoursClicks := t.mergeWorkingHoursClicks()

	for {
		select {
		case <-t.ctx.Done():
			return

		case <-t.mEnable.ClickedCh:
			t.conf.ToggleEnableDisable()
			t.updateMovementMenuState()

		case <-t.mDisable.ClickedCh:
			t.conf.ToggleEnableDisable()
			t.updateMovementMenuState()

		case interval := <-workingHoursClicks:
			t.conf.SetWorkingHoursInterval(interval)
			t.updateWorkingHoursChecks(interval)

		case <-t.mLockKeyboard.ClickedCh:
			t.keyboardController.LockKeyboard()
			t.updateKeyboardMenuState()

		case <-t.mUnlockKeyboard.ClickedCh:
			t.keyboardController.UnlockKeyboard()
			t.updateKeyboardMenuState()

		case <-t.mAbout.ClickedCh:
			if err := utils.OpenWebPage(config.GitRepo); err != nil {
				log.Printf("tray: failed to open GitHub: %v", err)
			}

		case <-t.mQuit.ClickedCh:
			systray.Quit()
			return
		}
	}
}

// mergeWorkingHoursClicks fans-in all working-hours sub-menu clicks into a
// single channel. Each goroutine is tied to t.ctx and stops on cancellation.
func (t *Tray) mergeWorkingHoursClicks() <-chan string {
	clicks := make(chan string)
	for interval, item := range t.workingHoursMenuItems {
		go func(interval string, item *systray.MenuItem) {
			for {
				select {
				case <-t.ctx.Done():
					return
				case <-item.ClickedCh:
					select {
					case <-t.ctx.Done():
						return
					case clicks <- interval:
					}
				}
			}
		}(interval, item)
	}
	return clicks
}

// updateWorkingHoursChecks sets the checkmark on the selected interval and
// clears all others.
func (t *Tray) updateWorkingHoursChecks(selected string) {
	for interval, item := range t.workingHoursMenuItems {
		if interval == selected {
			item.Check()
		} else {
			item.Uncheck()
		}
	}
}

// updateMovementMenuState synchronises the Enable/Disable menu item visibility
// and the Working Hours menu state with the current config.
func (t *Tray) updateMovementMenuState() {
	if t.conf.Enabled {
		t.mEnable.Hide()
		t.mDisable.Show()
		t.mWorkingHours.Enable()
	} else {
		t.mEnable.Show()
		t.mDisable.Hide()
		t.mWorkingHours.Disable()
	}
}

// updateKeyboardMenuState synchronises the Lock/Unlock keyboard menu items
// with the controller state. On non-darwin platforms both items are hidden.
func (t *Tray) updateKeyboardMenuState() {
	if runtime.GOOS != "darwin" {
		t.mLockKeyboard.Hide()
		t.mUnlockKeyboard.Hide()
		return
	}

	if t.keyboardController.KeyboardLocked {
		t.mLockKeyboard.Hide()
		t.mUnlockKeyboard.Show()
	} else {
		t.mLockKeyboard.Show()
		t.mUnlockKeyboard.Hide()
	}
}

func (t *Tray) onExit() {
	if t.cancel != nil {
		t.cancel()
	}
	if t.keyboardController.KeyboardLocked {
		t.keyboardController.UnlockKeyboard()
	}
}
