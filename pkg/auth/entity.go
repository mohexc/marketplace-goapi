package auth

import (
	"fmt"
	"marketplace-goapi/pkg/base"

	"gorm.io/gorm"
)

type User struct {
	base.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

var postgresqlDB *gorm.DB

func MigrateUser(db *gorm.DB) error {
	postgresqlDB = db
	err := postgresqlDB.AutoMigrate(&User{})
	if err != nil {

		return err
	}
	fmt.Println("User table migrated")
	return nil
}
