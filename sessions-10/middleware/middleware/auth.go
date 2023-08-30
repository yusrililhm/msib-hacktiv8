package middleware

import "net/http"

var USERNAME, PASSWORD = "USER", "RAHASIA"

func Auth(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
	
		if !ok {
			w.Write([]byte("somethng wrong"))
			return 
		}
	
		if !(username == USERNAME && password == PASSWORD) {
			w.Write([]byte("username dan password salah"))
			return 
		}
	
		handler.ServeHTTP(w, r)
	})
}

func AllowOnlyGet(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("only get"))
			return
		}
		
		handler.ServeHTTP(w, r)
	})
}