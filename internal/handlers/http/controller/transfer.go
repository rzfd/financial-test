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

func Transfer(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := utill.ExtractUserID(c)
		fmt.Println("Extracted UserID:", userID)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthenticated"})
		}

		input := new(struct {
			TargetUser string `json:"target_user"`
			Amount     int64  `json:"amount"`
			Remarks    string `json:"remarks"`
		})
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		var sender model.User
		if err := db.First(&sender, "id = ?", uuid.MustParse(userID)).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
		}

		if sender.Balance < input.Amount {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Balance is not enough"})
		}

		var receiver model.User
		if err := db.First(&receiver, "id = ?", uuid.MustParse(input.TargetUser)).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Target user not found"})
		}

		transfer := model.Transfer{
			ID:            uuid.New(),
			SenderID:      sender.ID,
			ReceiverID:    receiver.ID,
			Amount:        input.Amount,
			Remarks:       input.Remarks,
			BalanceBefore: sender.Balance,
			BalanceAfter:  sender.Balance - input.Amount,
			CreatedAt:     time.Now(),
		}

		sender.Balance -= input.Amount
		receiver.Balance += input.Amount

		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Save(&sender).Error; err != nil {
				return err
			}
			if err := tx.Save(&receiver).Error; err != nil {
				return err
			}
			if err := tx.Create(&transfer).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Transfer failed"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "SUCCESS",
			"result": map[string]interface{}{
				"transfer_id":    transfer.ID,
				"amount":         transfer.Amount,
				"remarks":        transfer.Remarks,
				"balance_before": transfer.BalanceBefore,
				"balance_after":  transfer.BalanceAfter,
				"created_date":   transfer.CreatedAt,
			},
		})
	}
}
