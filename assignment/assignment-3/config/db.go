package config

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	username = Init("PGUSERNAME")
	password = Init("PGPASSWORD")
	host = Init("HOST")
	port = Init("PORTPG")
	sslMode = Init("SSLMODE")
	dbName = Init("DBNAME")
	timezone = Init("TIMEZONE") 
)


func NewConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=%s", username, password, host, port, dbName, sslMode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}
