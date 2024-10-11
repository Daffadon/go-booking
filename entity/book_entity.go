package entity

import (
	"github.com/google/uuid"

	"github.com/shopspring/decimal"
)

type Book struct {
	ID          uuid.UUID       `json:"id" form:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title       string          `json:"title" form:"title"`
	Author      string          `json:"author" form:"author"`
	Cover       string          `json:"cover" form:"cover"`
	Description string          `json:"description" form:"description"`
	Stock       int             `json:"stock" form:"stock"`
	Price       decimal.Decimal `json:"price" form:"price" gorm:"type:decimal(10,2)"`

	TimeStamp
}
