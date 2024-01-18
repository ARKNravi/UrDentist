package routes

import (
	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/gin-gonic/gin"
)

func CameraRoutes(r *gin.Engine) {
	r.POST("/camera/upload", controller.UploadHandler)
}
