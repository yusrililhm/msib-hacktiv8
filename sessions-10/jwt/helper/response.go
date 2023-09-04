package helper

import "net/http"

type Response struct {
	Message string `json:"message"`
}

var (
	BadRequest  = http.StatusText(http.StatusBadRequest)
	Unathorized = http.StatusText(http.StatusUnauthorized)
)
