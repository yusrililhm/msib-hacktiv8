package handler

import (
	"basic-auth/middleware"
	"basic-auth/model"
	"basic-auth/util"
	"net/http"
	"strconv"
)

func ActionStudent(w http.ResponseWriter, r *http.Request)  {
	if !middleware.Auth(w, r) { return }
	if !middleware.AllowOnlyGet(w, r) { return }

	if id := r.URL.Query().Get("id"); id != "" {
		param, _ := strconv.Atoi(id)
		util.OutputJSON(w, model.SelectStudent(uint(param)))
		return
	}

	util.OutputJSON(w, model.GetStudent())
}