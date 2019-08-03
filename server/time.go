package main

import (
	"time"
)

// Time ...
type Time struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// NewTime ...
func NewTime(start, end time.Time) *Time {
	return &Time{
		Start: start,
		End:   end,
	}
}
