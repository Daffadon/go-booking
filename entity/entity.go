package entity

import (
	"time"

	"gorm.io/gorm"
)

type TimeStamp struct {
	CreatedAt time.Time ` gorm:"type: timestamp with time zone; autoCreateTime; notNull" json:"created_at"`
	UpdateAt  time.Time ` gorm:"type: timestamp with time zone; autoUpdateTime; notNull" json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
