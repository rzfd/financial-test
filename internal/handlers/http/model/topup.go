package model

import (
	"time"

	"github.com/google/uuid"
)

type TopUp struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"top_up_id"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;index" json:"user_id"`
	User      *User     `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Amount    int64     `json:"amount_top_up"`
	CreatedAt time.Time `json:"created_date"`
}
