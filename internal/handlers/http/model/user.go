package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key" json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `gorm:"unique" json:"phone_number"`
	Address     string    `json:"address"`
	PIN         string    `json:"pin"`
	Balance     int64     `json:"balance"` // Add this line
	CreatedAt   time.Time `json:"created_date"`
	UpdatedAt   time.Time `json:"updated_date,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	return
}
