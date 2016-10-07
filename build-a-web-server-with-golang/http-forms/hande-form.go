package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parse URL parameters passed, then parse the resposne packet from body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFileS("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// login parse of log in
		fmt.Println("Username:", r.Form["username"])
		fmt.Println("Password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayHelloName) // setting router rule
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // Setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
