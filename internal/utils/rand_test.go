package utils

import (
	"testing"
	"time"
)

func TestGetRandomSleepDuration(t *testing.T) {
	for range 100 {
		d := GetRandomSleepDuration()
		if d < 10*time.Second || d > 60*time.Second {
			t.Errorf("GetRandomSleepDuration() = %v, want [10s, 60s]", d)
		}
	}
}

func TestGetRandomOffset(t *testing.T) {
	for range 100 {
		v := GetRandomOffset()
		if v < -8 || v > 8 {
			t.Errorf("GetRandomOffset() = %v, want [-8, 8]", v)
		}
	}
}
