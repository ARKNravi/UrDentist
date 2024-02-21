package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ARKNravi/HACKFEST-BE/model"
	"github.com/ARKNravi/HACKFEST-BE/repository"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	repo repository.PaymentRepository
}

func NewPaymentController(repo repository.PaymentRepository) *PaymentController {
	return &PaymentController{repo: repo}
}

func (c *PaymentController) UpdatePayment(ctx *gin.Context) {
	paymentID, err := strconv.Atoi(ctx.Param("paymentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer openedFile.Close()

	data, err := ioutil.ReadAll(openedFile)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photoURL, err := uploadToGCS(data, file.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload the photo: %v", err)})
		return
	}	

	var payment model.Payment
	amount, _ := strconv.ParseFloat(ctx.PostForm("Amount"), 32)
	payment.Amount = float32(amount)
	payment.Status = true 
	payment.Method = ctx.PostForm("Method")
	appointmentID, _ := strconv.Atoi(ctx.PostForm("AppointmentID"))
	payment.AppointmentID = uint(appointmentID)

	payment.ID = uint(paymentID)
	payment.Photo = photoURL

	if err := c.repo.Update(&payment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the payment"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"payment": payment})
}

func (c *PaymentController) UpdatePaymentDummy(ctx *gin.Context) {
	paymentID, err := strconv.Atoi(ctx.Param("paymentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	var payment model.Payment
	amount, _ := strconv.ParseFloat(ctx.PostForm("Amount"), 32)
	payment.Amount = float32(amount)
	payment.Status = true 
	payment.Method = ctx.PostForm("Method")
	appointmentID, _ := strconv.Atoi(ctx.PostForm("AppointmentID"))
	payment.AppointmentID = uint(appointmentID)
	
	payment.ID = uint(paymentID)
	payment.Photo = "y"
	
	if err := c.repo.Update(&payment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the payment"})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message": "payment successfully"})
}	


func uploadToGCS(data []byte, name string) (string, error) {
	bucketName := os.Getenv("BUCKET_NAME")

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create new storage client: %w", err)
	}
	defer client.Close()

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := w.Write(data); err != nil {
		return "", fmt.Errorf("failed to write data to object: %w", err)
	}
	if err := w.Close(); err != nil {
		return "", fmt.Errorf("failed to close object writer: %w", err)
	}

	opts := &storage.SignedURLOptions{
		GoogleAccessID: os.Getenv("GOOGLE_ACCESS_ID"),
		PrivateKey:     []byte(os.Getenv("PRIVATE_KEY")),
		Method:         "GET",
		Expires:        time.Now().Add(24 * time.Hour),
	}

	u, err := storage.SignedURL(bucketName, name, opts)
	if err != nil {
		return "", fmt.Errorf("unable to generate a signed URL: %v", err)
	}

	return u, nil
}



