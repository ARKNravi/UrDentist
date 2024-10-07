package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OrderID       string      `json:"order_id" gorm:"type:varchar(50);not null;unique"`
	TransactionID string      `json:"transaction_id" gorm:"type:varchar(100);not null;"`
	Amount        int         `json:"amount" gorm:"type:int;not null"`
	Method        string      `json:"method" gorm:"type:varchar(100);not null"`
	VaNumber      string      `json:"va_number" gorm:"type:varchar(100);not null"`
	Status        string      `json:"status" gorm:"type:varchar(100);not null"`
	AppointmentID uint        `json:"appointment_id" gorm:"not null"`
	UserID        uint        `json:"user_id" gorm:"not null"`
	Appointment   Appointment `gorm:"foreignKey:AppointmentID"`
	Profile       Profile     `gorm:"foreignKey:UserID"`
}

type TransactionPost struct {
	Amount        int    `json:"amount" binding:"required"`
	Method        string `json:"method" binding:"required"`
	AppointmentID uint   `json:"appointment_id" binding:"required"`
}

type TransactionByID struct {
	ID            uint   `json:"id"`
	Method        string `json:"method"`
	Amount        int    `json:"amount"`
	VANumber      string `json:"va_number"`
	OrderID       string `json:"order_id"`
	Status        string `json:"status"`
	AppointmentID uint   `json:"appointment_id"`
}

type TransactionByUser struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Pict      string    `json:"pict"`
	Title     string    `json:"title"`
	Amount    int       `json:"amount"`
	Method    string    `json:"method"`
	Status    string    `json:"status"`
}