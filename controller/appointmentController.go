package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ARKNravi/HACKFEST-BE/model"
	"github.com/ARKNravi/HACKFEST-BE/repository"
	"github.com/gin-gonic/gin"
)

var AppointmentRepo, _ = repository.NewAppointmentRepository()

func GetAppointment(ctx *gin.Context) {
    profileID, err := strconv.Atoi(ctx.Param("profileID"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
        return
    }
    userID := uint(ctx.MustGet("userID").(float64))
    profileRepo := repository.NewProfileRepository()

    profile, err := profileRepo.GetProfile(uint(profileID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if profile.UserID == userID {
        appointmentID, err := strconv.Atoi(ctx.Param("appointmentID"))
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
            return
        }

        var appointment model.Appointment
        if err := AppointmentRepo.Get(&appointment, appointmentID); err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get the appointment"})
            return
        }

        if appointment.ProfileID != uint(profileID) {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "Profile ID does not match the appointment"})
            return
        }

        ctx.JSON(http.StatusOK, gin.H{"appointment": appointment})
        return
    }

    ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to access this appointment"})
}


func GetAllAppointments(ctx *gin.Context) {
	profileID, err := strconv.Atoi(ctx.Param("profileID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	userID := uint(ctx.MustGet("userID").(float64))
	profileRepo := repository.NewProfileRepository()

	profile, err := profileRepo.GetProfile(uint(profileID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if profile.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized"})
		return
	}

	var profiles []model.Profile
	if err := AppointmentRepo.GetProfilesByUserID(&profiles, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profiles"})
		return
	}

	var appointments []model.Appointment
	for _, profile := range profiles {
		var profileAppointments []model.Appointment
		if err := AppointmentRepo.GetAll(&profileAppointments, int(profile.ID)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all appointments"})
			return
		}
		appointments = append(appointments, profileAppointments...)
	}

	ctx.JSON(http.StatusOK, gin.H{"appointments": appointments})
}


func CreateAppointment(ctx *gin.Context) {
	profileID, err := strconv.Atoi(ctx.Param("profileID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	userID := uint(ctx.MustGet("userID").(float64))
	profileRepo := repository.NewProfileRepository()

	profile, err := profileRepo.GetProfile(uint(profileID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if profile.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized"})
		return
	}

	var appointment model.Appointment
	if err := ctx.ShouldBindJSON(&appointment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var consultation interface{}
	if appointment.OnlineConsultationID != nil && *appointment.OnlineConsultationID != 0 {
		var onlineConsultation model.OnlineConsultation
		if err := AppointmentRepo.GetOnlineConsultation(&onlineConsultation, int(*appointment.OnlineConsultationID)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get the online consultation: %v", err)})
			return
		}
		consultation = onlineConsultation
		appointment.OfflineConsultationID = nil  
		appointment.DentistID = onlineConsultation.DentistID  
	}
	
	if appointment.OfflineConsultationID != nil && *appointment.OfflineConsultationID != 0 {
		var offlineConsultation model.OfflineConsultation
		if err := AppointmentRepo.GetOfflineConsultation(&offlineConsultation, int(*appointment.OfflineConsultationID)); err != nil {
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

	if err := AppointmentRepo.Save(&appointment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the appointment"})
		return
	}

	payment := model.Payment{
		Amount:        appointment.TotalPrice * 1.05,
		Status:        false,
		Method:        "",
		AppointmentID: appointment.ID,
	}

	if err := AppointmentRepo.SavePayment(&payment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the payment"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": payment.ID})
}