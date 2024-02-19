package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ARKNravi/HACKFEST-BE/model"
	"github.com/ARKNravi/HACKFEST-BE/repository"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
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

func uploadToGCS(data []byte, name string) (string, error) {
	bucketName := "supple-hulling-408914.appspot.com"

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("C:/Users/rkunt/Downloads/supple-hulling-408914-b6a57c3d0519.json"))
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
		GoogleAccessID: "supple-hulling-408914@appspot.gserviceaccount.com",
		PrivateKey:     []byte("-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDROvBWdp7y0xHr\ne1j9Eg87Nog6IJtwXdxmm2QAy//+mxW4oVFZ8II+neBqTJadxpKLWZq6hWZ/yVlT\nmk3htTSbsdTubbyN+yOz03FXPinJ+MtwdOeJ3i0unajJeL/FEzOhthu7WeHvPzWC\nzc+YTOzg1AJff48hW7In/e1oGxGd3p8a8z9ogSxjvUSP2i500lD/fqI/surZwTZQ\nc3Qf/CZXa2+svIXLCjxPjiP1tuHfz0SlQ6dG5HvIfyEMGKHgHDnEmjaVfrVgkByh\nDYlzIVrginIfjECHncjiYwWG32eq91vt4aLygphb7FU4LI9CKiKerpbu55Yhi2+7\noWY++xzDAgMBAAECggEADew3G/nOv+fNtHbDyCtQeic5z2xmC4cjaGyErgzlHwMg\n4eVSLYL0l8gXq9sm1p7lF4LB6hGAbZZvbEHDVvag5o9h1O/WcTg5+vhh/WU0kK0O\nlJAi7CitpwF0vttbH3kUoXklxUTI5Qu2utqJKuBLjvZspgAt/RFF/KVIC/ppJLER\n2rUYPsTrxkJoxMPBuCg+Ry7rYr/hgyFSVebE+onHwI1MiLP4U5cGjOx/GZaQESLc\neOz2IzlPyWF/OuH+jYg9kgPVflGU7tqAPol6rRaNIGoflqR3pZn+KGEINo5+Qvgv\nQjywcuY9K87g3XI9HnkeOclZtAEi+aAfHX9iYiXNKQKBgQD0RSoawbYSuyGhZU3c\nEyKlLtS5P9lZZsbHVeGtc2rNS4ob/wkgh478T0ZFnpohY+riiOz+CCvSbDUYgV0g\n7l38QEaN7MSOypIRUUrz5X72ijw36M90xBT9fyb54YaoRZtTX7P2s3fHZWrNoJ8z\nfOCPEAKJHwR6SqmhXHumA6OILQKBgQDbRwZuVKjJSOmmnpJH8AqEOY2GDAtVC0Ln\nPxGT4PG4PeQyvxsP0bIv3Gk1vieDvrMJqLqCfZ+nYfd2WfATQgBGMJgEV8XCTA+X\nI0e1Rx6m2UFfF65HkAPZKSrAlwpkFkL0R0Tv9Dch2ghUNWVoaJUz0doqQwLtL2Yx\n1d9fwMLerwKBgEveC4TB843/xyM8tqEK5HDPicx7s0McM9MHro9T0LEwrBWj8a7D\ny7o72QSYjSCfyv1PL+R6nzm82ATjcQxgXJqTUBaWmjoLWrC8Qf5cokFqj+eBjKWk\nnSxayL1FubAb5nFPwTJ3bVVl/3UcVTYFrC1i+JakJpzhAayXb+QRL0KFAoGAN5CM\n6aJcTv1B7+3YxY/nKlBnM2OT7431+yE5NA7ZUcWlMNLKabzKeWRR6MNxwemt9rGh\n6XUp4sFpcr0hn8+mwCKKMveG7lBV1weioSYPd1owPYeDqzCsOPg8lCbyBCC8AKia\nqG9rFRHp8GTDeKyfukzgCruGX1IWhGRcwSfYeZMCgYBJ255vut1eylRZmNhylLpn\nEK3//yg8/UeLH/MY8QSoDKvpzhKoqc04iLJCoVb+ZXLyC5u8nY1+Xdkt1Nx4HN/R\nm/fI3ihvTsK5kwSGxZwRpXbO2aT2zw+BSMfSqSPC3stkXbD7NRakzcL4sMgFD40+\ncgC5BnSdJc5BpwNkvxryHA==\n-----END PRIVATE KEY-----\n"),
		Method:         "GET",
		Expires:        time.Now().Add(24 * time.Hour), 
	}

	u, err := storage.SignedURL(bucketName, name, opts)
	if err != nil {
		return "", fmt.Errorf("unable to generate a signed URL: %v", err)
	}

	return u, nil
}



