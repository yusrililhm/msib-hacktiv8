package util

import (
	"encoding/json"
	"net/http"
)

func OutputJSON(w http.ResponseWriter, o any)  {
	res, err := json.Marshal(o)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}