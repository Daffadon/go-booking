package entity

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password" gorm:"type:varchar(255)"`
	TelpNumber string    `json:"telp_number"`
	Age        uint8     `json:"age"`
	Role       string    `json:"role" gorm:"default:'user'"`
	IsVerified bool      `json:"is_verified"`

	TimeStamp
}
