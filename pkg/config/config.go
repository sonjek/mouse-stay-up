package config

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/ini.v1"
)

// Create a folder to save the config file if it doesn't exist yet
func CreateConfigFolder(appName string) error {
	// Use xdg to get the config file folder path
	configPath, err := xdg.ConfigFile(appName)
	if err != nil {
		return fmt.Errorf("error getting config file path: %w", err)
	}

	// Create a home folder if it does not exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(configPath, 0o750); err != nil {
			return fmt.Errorf("error creating config directory: %w", err)
		}
	}

	return nil
}

// Calculate the configuration file path
func GetConfigFilePath(appName, configFileName string) string {
	configPath, err := xdg.ConfigFile(appName + "/" + configFileName)
	if err != nil {
		fmt.Println("Error getting config file path:", err)
		return ""
	}
	return configPath
}

// Saves the struct to an configuration file
func SaveStructToFile(configFilePath string, config any) error {
	// Reflect data sources from given struct
	cfg := ini.Empty()
	if err := ini.ReflectFrom(cfg, config); err != nil {
		return fmt.Errorf("error reflecting config struct: %w", err)
	}

	// Save the configuration file into disk
	if err := cfg.SaveTo(configFilePath); err != nil {
		return fmt.Errorf("error saving config to file: %w", err)
	}

	return nil
}

// Restores the configuration file into the struct
func LoadFileToStruct(configFilePath string, config any) bool {
	// Try to load the configuration file from disk if it exists
	cfg, err := ini.Load(configFilePath)
	if err != nil {
		return false
	}

	// Try to map the configuration file from disk into the struct
	if err := cfg.MapTo(config); err != nil {
		fmt.Printf("error mapping INI to config struct: %s", err.Error())
		return false
	}

	return true
}
