package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/ARKNravi/HACKFEST-BE/model"
	"github.com/ARKNravi/HACKFEST-BE/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "random"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load .env file:", err)
	}

	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}
}

type UserController struct{}

type UserResponse struct {
	ID            uint   `json:"id"`
	FullName      string `json:"full_name"`
	NoPhone       string `json:"no_phone"`
	EmailAddress  string `json:"email_address"`
	IsVerified    bool   `json:"is_verified"`
}

func (UserController) Register(c *gin.Context) {
	var user model.TempUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Password != user.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	user.Password = string(hashedPassword)
	user.IsVerified = false

	loc, _ := time.LoadLocation("Asia/Jakarta")
	user.CreatedAt = time.Now().In(loc)

	user.VerificationCode = generateVerificationCode()
	repository := repository.NewUserRepository()
	if err := repository.CreateTempUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := sendVerificationEmail(user.EmailAddress, user.VerificationCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending verification email"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Verification email has been sent. Please verify your account within 10 minutes."})
}


func DeleteUnverifiedUsers(tenMinutesAgo time.Time) {
	repository := repository.NewUserRepository()
	err := repository.DeleteUnverifiedUsers(tenMinutesAgo)
	if err != nil {
		log.Println("Error deleting unverified users:", err)
	}
}

func (UserController) Verify(c *gin.Context) {
	var verificationRequest struct {
		Code string `json:"code"`
	}
	if err := c.ShouldBindJSON(&verificationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository := repository.NewUserRepository()
	user, err := repository.FindTempUserByCode(verificationRequest.Code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
		}
		return
	}

	user.IsVerified = true

	if err := repository.MoveUserToDB(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error moving user to database"})
		return
	}

	userResponse := UserResponse{
		ID:            user.ID,
		FullName:      user.FullName,
		NoPhone:       user.NoPhone,
		EmailAddress:  user.EmailAddress,
		IsVerified:    user.IsVerified,
	}

	c.JSON(http.StatusOK, gin.H{"status": "User verified successfully!", "user": userResponse})
}

func generateVerificationCode() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	rand.Seed(time.Now().In(loc).UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(10000)) // generates a random number between 0 and 9999
}


func sendVerificationEmail(email string, code string) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load .env file: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verification")
	m.SetBody("text/html", "Your verification code is: "+code)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func (UserController) Login(c *gin.Context) {
	var loginInfo struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository := repository.NewUserRepository()
	user, err := repository.FindUserByEmail(loginInfo.EmailAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().In(loc).Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "User logged in successfully!", "user": user, "token": tokenString})
}

func (UserController) GoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (UserController) GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid oauth state"))
		return
	}

	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	userInfo, err := getUserInfo(token)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	name, ok := userInfo["name"].(string)
	if !ok {
		name = "Unknown"
	}

	user := &model.User{
		FullName:      name,
		EmailAddress:  userInfo["email"].(string),
	}
	repository := repository.NewUserRepository()
	if err := repository.FindOrCreateUserByEmail(user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().In(loc).Add(time.Hour * 72).Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "User logged in successfully!", "user": user, "token": tokenString})
}

func getUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo map[string]interface{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}


func (UserController) ForgotPassword(c *gin.Context) {
	var request struct {
		EmailAddress string `json:"email_address"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository := repository.NewUserRepository()
	foundUser, err := repository.FindUserByEmail(request.EmailAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
		return
	}

	foundUser.VerificationCode = generateVerificationCode()

	if err := repository.UpdateUser(foundUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	if err := sendVerificationEmail(foundUser.EmailAddress, foundUser.VerificationCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending verification email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Verification email has been sent."})
}


func (UserController) VerifyToken(c *gin.Context) {
	var request struct {
		EmailAddress    string `json:"email_address"`
		VerificationCode string `json:"verification_code"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository := repository.NewUserRepository()
	foundUser, err := repository.FindUserByEmail(request.EmailAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
		return
	}

	if foundUser.VerificationCode != request.VerificationCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": foundUser.EmailAddress,
		"exp":   time.Now().In(loc).Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY_PASSWORD")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Token verified successfully!", "token": tokenString})
}

func (UserController) ResetPassword(c *gin.Context) {
	var request struct {
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Password != request.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	email := c.MustGet("email").(string)
	repository := repository.NewUserRepository()
	foundUser, err := repository.FindUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	foundUser.Password = string(hashedPassword)
	if err := repository.UpdateUser(foundUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Password reset successfully!"})
}


func (UserController) ShowProfile(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	userID := claims["userID"].(string)
	repository := repository.NewUserRepository()
	user, err := repository.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (UserController) ResendVerificationCode(c *gin.Context) {
	var request struct {
		EmailAddress string `json:"email_address"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository := repository.NewUserRepository()
	user, err := repository.FindTempUserByEmail(request.EmailAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
		return
	}

	user.VerificationCode = generateVerificationCode()

	if err := repository.UpdateTempUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	if err := sendVerificationEmail(user.EmailAddress, user.VerificationCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending verification email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Verification email has been resent. Please verify your account within 10 minutes."})
}



