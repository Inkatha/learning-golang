package main

import "fmt"

func main() {

	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	// Ways to declare a variable of type struct.
	// var DockerDeepDive courseMeta
	// DockerDeepDive := new(courseMeta)

	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
}
