package database

import (
	"fmt"
	"os"
	"token_app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type manager struct {
	connection *gorm.DB
}

var TxnDb *gorm.DB

var Mgr Manager

// db manager interface
type Manager interface {
	// login admin
	GetUserByEmailId(string) (models.Admin, error)
	// admin db call to create admin
	CreateAdmin(models.Admin) error
	// save token detail to db
	SaveTokenDetails(models.CreateToken) error
	// SaveAccount(models.UserLogin) error
	GetTokenWithToken(string) (models.CreateToken, error)
}

/*
 *	init funcion to create connection with database while running the server
 */
func DbInit() {
	dsn := os.Getenv("SQLUSERNAME") + ":" + os.Getenv("SQLPASSWORD") + "@tcp(" + os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT") + ")/" + os.Getenv("SQLDATABASE")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("not connected")
		return
	}
	TxnDb = db
	Mgr = &manager{connection: db}
}
