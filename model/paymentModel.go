package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Amount       float32 `gorm:"not null"`
	Status       string  `gorm:"type:varchar(100);not null"` 
	Method       string  `gorm:"type:varchar(100);not null"` 
	AppointmentID uint
}