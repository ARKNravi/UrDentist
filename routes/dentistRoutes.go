package routes

import (
	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/gin-gonic/gin"
)

func DentistRoutes(r *gin.Engine) {
	r.POST("/dentists", controller.CreateDentist)
	r.GET("/dentists", controller.GetAllDentists)
	r.GET("/dentists/:id", controller.GetDentistByID)
}
