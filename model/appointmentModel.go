package model

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	PatientName string `gorm:"type:varchar(100);not null"`
	DentistID   uint
	OnlineConsultationID *uint
	OfflineConsultationID *uint
	Day       string `gorm:"type:varchar(100);not null"`
	WorkHour  string       `gorm:"type:varchar(100);not null"`
	TotalPrice  float32 `gorm:"not null"`
	ProfileID   uint
}