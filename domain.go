package main

import "time"

type Counter struct {
	Days        int       `json:"days"`
	CurrentDate time.Time `json:"current_date"`
}

type Metrics struct {
	Id    string    `json:"id,omitempty"`
	Value string    `json:"value"`
	Type  string    `json:"type"`
	Time  time.Time `json:"time,omitempty"`
}
