package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rzfd/finance-test/internal/handlers/http/model"
	"github.com/rzfd/finance-test/internal/utill"
	"gorm.io/gorm"
)

func GetTransactions(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := utill.ExtractUserID(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthenticated"})
		}

		var transactions []model.Transaction
		if err := db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve transactions"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "SUCCESS",
			"result": transactions,
		})
	}
}
