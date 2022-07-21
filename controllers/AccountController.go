package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
	"token_app/database"
	"token_app/helpers"
	"token_app/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// admin login
func AdminLogin(c *gin.Context) {
	login := models.AdminLogin{}
	loginResp := models.LoginResponse{}
	c.ShouldBindJSON(&login)
	userResp, err := database.Mgr.GetUserByEmailId(login.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "You are not a authorized user."})
		return
	}
	if passErr := bcrypt.CompareHashAndPassword([]byte(userResp.Password), []byte(login.Password)); passErr != nil {
		// If the two passwords don't match, return an error
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "You are not a authorized user."})
		return
	}
	jwtWrapper := helpers.JwtWrapper{
		SecretKey:       os.Getenv("JWTSecret"),
		Issuer:          os.Getenv("JWTIssuer"),
		ExpirationHours: 48,
	}
	signedToken, jwtErr := jwtWrapper.GenerateToken(userResp.Id, userResp.Email)
	if jwtErr != nil {
		log.Println("Token genration error", jwtErr)
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "You are not a authorized user."})
		return
	}
	loginResp.Token = signedToken
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success", "data": loginResp})
}

/*
 *	Function to login
 *
 *	response c.JSON
 */
func Login(c *gin.Context) {
	requestLogin := models.UserLogin{}
	postBodyErr := json.NewDecoder(c.Request.Body).Decode(&requestLogin)
	if postBodyErr != nil {
		log.Println("postBodyErr ---------- ", postBodyErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": postBodyErr.Error()})
		return
	}
	if requestLogin.Token == "" {
		log.Println("Validation error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": errors.New("please pass a valid details")})
		return
	}
	tokenData, err := database.Mgr.GetTokenWithToken(requestLogin.Token)
	if err != nil {
		log.Println("token details not feched from db", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "internal server error"})
		return
	}
	// checking time expiration at loging
	if tokenData.TokenExpiration < time.Now().Unix() {
		log.Println("token expired")
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "token expired"})
		return
	}
	// create structure for user
	requestLogin.Email = "test@gmail.com"
	requestLogin.UserName = "test"
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success", "data": requestLogin})
}
