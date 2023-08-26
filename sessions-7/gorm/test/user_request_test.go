package test

import (
	"encoding/json"
	"fmt"
	"go-gorm/config"
	"go-gorm/model"
	"go-gorm/repository"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, err := config.NewConn()
	user := model.User{
		Email: "momochan@email.com",
	}

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	usr, err := repository.NewRepository(db).CreateUser(user)

	if err != nil {
		log.Println(err.Error())
		return
	}

	res, _ := json.Marshal(usr)
	fmt.Println(string(res))
}

func TestGetUser(t *testing.T) {
	db, err := config.NewConn()

	if err != nil {
		log.Println(err.Error())
		return
	}

	u, err := repository.NewRepository(db).GetUser(1)

	if err != nil {
		log.Println(err.Error())
		return
	}

	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}

func TestUpdateUser(t *testing.T) {
	db, err := config.NewConn()

	if err != nil {
		log.Println(err.Error())
		return
	}

	id := 7
	user := model.User{
		Email: "momo@email.com",
	}

	u, err := repository.NewRepository(db).UserUpdate(id, user)

	if err != nil {
		log.Println(err.Error())
		return
	}

	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}

func TestGetUserWithProduct(t *testing.T) {
	db, err := config.NewConn()
	user := model.User{}

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	u, err := repository.NewRepository(db).GetUserWithProduct(user)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}