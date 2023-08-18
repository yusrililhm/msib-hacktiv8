package test

import (
	"log"
	"reflect"
	"testing"
)

type person struct {
	Name, Address string
}

func (p *person) getPersonInfo() {
	reflectValue := reflect.ValueOf(p)

	// validasi
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	reflectType := reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		log.Println("name\t:", reflectType.Field(i).Name)
		log.Println("type\t:", reflectType.Field(i).Type)
		log.Println("value\t:", reflectValue.Field(i).Interface())
	}
}

func (p *person) SetName(name string) {
	p.Name = name
}

func TestReflect(t *testing.T) {
	name := "momo"
	reflectValue := reflect.ValueOf(name)
	log.Println("tipe data name adalah", reflectValue.Type())

	if reflect.String == reflectValue.Kind() {
		log.Println("value: ", reflectValue.String())
	}
}

func TestAccessVlueInterfaceKosong(t *testing.T) {
	number := 10
	v := reflect.ValueOf(number)

	log.Println("tipe\t:", v.Type())
	log.Println("value\t:", v.Interface())
	log.Println("nilai\t:", v.Interface().(int))
}

func TestAccessInformationObject(t *testing.T) {
	person := &person{Name: "momo", Address: "japan"}
	person.getPersonInfo()
}

func TestReflectMethod(t *testing.T) {
	person := &person{Name: "momo", Address: "japan"}

	log.Println("before: ", person.Name)
	reflectValue := reflect.ValueOf(person)

	reflectMethod := reflectValue.MethodByName("SetName")
	reflectMethod.Call([]reflect.Value{
		reflect.ValueOf("dahyun"),
	})

	log.Println("after:", person.Name)
}
