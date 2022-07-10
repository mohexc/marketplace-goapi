package profile

import (
	"fmt"
	"marketplace-goapi/modules/base"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	base.Model
	UserId    string    `json:"user_id"`
	Addresses []Address `gorm:"foreignKey:AddressRefer"`
	Avatar    string    `json:"avatar"`
}

type Address struct {
	base.Model
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Zip          string    `json:"zip"`
	Country      string    `json:"country"`
	ProfileRefer uuid.UUID ``
}

var postgresqlDB *gorm.DB

func MigrateProfile(db *gorm.DB) error {
	postgresqlDB = db
	err := postgresqlDB.AutoMigrate(&Profile{})
	if err != nil {
		fmt.Println("error in migrate profile")
		return err
	}
	err = postgresqlDB.AutoMigrate(&Address{})
	if err != nil {
		fmt.Println("error in migrate address")
		return err
	}
	return nil
}
