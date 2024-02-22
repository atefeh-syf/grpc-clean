package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DbConn() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open("host=localhost user=postgres dbname=yamigo_db password=admin port=5432 sslmode=disable"), &gorm.Config{},
	)
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	return db
}
