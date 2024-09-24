package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	Status          string    `json:"status"`
	UserID          uuid.UUID `json:"user_id" gorm:"type:char(36)"`
	TransactionType string    `json:"transaction_type"`
	Amount          int64     `json:"amount"`
	Remarks         string    `json:"remarks"`
	BalanceBefore   int64     `json:"balance_before"`
	BalanceAfter    int64     `json:"balance_after"`
	CreatedAt       time.Time `json:"created_date"`
}
