package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"github.com/rzfd/finance-test/internal/handlers/http/controller"
	"github.com/rzfd/finance-test/internal/utill"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConn := utill.ConnectDB()

	e := echo.New()

	jwtSecret := os.Getenv("JWT_SECRET")

	e.POST("/register", controller.Register(dbConn, jwtSecret))
	e.POST("/login", controller.Login(dbConn, jwtSecret))

	protected := e.Group("")
	protected.Use(utill.JWTMiddleware(jwtSecret))
	protected.POST("/topup", controller.TopUp(dbConn))
	protected.POST("/pay", controller.Pay(dbConn))
	protected.POST("/transfer", controller.Transfer(dbConn))
	protected.PUT("/profile", controller.UpdateUser(dbConn))
	protected.GET("/transaction", controller.GetTransactions(dbConn))

	e.Logger.Fatal(e.Start(":8080"))
}
