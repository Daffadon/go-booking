package cmd

import (
	"fmt"
	"go-booking/migrations"
	"os"

	"gorm.io/gorm"
)

func DatabaseCommand(db *gorm.DB) {
	migrate := false
	seed := false

	for _, arg := range os.Args[1:] {
		if arg == "migrate" {
			migrate = true
		}
		if arg == "seed" {
			seed = true
		}
	}
	if migrate {
		migrations.Migrate(db)
	}
	if seed {
		fmt.Println("Seeding data...")
		// Seed(db)
	}

}
