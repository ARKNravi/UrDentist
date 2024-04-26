package repository

import (
	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
	"gorm.io/gorm"
)

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository() (*AppointmentRepository, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &AppointmentRepository{db: db}, nil
}

func (r *AppointmentRepository) Save(appointment *model.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *AppointmentRepository) GetProfile(profile *model.Profile, profileID int) error {
	return r.db.First(profile, profileID).Error
}

func (r *AppointmentRepository) GetAll(appointments *[]model.Appointment, profileID int) error {
	return r.db.Where("profile_id = ?", profileID).Find(appointments).Error
}

func (r *AppointmentRepository) Get(appointment *model.Appointment, id int) error {
	return r.db.First(appointment, id).Error
}

func (r *AppointmentRepository) GetOnlineConsultation(consultation *model.OnlineConsultation, consultationID int) error {
    return r.db.First(consultation, consultationID).Error
}

func (r *AppointmentRepository) GetOfflineConsultation(consultation *model.OfflineConsultation, consultationID int) error {
	return r.db.First(consultation, consultationID).Error
}

func (r *AppointmentRepository) SavePayment(payment *model.Payment) error {
	return r.db.Create(payment).Error
}

func (r *AppointmentRepository) GetOfflineConsultationsByDentistID(dentistID uint) ([]model.OfflineConsultation, error) {
	var consultations []model.OfflineConsultation
	result := r.db.Where("dentist_id = ?", dentistID).Find(&consultations)
	return consultations, result.Error
}

func (r *AppointmentRepository) GetOnlineConsultationsByDentistID(dentistID uint) ([]model.OnlineConsultation, error) {
	var consultations []model.OnlineConsultation
	result := r.db.Where("dentist_id = ?", dentistID).Find(&consultations)
	return consultations, result.Error
}
func (r *AppointmentRepository) GetProfilesByUserID(profiles *[]model.Profile, userID uint) error {
	return r.db.Where("user_id = ?", userID).Find(profiles).Error
}

