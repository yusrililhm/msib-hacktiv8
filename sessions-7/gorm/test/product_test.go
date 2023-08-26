package test

import (
	"encoding/json"
	"fmt"
	"go-gorm/config"
	"go-gorm/model"
	"log"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	db, err := config.NewConn()
	product := model.Product{
		UserId: 7,
		Brand: "Apple",
		Name: "iPad",
	}

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	p, err := model.NewProductRepository(db).CreateProduct(product)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}