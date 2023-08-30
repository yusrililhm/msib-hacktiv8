package model

import "go-middleware/entity"



var students = []*entity.Student{
	{
		Id: 1,
		Name: "mina",
	},
	{
		Id: 2,
		Name: "sana",
	},
	{
		Id: 3,
		Name: "momo",
	},
}

func GetStudent() []*entity.Student {
	return students
}

func SelectStudent(id uint) *entity.Student {
	for _, v := range students {
		if v.Id == id {
			return v
		}
	}
	return nil
}