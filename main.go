package main

import (
	"fmt"

	"/src"
)

func main() {
	// Input: Year for which you want to find the day of the week
	year := 2023

	// Assuming a specific date, for example, January 1st of the given year
	month := 1
	day := 1

	// Calculate the day of the week using Zeller's Congruence
	result := src.ZellersCongruence(day, month, year)

	// Print the result
	fmt.Printf("The day of the week on %d-%02d-%02d is %s.\n", year, month, day, result)
}
