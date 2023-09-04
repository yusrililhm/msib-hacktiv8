package database

import (
	"fmt"
	"go-jwt/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	username = helper.GetEnv("PGUSERNAME")
	password = helper.GetEnv("PGPASSWORD")
	host     = helper.GetEnv("HOST")
	port     = helper.GetEnv("PORTPG")
	sslMode  = helper.GetEnv("SSLMODE")
	dbName   = helper.GetEnv("DBNAME")
	timezone = helper.GetEnv("TIMEZONE")
)

func NewConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=%s", username, password, host, port, dbName, sslMode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}
