package repository

import (
	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
)

type PaymentRepository struct {}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}

func (r *PaymentRepository) Update(payment *model.Payment) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	return db.Model(payment).Updates(payment).Error
}