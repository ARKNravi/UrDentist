package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type RequestBody struct {
	Contents struct {
		Role  string  `json:"role"`
		Parts []*Part `json:"parts"`
	} `json:"contents"`
	SafetySettings struct {
		Category  string `json:"category"`
		Threshold string `json:"threshold"`
	} `json:"safety_settings"`
	GenerationConfig struct {
		Temperature     float64 `json:"temperature"`
		TopP            float64 `json:"topP"`
		TopK            int     `json:"topK"`
		MaxOutputTokens int     `json:"maxOutputTokens"`
	} `json:"generation_config"`
}

type Part struct {
	FileData *struct {
		MimeType string `json:"mimeType"`
		FileURI  string `json:"fileUri"`
	} `json:"fileData,omitempty"`
	Text string `json:"text,omitempty"`
}


func HandleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	if _, err := os.Stat("temp-images"); os.IsNotExist(err) {
		os.Mkdir("temp-images", 0755)
	}

	fileName := handler.Filename

	tempFile, err := os.Create("temp-images/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	_, err = uploadToGCS(fileBytes, tempFile.Name())
	if err != nil {
		fmt.Println(err)
		return
	}

	gcsURI := fmt.Sprintf("gs://supple-hulling-408914.appspot.com/%s", strings.ReplaceAll(tempFile.Name(), "\\", "/"))

	reqBody := &RequestBody{}
	reqBody.Contents.Role = "user"
	reqBody.Contents.Parts = append(reqBody.Contents.Parts, &Part{
		FileData: &struct {
			MimeType string `json:"mimeType"`
			FileURI  string `json:"fileUri"`
		}{
			MimeType: "image/png",
			FileURI:  gcsURI,
		},
	})

	reqBody.Contents.Parts = append(reqBody.Contents.Parts, &Part{
		Text: "I have had toothache in my right rear molar since yesterday. The pain feels like aching and throbbing, especially when I eat or drink something cold or hot. The pain also spread to my ears and jaw. When I checked my teeth, I found that there was a small hole in the affected part of the tooth. The hole measures about 1 millimeter.", 
	})

	fmt.Print(gcsURI)

	reqBody.SafetySettings.Category = "HARM_CATEGORY_SEXUALLY_EXPLICIT"
	reqBody.SafetySettings.Threshold = "BLOCK_LOW_AND_ABOVE"
	reqBody.GenerationConfig.Temperature = 0.4
	reqBody.GenerationConfig.TopP = 1.0
	reqBody.GenerationConfig.TopK = 32
	reqBody.GenerationConfig.MaxOutputTokens = 2048

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://us-central1-aiplatform.googleapis.com/v1/projects/supple-hulling-408914/locations/us-central1/publishers/google/models/gemini-1.0-pro-vision:streamGenerateContent", bytes.NewBuffer(jsonReqBody))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", "Bearer ya29.a0AfB_byBQOYsDF0TlyLJWuGXSflIlnBi8OEKnWEJQZv5Mdomg2c0PBolM48Gf-46Q2NNtzcoZQFfr3V3wQ1vqfe4IPw4B410QBAlArXCZ1F8N9NylKM6zX38wER4IPMJeuv_MojqDeCldtFeuCuiuXbEwgnsuuQ6cD0g6FxCD7ddYaCgYKAdoSARESFQHGX2Mit1RTvd5uAdngT2Vs45JN1Q0179")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

jsonResponse := make([]interface{}, 0)

err = json.Unmarshal(body, &jsonResponse)
if err != nil {
    fmt.Println(err)
    return
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(jsonResponse)

}
