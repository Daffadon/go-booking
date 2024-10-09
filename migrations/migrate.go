package migrations

import (
	"go-booking/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		panic(err)
	}
	return nil
}
