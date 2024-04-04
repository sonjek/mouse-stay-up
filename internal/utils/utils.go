package utils

import (
	"errors"
	"fmt"
	"math/rand/v2"
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

	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}

func IsInWorkingHours(timeWindow string) bool {
	parts := strings.Split(timeWindow, "-")

	// If the start and end times are equal, it means the interval spans the entire day
	if parts[0] == parts[1] {
		return true
	}

	startTimeBefore := MakeRandomTime(parts[0])
	endTimeAfter := MakeRandomTime(parts[1])
	currentTime := time.Now()

	// Check if the current time is within the dynamic range
	if currentTime.After(startTimeBefore) && currentTime.Before(endTimeAfter) {
		return true
	}
	return false
}

func MakeRandomTime(inputTime string) time.Time {
	// Parse input time like 19:00 to date time
	parsedTime, err := time.Parse(timeLayout, inputTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return time.Time{}
	}

	// Generate a random interval between 3 to 14 minutes
	randomInterval := time.Duration(rand.N(12)+3) * time.Minute

	// Randomly decide whether to add or subtract the interval
	if rand.N(2) == 0 {
		parsedTime = parsedTime.Add(randomInterval)
	} else {
		parsedTime = parsedTime.Add(-randomInterval)
	}

	// Get the current date
	now := time.Now()

	// Combine the current date with the parsed time
	return time.Date(now.Year(), now.Month(), now.Day(),
		parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0, now.Location())
}
