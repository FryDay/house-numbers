package main

// ForceOn ...
type ForceOn struct {
	On bool `json:"on"`
}

// NewForceOn ...
func NewForceOn() *ForceOn {
	return &ForceOn{false}
}
