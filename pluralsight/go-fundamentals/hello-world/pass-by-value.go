package main

import (
    "fmt"
);

func main() {
    
    name := "Malik"
    course := "Go-Fundamentals"

    fmt.Println("\nHi", name, "you're currently watching", course)

    changeCourse(&course)

    fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {

    *course = "First look: Native Docker Clustering"

    fmt.Println("\nTrying to change your course to", *course)

    return *course
}