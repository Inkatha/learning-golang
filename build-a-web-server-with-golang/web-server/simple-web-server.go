package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // Parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print from information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // Send data to client side
}

func main() {
	http.HandleFunc("/", sayHelloName)       // Set router
	err := http.ListenAndServe(":9090", nil) // Set Listen Port
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
