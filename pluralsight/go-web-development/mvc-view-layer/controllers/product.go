package controllers

import (
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/viewmodels"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

type productController struct {
	template *template.Template
}

func (prod *productController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)

	if err == nil {
		vm := viewmodels.GetProduct(id)
		w.Header().Add("Context-Type", "text/html")
		prod.template.Execute(w, vm)
	} else {
		w.WriteHeader(404)
	}
}
