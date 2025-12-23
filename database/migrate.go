package database

import (
	"template/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []any{
		&model.User{},
		&model.Admin{},
		&model.Card{},
	}
	err := db.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}
