package postgres

import (
	"log"

	"auth/core/user"
	"auth/storage"
)

// Migrate updates the db with new columns, and tables automatically during application start
// PS: as the app grows we might need to use a third party migration tool.
func Migrate(database *storage.Database) {
	err := database.DB.AutoMigrate(
		user.User{},
	)

	if err != nil {
		log.Println(err)
	}
}
