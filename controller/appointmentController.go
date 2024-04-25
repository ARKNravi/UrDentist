package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ARKNravi/HACKFEST-BE/model"
	"github.com/ARKNravi/HACKFEST-BE/repository"
	"github.com/gin-gonic/gin"
)

type AppointmentController struct {
	repo *repository.AppointmentRepository
}

func NewAppointmentController() (*AppointmentController, error) {
	repo, err := repository.NewAppointmentRepository()
	if err != nil {
		return nil, err
	}
	return &AppointmentController{repo: repo}, nil
}

func (c *AppointmentController) GetAppointment(ctx *gin.Context) {
	profileID, err := strconv.Atoi(ctx.Param("profileID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	appointmentID, err := strconv.Atoi(ctx.Param("appointmentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	var appointment model.Appointment
	if err := c.repo.Get(&appointment, appointmentID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get the appointment"})
		return
	}

	if appointment.ProfileID != uint(profileID) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Profile ID does not match the appointment"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"appointment": appointment})
}

func (c *AppointmentController) GetAllAppointments(ctx *gin.Context) {
	profileID, err := strconv.Atoi(ctx.Param("profileID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	var appointments []model.Appointment
	if err := c.repo.GetAll(&appointments, profileID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all appointments"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"appointments": appointments})
}

func (c *AppointmentController) CreateAppointment(ctx *gin.Context) {
	profileID, err := strconv.Atoi(ctx.Param("profileID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dentist ID"})
		return
	}

	var appointment model.Appointment
	if err := ctx.ShouldBindJSON(&appointment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var profile model.Profile
	if err := c.repo.GetProfile(&profile, profileID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get the profile"})
		return
	}

	appointment.PatientName = profile.NamaLengkap

	var consultation interface{}
	if appointment.OnlineConsultationID != nil && *appointment.OnlineConsultationID != 0 {
		var onlineConsultation model.OnlineConsultation
		if err := c.repo.GetOnlineConsultation(&onlineConsultation, int(*appointment.OnlineConsultationID)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get the online consultation: %v", err)})
			return
		}
		consultation = onlineConsultation
		appointment.OfflineConsultationID = nil  
		appointment.DentistID = onlineConsultation.DentistID  
	}
	
	if appointment.OfflineConsultationID != nil && *appointment.OfflineConsultationID != 0 {
		var offlineConsultation model.OfflineConsultation
		if err := c.repo.GetOfflineConsultation(&offlineConsultation, int(*appointment.OfflineConsultationID)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get the offline consultation: %v", err)})
			return
		}
		consultation = offlineConsultation
		appointment.OnlineConsultationID = nil  
		appointment.DentistID = offlineConsultation.DentistID  
	}
	appointment.ProfileID = uint(profileID)
	switch v := consultation.(type) {
	case model.OnlineConsultation:
		appointment.WorkHour = v.WorkHour
		appointment.Day = v.Day
		appointment.TotalPrice = v.Price
	case model.OfflineConsultation:
		appointment.WorkHour = v.WorkHour
		appointment.Day = v.Day
		appointment.TotalPrice = v.Price
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Consultation is neither online nor offline"})
		return
	}

	if err := c.repo.Save(&appointment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the appointment"})
		return
	}

	payment := model.Payment{
		Amount:        appointment.TotalPrice * 1.05, 
		Status:        false,         
		Method:        "",         
		AppointmentID: appointment.ID, 
	}

	if err := c.repo.SavePayment(&payment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the payment"})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"id": payment.ID})
}