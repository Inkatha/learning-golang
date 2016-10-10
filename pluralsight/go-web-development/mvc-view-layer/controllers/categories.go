package controllers

import (
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/viewmodels"
	"net/http"
	"text/template"
)

type categoriesController struct {
	template *template.Template
}

func (categories *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetCategories()
	w.Header().Add("Context-Type", "text/html")
	categories.template.Execute(w, vm)
}
