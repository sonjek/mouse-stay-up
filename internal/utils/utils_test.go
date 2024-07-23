package utils

import (
	"testing"
	"time"
)

// Test value is within the range of 3 to 14 minutes
func TestGenerateTimeInterval(t *testing.T) {
	for range 100 {
		result := generateTimeInterval()
		if result < 3*time.Minute || result >= 15*time.Minute {
			t.Errorf("generateTimeInterval() = %v, want %v to %v", result, 3*time.Minute, 15*time.Minute)
		}
	}
}

// Test checks if the function merges two times correctly
func TestCombineNowAndShiftTime(t *testing.T) {
	now := time.Date(2024, 4, 7, 10, 0, 45, 0, time.UTC)
	shiftTime := time.Date(2000, 1, 1, 15, 30, 0, 0, time.UTC)
	expected := time.Date(2024, 4, 7, 15, 30, 45, 0, time.UTC)
	result := combineNowAndShiftTime(now, shiftTime)

	if !result.Equal(expected) {
		t.Errorf("combineNowAndShiftTime(%v, %v) = %v, want %v", now, shiftTime, result, expected)
	}
}

// Test value is within the range 10-60 sec
func TestGetRandomSleepDuration(t *testing.T) {
	for range 100 {
		result := GetRandomSleepDuration()
		if result < 10*time.Second || result > 60*time.Second {
			t.Errorf("GetRandomSleepDuration() = %v, want %v to %v", result, 10*time.Second, 60*time.Second)
		}
	}
}

// Test value is within the range -8 and 8
func TestGetRandomOffset(t *testing.T) {
	for range 100 {
		result := GetRandomOffset()
		if result < -8 || result > 8 {
			t.Errorf("GetRandomSleepDuration() = %v, want %v to %v", result, -8, 8)
		}
	}
}
