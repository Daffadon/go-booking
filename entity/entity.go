package entity

import (
	"time"

	"gorm.io/gorm"
)

type TimeStamp struct {
	CreatedAt time.Time ` gorm:"type: timestamp with time zone" json:"created_at"`
	UpdateAt  time.Time ` gorm:"type: timestamp with time zone" json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
