package util

import "time"

// Clock interface for time-related operations
type Clock interface {
	Now() time.Time
}

// RealClock uses the actual time
type RealClock struct{}

// Now returns the current time
func (r *RealClock) Now() time.Time {
	return time.Now()
}
