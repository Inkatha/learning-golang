package main

import (
    "fmt"
    "reflect"
);

func main() {
    name := "Malik" // Name of the subscriber   

    module := 3.2   // Current position in the course

    pointer := &module 

    fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
    fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))  
    fmt.Println("Memory address of *module* varible is", 
        pointer, "and the value of *module* is", *pointer);
}