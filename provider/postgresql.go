package provider

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgresql() *gorm.DB {
	dsn := "host=postgresDB user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
