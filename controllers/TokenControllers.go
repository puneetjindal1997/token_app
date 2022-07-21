package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"token_app/database"
	"token_app/helpers"
	"token_app/models"

	"github.com/gin-gonic/gin"
)

/*
 *	Function to save the token details
 *
 *	return c.JSON
 */
func CreateToken(c *gin.Context) {
	requestToken := models.CreateToken{}
	postBodyErr := json.NewDecoder(c.Request.Body).Decode(&requestToken)
	if postBodyErr != nil {
		log.Println("postBodyErr ---------- ", postBodyErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": postBodyErr.Error()})
		return
	}
	rand.Seed(time.Now().UnixNano())
	// random token
	requestToken.Token = helpers.RandomString(6)
	fmt.Println("requestToken.Token----------->>", requestToken.Token)
	// created date
	requestToken.CreatedAt = time.Now().Unix()
	// valid till 7 days
	requestToken.TokenExpiration = time.Now().Unix() + 604800
	errResp := database.Mgr.SaveTokenDetails(requestToken)
	if errResp != nil {
		log.Println("token details not save in db", errResp)
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "success"})
}
