package utils

import (
	"math/rand/v2"
	"time"
)

// GetRandomSleepDuration returns a random duration between 10 and 60 seconds.
func GetRandomSleepDuration() time.Duration {
	return time.Duration(rand.IntN(51)+10) * time.Second //nolint:gosec // cryptographic randomness not needed for sleep
}

// GetRandomOffset returns a random movement offset in the range [-8, 8]
// (even values only: -8, -6, -4, -2, 0, 2, 4, 6, 8).
func GetRandomOffset() int {
	return (rand.IntN(9) - 4) * 2 //nolint:gosec // cryptographic randomness not needed for mouse offset
}
