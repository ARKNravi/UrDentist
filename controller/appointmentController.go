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
	repo repository.AppointmentRepository
}

func NewAppointmentController() (*AppointmentController, error) {
	repo, err := repository.NewAppointmentRepository()
	if err != nil {
		return nil, err
	}
	return &AppointmentController{repo: repo}, nil
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
	if *appointment.OnlineConsultationID != 0 {
		var onlineConsultation model.OnlineConsultation
		if err := c.repo.GetOnlineConsultation(&onlineConsultation, int(*appointment.OnlineConsultationID)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get the online consultation: %v", err)})
			return
		}
		consultation = onlineConsultation
		appointment.OfflineConsultationID = nil  
		appointment.DentistID = onlineConsultation.DentistID  
	} else if *appointment.OfflineConsultationID != 0 {
		var offlineConsultation model.OfflineConsultation
		if err := c.repo.GetOfflineConsultation(&offlineConsultation, int(*appointment.OfflineConsultationID)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get the offline consultation: %v", err)})
			return
		}
		consultation = offlineConsultation
		appointment.OnlineConsultationID = nil  
		appointment.DentistID = offlineConsultation.DentistID  
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Either OnlineConsultationID or OfflineConsultationID must be set"})
		return
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save the appointment: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"appointment": appointment})
}
