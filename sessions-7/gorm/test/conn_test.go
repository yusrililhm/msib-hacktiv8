package test

import (
	"go-gorm/config"
	"go-gorm/model"
	"log"
	"testing"
)

func TestConn(t *testing.T) {
	db, err := config.NewConn()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	d, _ := db.DB()

	if err := d.Ping(); err != nil {
		return
	}

	if err := db.AutoMigrate(model.User{}, model.Product{}); err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Println("Connect success")
}