package test

import (
	"encoding/json"
	"fmt"
	"go-postgresql/config"
	"go-postgresql/model"
	"go-postgresql/repository"
	"log"
	"testing"
)

func TestConn(t *testing.T) {
	db, _ := config.NewConn()

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Println("connect success")
}

func TestCreateEmployee(t *testing.T) {
	db, err := config.NewConn()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	employee := model.Employees{
		FullName: "Kim Dahyun",
		Email:    "dahyun@email.com",
		Age:      25,
		Division: "Rapper",
	}
	repository, err := repository.NewEmployeesRepository(db).CreateEmployee(employee)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	res, _ := json.Marshal(repository)

	fmt.Println(string(res))
}

func TestGetEmployee(t *testing.T) {
	db, err := config.NewConn()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	employees, err := repository.NewEmployeesRepository(db).GetEmployee()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	res, _ := json.Marshal(employees)
	fmt.Println(string(res))
}

func TestUpdate(t *testing.T) {
	db, err := config.NewConn()
	employee := model.Employees{
		Id: 25,
		FullName: "Mina",
		Email: "minachan@email.com",
		Age: 26,
		Division: "Vocal",
	}

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	s, err := repository.NewEmployeesRepository(db).UpdateEmployee(employee)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println(s)
}

func TestDelete(t *testing.T) {
	db, err := config.NewConn()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	s, err := repository.NewEmployeesRepository(db).DeleteEmployee(5)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println(s)
}