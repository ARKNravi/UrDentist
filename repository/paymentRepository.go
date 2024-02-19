// repository/paymentRepository.go
package repository

import (
	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Update(payment *model.Payment) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository() (PaymentRepository, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &paymentRepository{db: db}, nil
}

func (r *paymentRepository) Update(payment *model.Payment) error {
	return r.db.Model(payment).Updates(payment).Error
}

