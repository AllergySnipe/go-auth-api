package models

import "time"

type RevokedToken struct {
	ID        uint   `gorm:"primaryKey"`
	Token     string `gorm:"uniqueIndex"`
	RevokedAt time.Time
}
