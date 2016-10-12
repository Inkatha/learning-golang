package controllers

import (
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/converters"
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/models"
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/viewmodels"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

type categoriesController struct {
	template *template.Template
}

func (categoryCont *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	categories := models.GetCategories()

	categoriesVM := []viewmodels.Category{}
	for _, category := range categories {
		categoriesVM = append(categoriesVM, converters.ConvertCategoryToViewModel(category))
	}

	vm := viewmodels.GetCategories()
	vm.Categories = categoriesVM

	w.Header().Add("Context-Type", "text/html")
	categoryCont.template.Execute(w, vm)
}

type categoryController struct {
	template *template.Template
}

func (category *categoryController) get(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)

	if err == nil {
		vm := viewmodels.GetProducts(id)
		w.Header().Add("Context-Type", "text/html")
		category.template.Execute(w, vm)
	} else {
		w.WriteHeader(404)
	}
}
