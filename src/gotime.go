package src

import (
	"fmt"
	"time"
)

// TimeCongruence is a function that takes a time.Time object and returns the day of the week based on a given year
func TimeCongruence() {
	year := 2022 // replace with the year you want

	start := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.UTC)

	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		fmt.Printf("%v %v\n", d.Format("2006-01-02"), d.Weekday())
	}
}
