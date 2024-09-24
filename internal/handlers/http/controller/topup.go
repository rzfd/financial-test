package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rzfd/finance-test/internal/handlers/http/model"
	"github.com/rzfd/finance-test/internal/utill"
	"gorm.io/gorm"
)

func TopUp(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := utill.ExtractUserID(c)
		fmt.Println("Extracted UserID:", userID)
		fmt.Println(userID)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthenticated"})
		}

		input := new(struct {
			Amount int64 `json:"amount"`
		})
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		topUp := model.TopUp{
			ID:        uuid.New(),
			UserID:    uuid.MustParse(userID),
			Amount:    input.Amount,
			CreatedAt: time.Now(),
		}

		var user model.User
		if err := db.First(&user, topUp.UserID).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
		}

		user.Balance += topUp.Amount

		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Save(&user).Error; err != nil {
				return err
			}
			if err := tx.Create(&topUp).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Top-up failed"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "SUCCESS",
			"result": map[string]interface{}{
				"top_up_id":      topUp.ID,
				"amount_top_up":  topUp.Amount,
				"balance_before": user.Balance - topUp.Amount,
				"balance_after":  user.Balance,
				"created_date":   topUp.CreatedAt,
			},
		})
	}
}
