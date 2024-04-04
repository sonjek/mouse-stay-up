package config

import (
	"time"
)

const (
	gitRepo string = "https://github.com/sonjek/mouse-stay-up"
)

type Config struct {
	Enabled       bool
	GitRepo       string
	SleepInterval time.Duration
}

func NewConfig() *Config {
	return &Config{
		Enabled:       true,
		GitRepo:       gitRepo,
		SleepInterval: -1,
	}
}

func (c *Config) SetSleepInterval(interval time.Duration) {
	c.SleepInterval = interval
}

func (c *Config) SetSleepIntervalSec(sec int) {
	c.SetSleepInterval(time.Duration(sec) * time.Second)
}
