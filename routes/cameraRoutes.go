package routes

import (
	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/ARKNravi/HACKFEST-BE/middleware"
	"github.com/gin-gonic/gin"
)

func CameraRoutes(r *gin.Engine) {
	r.POST("/upload", middleware.AuthMiddleware(), func(c *gin.Context) {
        h := c.Request.Header
        w := c.Writer
        r := c.Request
        r.Header = h
        controller.HandleUpload(w, r)
    })
}
