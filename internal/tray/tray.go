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
	conf                  *config.Config
	workingHoursMenuItems map[string]*systray.MenuItem
}

func NewTray(mouseController *mouse.Controller, conf *config.Config) *Tray {
	return &Tray{
		mouseController:       mouseController,
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
	mEnable := systray.AddMenuItem("Enable", "Enable mouse movement")
	mDisable := systray.AddMenuItem("Disable", "Disable mouse movement")
	mWorkingHours := systray.AddMenuItem("Working hours", "Select a range of working hours")
	systray.AddSeparator()
	mAbout := systray.AddMenuItem("About", "Open GitHub repo")
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	// Hide the enable option since it's already enabled by default
	if t.conf.Enabled {
		mEnable.Hide()
	} else {
		mDisable.Hide()
	}

	// Add interval selection submenu items
	for _, hours := range t.conf.WorkingHours {
		t.addWorkingHoursItems(mWorkingHours, hours)
	}

	// Set a marker for the default working hours interval
	t.workingHoursMenuItems[t.conf.WorkingHoursInterval].Check()

	workingHoursIntervalClicks := t.createWorkingHoursIntervalClicksChannel()

	go func() {
		for {
			select {
			case <-mEnable.ClickedCh:
				t.conf.ToggleEnableDisable()
				mEnable.Hide()
				mDisable.Show()
				mWorkingHours.Enable()
				go t.mouseController.MoveMouse()
			case <-mDisable.ClickedCh:
				t.conf.ToggleEnableDisable()
				mDisable.Hide()
				mEnable.Show()
				mWorkingHours.Disable()
			case workingHoursInterval := <-workingHoursIntervalClicks:
				// When an hours interval item is clicked, update the workingHoursInterval interval and checkmarks
				t.conf.SetWorkingHoursInterval(workingHoursInterval)
				t.updateNightModeIntervalChecks(t.conf.WorkingHoursInterval)
			case <-mAbout.ClickedCh:
				if err := utils.OpenWebPage(t.conf.GitRepo); err != nil {
					panic(err)
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

	// Start moving the mouse in a circle immediately if enabled
	if t.conf.Enabled {
		go t.mouseController.MoveMouse()
	}
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

func (t *Tray) onExit() {
}
