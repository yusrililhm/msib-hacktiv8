package config

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	user = "postgres"
	password = "postgres" 
	host = "localhost"
	port = "5432"
	dbName = "test-gorm"
)

func NewConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&Timezone=Asia/Jakarta", user, password, host, port, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return db, nil
}
