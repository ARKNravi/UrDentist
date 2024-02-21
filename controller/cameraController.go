package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
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

func getAccessToken() (string, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create new storage client: %w", err)
	}

	bucketName := os.Getenv("BUCKET_NAME")
	objectName := os.Getenv("OBJECT_NAME")

	rc, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get credentials file from GCS: %w", err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return "", fmt.Errorf("failed to read credentials file: %w", err)
	}

	creds, err := google.CredentialsFromJSON(ctx, data, os.Getenv("GOOGLE_AUTH_URL"))
	if err != nil {
		return "", err
	}

	tokenSource := creds.TokenSource
	token, err := tokenSource.Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {

	accessToken, err := getAccessToken()
	if err != nil {
		fmt.Println(err)
		return
	}

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

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	
	_, err = uploadToGCS(fileBytes, handler.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	gcsURI := fmt.Sprintf("gs://%s/%s", os.Getenv("BUCKET_NAME"), handler.Filename)
	text := r.FormValue("text")
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
		Text: "how much percentage from  0% to 100% caries is that be precise only number no alphabet or symbol only number% and  " + text +"what to do when i have that much caries can you give me the solution and what to avoid", 
	})

	fmt.Print(gcsURI)

	reqBody.SafetySettings.Category = os.Getenv("SAFETY_CATEGORY")
	reqBody.SafetySettings.Threshold = os.Getenv("SAFETY_THRESHOLD")
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

	req, err := http.NewRequest("POST", os.Getenv("API_URL"), bytes.NewBuffer(jsonReqBody))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("streamGenerateContent API call failed with status code %d and response body: %s\n", resp.StatusCode, string(bodyBytes))
		return
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	
	jsonResponse := make([]map[string]interface{}, 0)
	
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	var persenText, restText string

	for _, item := range jsonResponse {
		candidates := item["candidates"].([]interface{})
		for _, candidate := range candidates {
			content := candidate.(map[string]interface{})["content"].(map[string]interface{})
			parts := content["parts"].([]interface{})
			for _, part := range parts {
				text := part.(map[string]interface{})["text"].(string)
				splitText := strings.SplitN(text, "\n", 2)
				re := regexp.MustCompile(`(\d+%).*`)
				matches := re.FindStringSubmatch(splitText[0])
				if len(matches) > 1 {
					persenText += matches[1] + " "
				}
				if len(splitText) > 1 {
					restText += splitText[1] + " "
				}
			}
		}
	}
	
	responseData := make(map[string]string)
	responseData["persen"] = persenText
	responseData["text"] = restText
	
	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(responseData)
}