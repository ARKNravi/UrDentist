package main

import (
	"log"
	"os"
	"time"

	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
	"github.com/ARKNravi/HACKFEST-BE/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.TempUser{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, User!")
	})

	routes.UserRoutes(r)

	c := cron.New()
	_, err = c.AddFunc("@every 1m", func() {
		tenMinutesAgo := time.Now().Add(-10 * time.Minute)
		controller.DeleteUnverifiedUsers(tenMinutesAgo)
	})
	if err != nil {
		log.Fatal("Error scheduling DeleteUnverifiedUsers job:", err)
	}
	c.Start()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
