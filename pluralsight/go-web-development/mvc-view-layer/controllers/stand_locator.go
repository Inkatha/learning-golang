package controllers

import (
	"encoding/json"
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/controllers/util"
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/viewmodels"
	"net/http"
	"text/template"
)

type standLocatorController struct {
	template *template.Template
}

func (standLocator *standLocatorController) get(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocator()
	responseWriter.Header().Add("Content-Type", "text/html")
	standLocator.template.Execute(responseWriter, vm)
}

func (standLocator *standLocatorController) apiSearch(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocations()
	responseWriter.Header().Add("Content-Type", "application/json")
	data, err := json.Marshal(vm)

	if err == nil {
		responseWriter.Write(data)
	} else {
		responseWriter.WriteHeader(404)
	}
}
