package main

import "fmt"

func main() {
	// Variables to store course rankings

	if firstRank, secondRank := 685, 614; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better" +
			" than the second course.")
		if !isAcceptableRank(firstRank) {
			fmt.Println("You may want to consider getting another job... :-D")
		}
	} else if firstRank > secondRank {
		fmt.Println("\nOh dear... your first " +
			"course must be doing abysmally.")
		if !isAcceptableRank(firstRank) {
			fmt.Println("You may want to consider getting another job... :-D")
		}
	} else {
		fmt.Println("\nBoth courses are either " +
			"the same, or something weird is going on.")
	}
}

func isAcceptableRank(input int) bool {
	if input > 100 {
		return false
	}
	return true
}
