package model

import "go-web-server/entity"

var Employees = []entity.Employee{
	{
		Id:       1,
		Name:     "Momo",
		Age:      26,
		Division: "Development and Operations",
	},
	{
		Id:       2,
		Name:     "Sana",
		Age:      26,
		Division: "Multimedia",
	},
	{
		Id:       3,
		Name:     "Dahyun",
		Age:      25,
		Division: "IT Develepoments",
	},
}
