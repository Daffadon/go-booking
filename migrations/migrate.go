package migrations

import (
	"go-booking/entity"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		panic(err)
	}
	log.Println("Migration has been processed")
	return nil
}
