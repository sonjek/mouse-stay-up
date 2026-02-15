package utils

import (
	"testing"
	"time"
)

var windowFromNineToEighteen = "09:00-19:00"

func TestIsInWorkingHours(t *testing.T) {
	tests := []struct {
		name   string
		now    time.Time
		window string
		want   bool
	}{
		{"FullDay", time.Now(), "00:00-00:00", true},
		{"Inside", time.Date(2024, 4, 7, 12, 0, 0, 0, time.UTC), windowFromNineToEighteen, true},
		{"AtStart", time.Date(2024, 4, 7, 9, 0, 0, 0, time.UTC), windowFromNineToEighteen, true},
		{"BeforeStart", time.Date(2024, 4, 7, 8, 59, 0, 0, time.UTC), windowFromNineToEighteen, false},
		{"AtEnd", time.Date(2024, 4, 7, 19, 0, 0, 0, time.UTC), windowFromNineToEighteen, false},
		{"AfterEnd", time.Date(2024, 4, 7, 20, 0, 0, 0, time.UTC), windowFromNineToEighteen, false},
		{"InvalidFormat", time.Now(), "invalid", false},
		{"InvalidStart", time.Now(), "xx:00-19:00", false},
		{"InvalidEnd", time.Now(), "09:00-yy:00", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInWorkingHours(tt.now, tt.window); got != tt.want {
				t.Errorf("IsInWorkingHours(%v, %q) = %v, want %v", tt.now, tt.window, got, tt.want)
			}
		})
	}
}

func TestMinutesSinceMidnight(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want int
	}{
		{"Midnight", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), 0},
		{"OneMinute", time.Date(2024, 1, 1, 0, 1, 0, 0, time.UTC), 1},
		{"Noon", time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), 720},
		{"EndOfDay", time.Date(2024, 1, 1, 23, 59, 59, 0, time.UTC), 1439},
		{"SecondsIgnored", time.Date(2024, 1, 1, 9, 30, 45, 0, time.UTC), 570},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minutesSinceMidnight(tt.time); got != tt.want {
				t.Errorf("minutesSinceMidnight(%v) = %d, want %d", tt.time, got, tt.want)
			}
		})
	}
}

func TestParseHHMM(t *testing.T) {
	tests := []struct {
		input  string
		want   int
		wantOK bool
	}{
		{"00:00", 0, true},
		{"09:30", 570, true},
		{"23:59", 1439, true},
		{"invalid", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, ok := parseTimeToMinutes(tt.input)
			if ok != tt.wantOK || got != tt.want {
				t.Errorf("parseTimeToMinutes(%q) = (%d, %v), want (%d, %v)", tt.input, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}
