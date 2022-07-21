package main

import (
	"log"
	"time"
	"token_app/database"
	"token_app/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin() {
	password := "password"
	requestAdmin := models.Admin{}
	requestAdmin.CreatedAt = time.Now().Unix()
	requestAdmin.Email = "admin@gmail.com"
	requestAdmin.FullName = "Tom carter"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		panic(err)
	}
	requestAdmin.Password = string(hashedPassword)

	err = database.Mgr.CreateAdmin(requestAdmin)
	if err != nil {
		panic(err)
	}
	log.Println("Admin has been created")
}
