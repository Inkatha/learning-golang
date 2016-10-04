package main

import (
    "fmt"
    "reflect"
);

var (
    name    string
    course  string
    module  float64        
)

func main() {
    fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
    fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))

    make  := "Dodge"     // Make of the car
    model := "Charger"   // Model of the car
    year  := 2010        // Current Year of the Model

    fmt.Println("The make is", make, "the model is", model, "The year is", year)

    a := 10.0000000
    b := 3

    c := int(a) + b

    fmt.Println("\nC has a value:", c, "and is of type", reflect.TypeOf(c))
}