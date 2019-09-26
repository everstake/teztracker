package models

import "time"

type Fee struct {
	Low       uint      `json:"low"`
	Medium    uint      `json:"medium"`
	High      uint      `json:"high"`
	Timestamp time.Time `json:"timestamp"`
	Kind      string    `json:"kind"`
}
