package model

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	Score      float32 `gorm:"not null"`
	Comment    string  `gorm:"type:text"`
	DentistID  uint
	ProfileID  uint
}
