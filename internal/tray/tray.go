package tray

import (
	"embed"
	"fmt"
	"io"
	"runtime"

	"fyne.io/systray"
	"github.com/sonjek/mouse-stay-up/internal/config"
	"github.com/sonjek/mouse-stay-up/internal/keyboard"
	"github.com/sonjek/mouse-stay-up/internal/mouse"
	"github.com/sonjek/mouse-stay-up/internal/utils"
)

//go:embed icon.png
var iconFile embed.FS

func loadIcon() ([]byte, error) {
	file, err := iconFile.Open("icon.png")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type Tray struct {
	mouseController       *mouse.Controller
	keyboardController    *keyboard.Controller
	conf                  *config.Config
	mEnable               *systray.MenuItem
	mDisable              *systray.MenuItem
	mWorkingHours         *systray.MenuItem
	kUnlockKeyboard       *systray.MenuItem
	kLockKeyboard         *systray.MenuItem
	workingHoursMenuItems map[string]*systray.MenuItem
}

func NewTray(mouseController *mouse.Controller, keyboardController *keyboard.Controller, conf *config.Config) *Tray {
	return &Tray{
		mouseController:       mouseController,
		keyboardController:    keyboardController,
		conf:                  conf,
		workingHoursMenuItems: make(map[string]*systray.MenuItem),
	}
}

func (t *Tray) Run() {
	systray.Run(t.onReady, t.onExit)
}

func (t *Tray) onReady() {
	icon, err := loadIcon()
	if err != nil {
		panic(fmt.Errorf("could not load icon due to error `%s`", err.Error()))
	}

	// Set icon and tooltip for the systray icon
	systray.SetTemplateIcon(icon, icon)
	systray.SetTooltip("Enable or Disable periodic mouse movements")

	// Create menu items for enable/disable mouse movement, change working hours and exit
	t.mEnable = systray.AddMenuItem("Enable movement", "Enable mouse movement")
	t.mDisable = systray.AddMenuItem("Disable movement", "Disable mouse movement")
	t.mWorkingHours = systray.AddMenuItem("Working hours", "Select a range of working hours")
	systray.AddSeparator()
	t.kUnlockKeyboard = systray.AddMenuItem("Unlock keyboard", "Unlock keyboard")
	t.kLockKeyboard = systray.AddMenuItem("Lock keyboard", "Lock keyboard")
	systray.AddSeparator()
	mAbout := systray.AddMenuItem("About", "Open GitHub repo")
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	// Add interval selection submenu items
	for _, hours := range t.conf.WorkingHours {
		t.addWorkingHoursItems(t.mWorkingHours, hours)
	}

	// Adjust visibilities based on the app state
	t.applyEnableDisable()
	t.initLockKeyboard()

	// Set a marker for the default working hours interval
	t.workingHoursMenuItems[t.conf.WorkingHoursInterval].Check()

	workingHoursIntervalClicks := t.createWorkingHoursIntervalClicksChannel()

	go func() {
		for {
			select {
			case <-t.mEnable.ClickedCh:
				t.conf.ToggleEnableDisable()
				t.applyEnableDisable()
			case <-t.mDisable.ClickedCh:
				t.conf.ToggleEnableDisable()
				t.applyEnableDisable()
			case workingHoursInterval := <-workingHoursIntervalClicks:
				// When an hours interval item is clicked, update the workingHoursInterval interval and checkmarks
				t.conf.SetWorkingHoursInterval(workingHoursInterval)
				t.updateNightModeIntervalChecks(t.conf.WorkingHoursInterval)
			case <-t.kLockKeyboard.ClickedCh:
				t.applyEnableDisableKeyboard()
				t.keyboardController.LockKeyboard()
			case <-t.kUnlockKeyboard.ClickedCh:
				t.applyEnableDisableKeyboard()
				t.keyboardController.UnlockKeyboard()
			case <-mAbout.ClickedCh:
				if err := utils.OpenWebPage(config.GitRepo); err != nil {
					panic(err)
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

	// Start moving the mouse if app enabled
	go t.mouseController.MoveMouse()
}

// Adds a submenu item for selecting a working hours interval
func (t *Tray) addWorkingHoursItems(parent *systray.MenuItem, interval string) {
	t.workingHoursMenuItems[interval] = parent.AddSubMenuItem(interval, interval)
}

// Creates and returns a channel that listens to all working hours interval item clicks
func (t *Tray) createWorkingHoursIntervalClicksChannel() <-chan string {
	clicks := make(chan string)
	for interval, item := range t.workingHoursMenuItems {
		go func(interval string, item *systray.MenuItem) {
			for {
				<-item.ClickedCh
				clicks <- interval
			}
		}(interval, item)
	}
	return clicks
}

// Updates the checkmarks for interval selection
func (t *Tray) updateNightModeIntervalChecks(selectedInterval string) {
	for interval, item := range t.workingHoursMenuItems {
		if interval == selectedInterval {
			item.Check()
		} else {
			item.Uncheck()
		}
	}
}

// Adjust visibilities based on the app state
func (t *Tray) applyEnableDisable() {
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

// Init LockKeyboard visibility
func (t *Tray) initLockKeyboard() {
	switch runtime.GOOS {
	case "darwin":
		t.kLockKeyboard.Show()
	default:
		t.kLockKeyboard.Hide()
	}

	t.kUnlockKeyboard.Hide()
}

// Adjust visibilities and activity based on the app state
func (t *Tray) applyEnableDisableKeyboard() {
	if t.keyboardController.KeyboardLocked {
		t.kUnlockKeyboard.Hide()
		t.kLockKeyboard.Show()
	} else {
		t.kUnlockKeyboard.Show()
		t.kLockKeyboard.Hide()
	}
}

func (t *Tray) onExit() {
}
