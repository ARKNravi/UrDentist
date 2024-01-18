package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VertexAIResponse struct {
	Predictions []struct {
		Classes []string `json:"classes"`
	} `json:"predictions"`
}

type Response struct {
	Classification string `json:"classification"`
}

func UploadHandler(c *gin.Context) {
	// Parse the multipart form
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the file from form data
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Open the file
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer openedFile.Close()

	// Read the file into a byte array
	fileBytes, err := ioutil.ReadAll(openedFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vertexAIResponse, err := SendToVertexAI(fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check that Predictions and Classes are not empty
	if len(vertexAIResponse.Predictions) > 0 && len(vertexAIResponse.Predictions[0].Classes) > 0 {
		// Create the response
		response := Response{
			Classification: vertexAIResponse.Predictions[0].Classes[0],
		}

		// Convert the response to JSON
		responseJSON, err := json.Marshal(response)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Write the response
		c.Data(http.StatusOK, "application/json", responseJSON)
		fmt.Println(vertexAIResponse)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "No predictions returned by the model"})
		fmt.Println(vertexAIResponse)

	}
}

func SendToVertexAI(image []byte) (*VertexAIResponse, error) {
	// Create a new POST request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "image.jpg")
	if err != nil {
		return nil, err
	}
	part.Write(image)
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://vertexai.googleapis.com/v1/projects/supple-hulling-408914/locations/us-central1/endpoints/2073461226683236352:predict", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Log the status code and body of the response
	fmt.Println("Status code:", res.StatusCode)
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Body:", string(bodyBytes))

	// Parse the response
	vertexAIResponse := &VertexAIResponse{}
	json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(vertexAIResponse)

	return vertexAIResponse, nil
}

