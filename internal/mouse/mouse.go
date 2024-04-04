package mouse

import (
	"math/rand/v2"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/sonjek/mouse-stay-up/internal/config"
	"github.com/sonjek/mouse-stay-up/internal/utils"
)

type Controller struct {
	config       *config.Config
	LastX, LastY int
}

func NewController(config *config.Config) *Controller {
	return &Controller{
		config: config,
	}
}

func (c *Controller) MoveMouse() {
	for c.config.Enabled {
		// Sleep before the check
		c.sleep()

		// Check if the current time is within working hours.
		// If not, there is no reason to move the cursor.
		isWorkingHours := utils.IsInWorkingHours(c.config.WorkingHoursInterval)
		if !isWorkingHours {
			continue
		}

		// Get current mouse position
		curX, curY := robotgo.Location()

		// Check if the mouse position has changed since the previous run
		// If position has not changed, then move the cursor
		if c.LastX == curX && c.LastY == curY {
			// Random movement distance along the X-axis and Y-axis (between -8 and 8)
			relX := (rand.N(9) - 4) * 2
			relY := (rand.N(9) - 4) * 2

			// Move the cursor
			robotgo.MoveSmoothRelative(relX, relY)
		}

		// Update last known mouse position
		c.LastX, c.LastY = robotgo.Location()
	}
}

func (c *Controller) sleep() {
	// Get random duration between 10-60 sec
	duration := time.Duration(rand.N(51)+10) * time.Second
	time.Sleep(duration)
}
