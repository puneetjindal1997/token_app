package database

import (
	"log"
	"token_app/models"
)

// tables
const (
	AdminTable = "admin"
	TokenTable = "token"
)

// Function to create admin
// return error
func (mgr *manager) CreateAdmin(request models.Admin) error {
	resp := mgr.connection.Table(AdminTable).Create(&request)
	return resp.Error
}

/*
 *	Function to get the token details and save them
 *	return id int, error
 */
func (mgr *manager) SaveTokenDetails(request models.CreateToken) error {
	resp := mgr.connection.Table(TokenTable).Create(&request)
	return resp.Error
}

/*
 *	Function to get the token using token id
 *
 *	return models.CreateToken, error
 */
func (mgr *manager) GetTokenWithToken(token string) (resp models.CreateToken, err error) {
	log.Println("method token", token)
	dbResp := mgr.connection.Table(TokenTable).Where("token = ?", token).Find(&resp)
	err = dbResp.Error
	return resp, err
}

// get admin
func (mgr *manager) GetUserByEmailId(email string) (user models.Admin, err error) {
	dbResp := mgr.connection.Table(AdminTable).Where("email = ?", email).Find(&user)
	return user, dbResp.Error
}
