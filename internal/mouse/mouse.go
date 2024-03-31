package mouse

import (
	"math/rand/v2"
	"time"

	"github.com/go-vgo/robotgo"
)

type Controller struct {
	Enabled       bool
	GitRepo       string
	SleepInterval time.Duration
	LastX, LastY  int
}

func NewController(gitRepo string) *Controller {
	return &Controller{
		Enabled: true,
		GitRepo: gitRepo,
	}
}

func (c *Controller) MoveMouse() {
	for c.Enabled {
		// Get current mouse position
		curX, curY := robotgo.Location()

		// Check if the mouse position has changed since the previous run
		// If position has not changed, then move the cursor
		if c.LastX == curX && c.LastY == curY {
			// Random movement distance along the X-axis and Y-axis (between 2 and 5)
			relX := (rand.N(8) - 4) * 2
			relY := (rand.N(8) - 4) * 2

			// Move the cursor
			robotgo.MoveSmoothRelative(relX, relY)
		}

		// Get new mouse position
		newX, newY := robotgo.Location()

		// Update last known mouse position
		c.LastX, c.LastY = newX, newY

		time.Sleep(c.SleepInterval)
	}
}

func (c *Controller) SetSleepInterval(interval time.Duration) {
	c.SleepInterval = interval
}

func (c *Controller) SetSleepIntervalSec(sec int) {
	c.SetSleepInterval(time.Duration(sec) * time.Second)
}
