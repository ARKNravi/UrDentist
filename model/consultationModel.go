package model

import (
	"time"

	"gorm.io/gorm"
)

type OnlineConsultation struct {
	gorm.Model
	Day       time.Weekday `gorm:"type:varchar(100);not null"`
	WorkHour  string       `gorm:"type:varchar(100);not null"`
	ChatPrice float32      `gorm:"not null"`
	CallPrice float32      `gorm:"not null"`
	VidCallPrice float32   `gorm:"not null"`
	DentistID uint
}

type OfflineConsultation struct {
	gorm.Model
	Place     string `gorm:"type:varchar(100);not null"`
	Day       time.Weekday `gorm:"not null"`
	WorkHour  string `gorm:"type:varchar(100);not null"`
	Price     float32      `gorm:"not null"`
	DentistID uint
}