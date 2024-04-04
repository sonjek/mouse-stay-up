package tray

import (
	"embed"
	"fmt"
	"io"

	"github.com/getlantern/systray"

	"github.com/sonjek/mouse-stay-up/internal/config"
	"github.com/sonjek/mouse-stay-up/internal/mouse"
	"github.com/sonjek/mouse-stay-up/internal/utils"
)

var (
	//go:embed icon.png
	iconFile embed.FS
)

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
	mouseController *mouse.Controller
	config          *config.Config
	intervalItems   map[int]*systray.MenuItem
}

func NewTray(mouseController *mouse.Controller, config *config.Config) *Tray {
	return &Tray{
		mouseController: mouseController,
		config:          config,
		intervalItems:   make(map[int]*systray.MenuItem),
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

	// Create menu items for enable/disable mouse movement, change sleep interval and exit
	mEnable := systray.AddMenuItem("Enable", "Enable mouse movement")
	mDisable := systray.AddMenuItem("Disable", "Disable mouse movement")
	mInterval := systray.AddMenuItem("Check Interval", "Set mouse movement interval")
	systray.AddSeparator()
	mAbout := systray.AddMenuItem("About", "Open GitHub repo")
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	// Hide the enable option since it's already enabled by default
	if t.config.Enabled {
		mEnable.Hide()
	} else {
		mDisable.Hide()
	}

	// Add interval selection submenu items
	t.addIntervalItem(mInterval, "10-60 sec", -1)
	t.addIntervalItem(mInterval, "30 sec", 30)
	t.addIntervalItem(mInterval, "60 sec", 60)

	// Mark the default interval
	t.intervalItems[int(t.config.SleepInterval)].Check()

	// Create a channel to listen for interval item clicks
	intervalClicks := t.createIntervalClicksChannel()

	go func() {
		for {
			select {
			case <-mEnable.ClickedCh:
				t.config.Enabled = true
				mEnable.Hide()
				mDisable.Show()
				mInterval.Enable()
				go t.mouseController.MoveMouse()
			case <-mDisable.ClickedCh:
				t.config.Enabled = false
				mDisable.Hide()
				mEnable.Show()
				mInterval.Disable()
			case interval := <-intervalClicks:
				// When an interval item is clicked, update the sleep interval and checkmarks
				t.config.SetSleepIntervalSec(interval)
				t.updateIntervalChecks(interval)
			case <-mAbout.ClickedCh:
				utils.OpenWebPage(t.config.GitRepo)
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

	// Start moving the mouse in a circle immediately if enabled
	if t.config.Enabled {
		go t.mouseController.MoveMouse()
	}
}

// Adds a submenu item for selecting a sleep interval
func (t *Tray) addIntervalItem(parent *systray.MenuItem, title string, interval int) {
	t.intervalItems[interval] = parent.AddSubMenuItem(title, "Set interval to "+title)
}

// Creates and returns a channel that listens to all interval item clicks
func (t *Tray) createIntervalClicksChannel() <-chan int {
	clicks := make(chan int)
	for interval, item := range t.intervalItems {
		go func(interval int, item *systray.MenuItem) {
			for {
				<-item.ClickedCh
				clicks <- interval
			}
		}(interval, item)
	}
	return clicks
}

// Updates the checkmarks for interval selection
func (t *Tray) updateIntervalChecks(selectedInterval int) {
	for interval, item := range t.intervalItems {
		if interval == selectedInterval {
			item.Check()
		} else {
			item.Uncheck()
		}
	}
}

func (t *Tray) onExit() {
}
