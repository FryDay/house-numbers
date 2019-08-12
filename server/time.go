package main

import (
	"time"
)

// Time ...
type Time struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// NewTime ...
func NewTime(start, end time.Time) *Time {
	return &Time{
		Start: start.Format("3:04 PM"),
		End:   end.Format("3:04 PM"),
	}
}
