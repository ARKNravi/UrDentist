// repository/paymentRepository.go
package repository

import (
	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Save(payment *model.Payment) error
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

func (r *paymentRepository) Save(payment *model.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) Update(payment *model.Payment) error {
	return r.db.Save(payment).Error
}
