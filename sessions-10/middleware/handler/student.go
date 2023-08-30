package handler

import (
	"go-middleware/model"
	"go-middleware/util"
	"net/http"
	"strconv"
)

func ActionStudent(w http.ResponseWriter, r *http.Request)  {

	if id := r.URL.Query().Get("id"); id != "" {
		param, _ := strconv.Atoi(id)
		util.OutputJSON(w, model.SelectStudent(uint(param)))
		return
	}

	util.OutputJSON(w, model.GetStudent())
}