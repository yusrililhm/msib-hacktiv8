package controller

import (
	"encoding/json"
	"go-http-request/entity"
	"net/http"
)

var baseUrl = "http://localhost:5000/"

func FetchUser() ([]entity.Employee, error) {
	client := &http.Client{}
	data := []entity.Employee{}

	request, err := http.NewRequest("POST", baseUrl + "/employees", nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
