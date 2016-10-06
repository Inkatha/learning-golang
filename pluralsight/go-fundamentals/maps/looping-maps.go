package main

import (
	"fmt"
)

func main() {

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6}

	for key, value := range testMap {
		fmt.Printf("\n Key is: %v Value is %v\n", key, value)
	}

	fmt.Println(testMap["C"])
	testMap["C"] = 100
	fmt.Println(testMap["C"])

	delete(testMap, "F")
	fmt.Println(testMap)
}
