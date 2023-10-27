package db

import (
	//"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(dbURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//db.AutoMigrate(&YourModel{}) // Replace YourModel with your actual model struct.
	return db, nil
}
