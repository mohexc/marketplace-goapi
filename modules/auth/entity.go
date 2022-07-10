package auth

import (
	"marketplace-goapi/modules/base"
	"marketplace-goapi/modules/profile"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	base.Model
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	ProfileID uuid.UUID
	Profile   profile.Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

var postgresqlDB *gorm.DB

func MigrateUser(db *gorm.DB) error {
	postgresqlDB = db
	err := postgresqlDB.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
