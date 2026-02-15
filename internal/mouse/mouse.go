package mouse

import (
	"context"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/sonjek/mouse-stay-up/internal/config"
	"github.com/sonjek/mouse-stay-up/internal/utils"
)

type Controller struct {
	conf         *config.Config
	lastX, lastY int
}

func NewController(conf *config.Config) *Controller {
	return &Controller{
		conf: conf,
	}
}

func (c *Controller) MoveMouse(ctx context.Context) {
	timer := time.NewTimer(0)
	defer timer.Stop()
	<-timer.C // prepare timer for Reset inside sleepCtx

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		// Check if the application is enabled
		if !c.conf.Enabled {
			if !sleepCtx(ctx, timer, 10*time.Second) {
				return
			}
			continue
		}

		// Check if the current time is within working hours.
		// If not, there is no reason to move the cursor.
		if !utils.IsInWorkingHours(time.Now(), c.conf.WorkingHoursInterval) {
			if !sleepCtx(ctx, timer, 1*time.Minute) {
				return
			}
			continue
		}

		// Get current mouse position
		curX, curY := robotgo.Location()

		// Check if the mouse position has changed since the previous run
		// If position has not changed, then move the cursor randomly (between -8px and 8px)
		if c.lastX == curX && c.lastY == curY {
			relX := utils.GetRandomOffset()
			relY := utils.GetRandomOffset()
			robotgo.MoveSmoothRelative(relX, relY)
		}

		// Update last known mouse position
		c.lastX, c.lastY = robotgo.Location()

		// Sleep for a random duration for rendomizing the mouse movement delay
		if !sleepCtx(ctx, timer, utils.GetRandomSleepDuration()) {
			return
		}
	}
}

// sleepCtx blocks for the given duration or until ctx is cancelled.
// Returns true if sleep completed normally, false if the context was cancelled.
func sleepCtx(ctx context.Context, t *time.Timer, d time.Duration) bool {
	t.Reset(d)
	select {
	case <-ctx.Done():
		if !t.Stop() {
			<-t.C
		}
		return false
	case <-t.C:
		return true
	}
}
