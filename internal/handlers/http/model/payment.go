package model

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID            uuid.UUID `gorm:"type:char(36);primary_key" json:"top_up_id"`
	UserID        uuid.UUID `gorm:"type:char(36);not null;index" json:"user_id"`
	User          *User     `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Amount        int64     `json:"amount_top_up"`
	Remarks       string    `json:"remarks"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedAt     time.Time `json:"created_date"`
}
