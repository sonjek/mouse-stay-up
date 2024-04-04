package config

import (
	"time"
)

const (
	gitRepo string = "https://github.com/sonjek/mouse-stay-up"

	workingHours00_00 = "00:00-00:00"
	workingHours08_18 = "08:00-18:00"
	workingHours09_19 = "09:00-19:00"
	workingHours10_19 = "10:00-19:00"
	workingHours10_20 = "10:00-20:00"
)

var workingHours = []string{
	workingHours08_18,
	workingHours09_19,
	workingHours10_19,
	workingHours10_20,
	workingHours00_00,
}

type Config struct {
	Enabled              bool
	GitRepo              string
	SleepInterval        time.Duration
	WorkingHoursInterval string
	WorkingHours         []string
}

func NewConfig() *Config {
	return &Config{
		Enabled:              true,
		GitRepo:              gitRepo,
		SleepInterval:        -1,
		WorkingHoursInterval: workingHours10_19,
		WorkingHours:         workingHours,
	}
}

func (c *Config) SetSleepInterval(interval time.Duration) {
	c.SleepInterval = interval
}

func (c *Config) SetSleepIntervalSec(sec int) {
	c.SetSleepInterval(time.Duration(sec) * time.Second)
}

func (c *Config) SetWorkingHoursInterval(interval string) {
	c.WorkingHoursInterval = interval
}
