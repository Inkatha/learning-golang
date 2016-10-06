package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 1, 4)

	for i := 0; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity is: %d", cap(mySlice))
	}
}
