// routes/routes.go
package routes

import (
	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/ARKNravi/HACKFEST-BE/middleware"
	"github.com/gin-gonic/gin"
)

func AppointmentRoutes(r *gin.Engine) {
	r.GET("/profile/:profileID/appointments",  middleware.AuthMiddleware(), controller.GetAllAppointments)
	r.GET("/profile/:profileID/appointment/:appointmentID", middleware.AuthMiddleware(), controller.GetAppointment)
	r.POST("/profile/:profileID/appointment", middleware.AuthMiddleware(), controller.CreateAppointment)
}
