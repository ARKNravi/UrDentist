// routes/routes.go
package routes

import (
	"log"

	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/gin-gonic/gin"
)

func AppointmentRoutes(r *gin.Engine) {
	controller, err := controller.NewAppointmentController()
	if err != nil {
		log.Fatalf("Failed to create controller: %v", err)
	}

	r.POST("/profile/:profileID/appointment", controller.CreateAppointment)
}
