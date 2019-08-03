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
		Start: start.Format("15:04"),
		End:   end.Format("15:04"),
	}
}
