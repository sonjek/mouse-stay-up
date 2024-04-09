package config

import (
	"fmt"

	"github.com/sonjek/mouse-stay-up/pkg/config"
)

const (
	GitRepo        string = "https://github.com/sonjek/mouse-stay-up"
	appName        string = "mouse-stay-up"
	configFileName string = "config.conf"

	workingHours00_00 = "00:00-00:00"
	workingHours08_18 = "08:00-18:00"
	workingHours09_19 = "09:00-19:00"
	workingHours10_19 = "10:00-19:00"
	workingHours10_20 = "10:00-20:00"
)

var (
	configFilePath string

	workingHours = []string{
		workingHours08_18,
		workingHours09_19,
		workingHours10_19,
		workingHours10_20,
		workingHours00_00,
	}
)

type Config struct {
	Enabled              bool     `ini:"enabled"`
	WorkingHoursInterval string   `ini:"working-hours-interval"`
	WorkingHours         []string `ini:"working-hours,omitempty"`
}

func init() {
	// Create a folder to save the config file if it doesn't exist yet
	if err := config.CreateConfigFolder(appName); err != nil {
		fmt.Println(err)
	}

	// Calculate the configuration file path for saving when app parameters change
	configFilePath = config.GetConfigFilePath(appName, configFileName)
}

func NewConfig() *Config {
	// Define the default Config struct
	appConfig := &Config{
		Enabled:              true,
		WorkingHoursInterval: workingHours10_19,
		WorkingHours:         workingHours,
	}

	// Restores the configuration from disk if it exists
	isRestored := config.LoadFileToStruct(configFilePath, appConfig)

	// Saves the struct to an configuration file if it doesn't exist yet
	if !isRestored {
		saveConfig(appConfig)
	}

	return appConfig
}

func (c *Config) ToggleEnableDisable() {
	c.Enabled = !c.Enabled

	saveConfig(c)
}

func (c *Config) SetWorkingHoursInterval(interval string) {
	c.WorkingHoursInterval = interval

	saveConfig(c)
}

func saveConfig(appConfig *Config) {
	if err := config.SaveStructToFile(configFilePath, appConfig); err != nil {
		fmt.Println(err)
	}
}
