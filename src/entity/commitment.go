package entity

import "time"

type Commitment struct {
	Title     string
	StartTime time.Time
	EndTime   time.Time
	Place     string
}
