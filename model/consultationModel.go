package model

import (
	"gorm.io/gorm"
)

type OfflineConsultation struct {
	gorm.Model
	Place     string `gorm:"type:varchar(100);not null"`
	Day       string `gorm:"type:varchar(100);not null"`
	WorkHour  string `gorm:"type:varchar(100);not null"`
	Price     float32      `gorm:"not null"`
	DentistID uint
	ServiceID uint
	Appointment []Appointment `gorm:"foreignKey:OfflineConsultationID"`
}

type OnlineConsultation struct {
	gorm.Model
	Day       string `gorm:"type:varchar(100);not null"`
	WorkHour  string       `gorm:"type:varchar(100);not null"`
	Price     float32      `gorm:"not null"`
	DentistID uint
	ServiceID uint
	Appointment []Appointment `gorm:"foreignKey:OnlineConsultationID"`
}

