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

func Pay(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := utill.ExtractUserID(c)
		fmt.Println("Extracted UserID:", userID)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthenticated"})
		}

		input := new(struct {
			Amount  int64  `json:"amount"`
			Remarks string `json:"remarks"`
		})
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		var user model.User
		if err := db.First(&user, "id = ?", uuid.MustParse(userID)).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
		}

		if user.Balance < input.Amount {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Balance is not enough"})
		}

		payment := model.Payment{
			ID:            uuid.New(),
			UserID:        uuid.MustParse(userID),
			Amount:        input.Amount,
			Remarks:       input.Remarks,
			BalanceBefore: user.Balance,
			BalanceAfter:  user.Balance - input.Amount,
			CreatedAt:     time.Now(),
		}

		user.Balance -= input.Amount

		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Save(&user).Error; err != nil {
				return err
			}
			if err := tx.Create(&payment).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Payment failed"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "SUCCESS",
			"result": map[string]interface{}{
				"payment_id":     payment.ID,
				"amount":         payment.Amount,
				"remarks":        payment.Remarks,
				"balance_before": payment.BalanceBefore,
				"balance_after":  payment.BalanceAfter,
				"created_date":   payment.CreatedAt,
			},
		})
	}
}
