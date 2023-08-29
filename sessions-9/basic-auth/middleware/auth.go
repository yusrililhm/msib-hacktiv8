package middleware

import "net/http"

var USERNAME, PASSWORD = "USER", "RAHASIA"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte("somethng wrong"))
		return false
	}

	if !(username == USERNAME && password == PASSWORD) {
		w.Write([]byte("username dan password salah"))
		return false
	}

	return true
}

func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("only get"))
		return false
	}

	return true
}