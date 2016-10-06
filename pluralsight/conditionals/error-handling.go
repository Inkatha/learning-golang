package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("C:\\Users\\mink\\temp\\test.txt")

	if err != nil {
		fmt.Println("Error returned was:", err)
	}
	fmt.Println("Error returned was:", err)
}
