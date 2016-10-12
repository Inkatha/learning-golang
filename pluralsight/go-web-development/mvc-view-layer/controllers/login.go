package controllers

import (
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/viewmodels"
	"net/http"
	"text/template"
)

type loginController struct {
	template *template.Template
}

func (login *loginController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetLogin()
	w.Header().Add("Context-Type", "text/html")

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")

		member, err := models.GetMember(email, password)
	}

	login.template.Execute(w, vm)
}
