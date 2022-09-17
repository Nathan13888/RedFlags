package main

import (
	"time"
	"uuid"
)

type StreetEntry struct {
	Time        time.Time
	ImageLink   string
	SafetyScore int
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
