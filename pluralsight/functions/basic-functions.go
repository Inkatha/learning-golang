package main

import (
	"fmt"
	"strings"
)

func main() {

	module := "function basics"
	author := "nigel pulton"

	fmt.Println(converter(module, author))
}

func converter(s1, s2 string) (s1, s2 string) {
	s1 = strings.Title(s1)
	s2 = strings.ToUpper(s2)

	return s1, s2
}
