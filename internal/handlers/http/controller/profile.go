package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rzfd/finance-test/internal/handlers/http/model"
	"github.com/rzfd/finance-test/internal/utill"
	"gorm.io/gorm"
)

func UpdateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := utill.ExtractUserID(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthenticated"})
		}

		input := new(struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Address   string `json:"address"`
		})
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		var user model.User
		if err := db.First(&user, "id = ?", userID).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
		}

		user.FirstName = input.FirstName
		user.LastName = input.LastName
		user.Address = input.Address
		user.UpdatedAt = time.Now()

		if err := db.Save(&user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "SUCCESS",
			"result": user,
		})
	}
}
