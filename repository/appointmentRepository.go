// repository/appointmentRepository.go
package repository

import (
	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	Save(appointment *model.Appointment) error
	SavePayment(payment *model.Payment) error
	GetProfile(profile *model.Profile, profileID int) error
	GetOnlineConsultation(consultation *model.OnlineConsultation, consultationID int) error
	GetOfflineConsultation(consultation *model.OfflineConsultation, consultationID int) error
}

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository() (AppointmentRepository, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &appointmentRepository{db: db}, nil
}

func (r *appointmentRepository) Save(appointment *model.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *appointmentRepository) GetProfile(profile *model.Profile, profileID int) error {
	return r.db.First(profile, profileID).Error
}

func (r *appointmentRepository) GetOnlineConsultation(consultation *model.OnlineConsultation, consultationID int) error {
    return r.db.First(consultation, consultationID).Error
}

func (r *appointmentRepository) GetOfflineConsultation(consultation *model.OfflineConsultation, consultationID int) error {
	return r.db.First(consultation, consultationID).Error
}

func (r *appointmentRepository) SavePayment(payment *model.Payment) error {
	return r.db.Create(payment).Error
}

type OfflineConsultationRepository struct {}

func NewOfflineConsultationRepository() *OfflineConsultationRepository {
	return &OfflineConsultationRepository{}
}

func (r *OfflineConsultationRepository) GetOfflineConsultationsByDentistID(dentistID uint) ([]model.OfflineConsultation, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var consultations []model.OfflineConsultation
	result := db.Where("dentist_id = ?", dentistID).Find(&consultations)
	return consultations, result.Error
}

type OnlineConsultationRepository struct {}

func NewOnlineConsultationRepository() *OnlineConsultationRepository {
	return &OnlineConsultationRepository{}
}

func (r *OnlineConsultationRepository) GetOnlineConsultationsByDentistID(dentistID uint) ([]model.OnlineConsultation, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var consultations []model.OnlineConsultation
	result := db.Where("dentist_id = ?", dentistID).Find(&consultations)
	return consultations, result.Error
}