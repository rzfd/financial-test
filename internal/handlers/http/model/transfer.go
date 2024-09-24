package model

import (
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	ID            uuid.UUID `json:"transfer_id" gorm:"type:char(36);primary_key"`
	SenderID      uuid.UUID `json:"sender_id" gorm:"type:char(36)"`
	ReceiverID    uuid.UUID `json:"receiver_id" gorm:"type:char(36)"`
	Amount        int64     `json:"amount"`
	Remarks       string    `json:"remarks" gorm:"type:longtext"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedAt     time.Time `json:"created_date"`
}
