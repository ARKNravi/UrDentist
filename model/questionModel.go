package model

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Tag         string `gorm:"type:varchar(100);not null"`
	Question    string `gorm:"type:text;not null"`
	Answer      string `gorm:"type:text"`
	AnsweredAt  *time.Time
	ProfileID   uint
	DentistID   *uint
}