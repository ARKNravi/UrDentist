package model

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	PatientName string `gorm:"type:varchar(100);not null"`
	DentistID   uint
}