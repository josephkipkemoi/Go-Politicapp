package database

import "gorm.io/gorm"

// Migrations func is responsible for creating necessary tables and running our migrations
func Migration(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
	)
}
