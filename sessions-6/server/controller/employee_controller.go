package controller

import (
	"encoding/json"
	"go-web-server/entity"
	"go-web-server/model"
	"net/http"
)

func GetEmployees(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(model.Employees)
}

func AddEmployee(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	employee := entity.Employee{}

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid body request", http.StatusBadRequest)
		return
	}

	model.Employees = append(model.Employees, employee)
	json.NewEncoder(w).Encode(employee)
}
