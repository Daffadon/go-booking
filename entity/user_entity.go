package entity

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	TelpNumber string    `json:"telp_number"`
	Age        int       `json:"age"`
	Role       string    `json:"role"`
	IsVerified bool      `json:"is_verified"`

	TimeStamp
}
