package main

import (
	"fmt"
)

func main() {

	courseInProg := []string{"Docker Deep Dive",
		"Docker Clustering", "Docker and Kubernetes"}

	courseCompleted := []string{"Puppet Fundamentals",
		"Go Fundamentals", "Docker Deep Dive"}

	for _, i := range courseInProg {
		for _, j := range courseCompleted {
			if i == j {
				fmt.Println("Clash found.", i, "is in both lists")
			}
		}
	}
}
