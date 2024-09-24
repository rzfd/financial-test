package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/rzfd/finance-test/internal/handlers/http/model"
	"github.com/rzfd/finance-test/internal/utill"
)

func Register(db *gorm.DB, jwtSecret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(model.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		var existingUser model.User
		if err := db.Where("phone_number = ?", user.PhoneNumber).First(&existingUser).Error; err == nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "User already exists"})
		}

		if err := db.Create(user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Registration failed"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "SUCCESS",
			"result": map[string]interface{}{
				"user_id":      user.ID.String(),
				"first_name":   user.FirstName,
				"last_name":    user.LastName,
				"phone_number": user.PhoneNumber,
				"address":      user.Address,
				"created_date": user.CreatedAt.Format("2006-01-02 15:04:05"),
			},
		})
	}
}

func Login(db *gorm.DB, jwtSecret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		type LoginInput struct {
			PhoneNumber string `json:"phone_number"`
			PIN         string `json:"pin"`
		}

		input := new(LoginInput)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		var user model.User
		if err := db.Where("phone_number = ?", input.PhoneNumber).First(&user).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Phone Number and PIN doesnt match"})
		}

		if input.PIN != user.PIN {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Phone Number and PIN doesnt match"})
		}

		token, err := utill.GenerateToken(user.ID.String(), jwtSecret)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "SUCCESS",
			"result": map[string]string{
				"access_token":  token,
				"refresh_token": "",
			},
		})
	}
}
