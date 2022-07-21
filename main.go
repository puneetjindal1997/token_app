/*
 *	Author:- Puneet Jindal
 *
 *	Main.go
 */
package main

import (
	"log"
	"net/http"
	"os"
	"token_app/database"
	"token_app/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}

type routes struct {
	router *gin.Engine
}
type Routes []Route

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// route calling
	database.DbInit()
	database.TxnDb.AutoMigrate(&models.CreateToken{})
	database.TxnDb.AutoMigrate(&models.UserLogin{})
	database.TxnDb.AutoMigrate(&models.Admin{})
	// creating admin
	CreateAdmin()
	// routing
	ClientRoutes()

}

/*
 *	Function for grouping log routes
 */
func (r routes) TokenGrouping(rg *gin.RouterGroup) {
	orderRouteGrouping := rg.Group("/admin")
	for _, route := range AdminRoutes {
		switch route.Method {
		case "GET":
			orderRouteGrouping.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			orderRouteGrouping.POST(route.Pattern, route.HandlerFunc)
		case "OPTIONS":
			orderRouteGrouping.OPTIONS(route.Pattern, route.HandlerFunc)
		case "PUT":
			orderRouteGrouping.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			orderRouteGrouping.DELETE(route.Pattern, route.HandlerFunc)
		default:
			orderRouteGrouping.GET(route.Pattern, func(c *gin.Context) {
				c.JSON(200, gin.H{
					"result": "Specify a valid http method with this route.",
				})
			})
		}
	}
}

/*
 *	Function for grouping UserLogin routes
 */
func (r routes) UserLoginGrouping(rg *gin.RouterGroup) {
	tradeRouteGrouping := rg.Group("/auth")
	tradeRouteGrouping.Use(CORSMiddleware())
	for _, route := range UserLoginRoutes {
		switch route.Method {
		case "GET":
			tradeRouteGrouping.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			tradeRouteGrouping.POST(route.Pattern, route.HandlerFunc)
		case "OPTIONS":
			tradeRouteGrouping.OPTIONS(route.Pattern, route.HandlerFunc)
		case "PUT":
			tradeRouteGrouping.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			tradeRouteGrouping.DELETE(route.Pattern, route.HandlerFunc)
		default:
			tradeRouteGrouping.GET(route.Pattern, func(c *gin.Context) {
				c.JSON(200, gin.H{
					"result": "Specify a valid http method with this route.",
				})
			})
		}
	}
}

// append routes with versions
func ClientRoutes() {
	r := routes{
		router: gin.Default(),
	}
	v1 := r.router.Group(os.Getenv("API_VERSION"))
	r.TokenGrouping(v1)
	r.UserLoginGrouping(v1)
	if err := r.router.Run(":" + os.Getenv("CLIENT_PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// Middlewares
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
		}
	}
}
