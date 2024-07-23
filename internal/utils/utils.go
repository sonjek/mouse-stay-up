package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// layout is the time format layout
const timeLayout = "15:04"

func OpenWebPage(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return errors.New("unsupported platform")
	}

	return cmd.Start()
}

func IsInWorkingHours(timeWindow string) bool {
	parts := strings.Split(timeWindow, "-")

	// If the start and end times are equal, it means the interval spans the entire day
	if parts[0] == parts[1] {
		return true
	}

	// Get the current date
	currentTime := time.Now()

	// Generate a datetime 3-14 minutes earlier than the specified time
	startTimeBefore := addRandomTimeInterval(currentTime, parts[0])

	// Generate a datetime 3-14 minutes later than the specified time
	endTimeAfter := subtractRandomTimeInterval(currentTime, parts[1])

	// Check if the current time is within the dynamic range
	if currentTime.After(startTimeBefore) && currentTime.Before(endTimeAfter) {
		return true
	}
	return false
}

func addRandomTimeInterval(currentTime time.Time, inputTime string) time.Time {
	shiftTime, err := time.Parse(timeLayout, inputTime)
	if err != nil {
		panic(err)
	}

	shiftTime = shiftTime.Add(generateTimeInterval())

	// Combine the current date with the parsed time
	return combineNowAndShiftTime(currentTime, shiftTime)
}

func subtractRandomTimeInterval(currentTime time.Time, inputTime string) time.Time {
	shiftTime, err := time.Parse(timeLayout, inputTime)
	if err != nil {
		panic(err)
	}

	shiftTime = shiftTime.Add(-generateTimeInterval())

	// Combine the current date with the parsed time
	return combineNowAndShiftTime(currentTime, shiftTime)
}

// Generate a random interval between 3 to 14 minutes
func generateTimeInterval() time.Duration {
	n, err := rand.Int(rand.Reader, big.NewInt(12))
	if err != nil {
		panic(err)
	}

	return time.Duration(n.Int64()+3) * time.Minute
}

// Combine the current date with the parsed time (hour-minute)
func combineNowAndShiftTime(now, shiftTime time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(),
		shiftTime.Hour(), shiftTime.Minute(), now.Second(), 0, now.Location())
}

// Random duration between 10-60 sec
func GetRandomSleepDuration() time.Duration {
	max := big.NewInt(51) // Upper bound is 51 to get values from 0 to 50
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return time.Duration(n.Int64()+10) * time.Second
}

// Random movement distance between -8 and 8
func GetRandomOffset() int {
	n, err := rand.Int(rand.Reader, big.NewInt(9))
	if err != nil {
		panic(err)
	}
	return int((n.Int64() - 4) * 2)
}
