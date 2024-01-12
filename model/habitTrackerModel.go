package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null"`
	Points      int    `gorm:"not null"`
}

type TaskCompletion struct {
	gorm.Model
	ProfileID uint `gorm:"not null"`
	TaskID    uint `gorm:"not null"`
	Completed bool `gorm:"type:boolean;not null"`
	Date      *time.Time `gorm:"type:date;not null"`
	Profile   Profile `gorm:"foreignKey:ProfileID"`
	Task      Task `gorm:"foreignKey:TaskID"`
}
