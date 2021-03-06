package controllers

import (
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/viewmodels"
	"net/http"
	"text/template"
)

type homeController struct {
	template *template.Template
}

func (home *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()
	w.Header().Add("Context-Type", "text/html")
	home.template.Execute(w, vm)
}
