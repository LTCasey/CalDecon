package src

func ZellersCongruence(day, month, year int) string {
	if month < 3 {
		month += 12
		year--
	}

	k := year % 100
	j := year / 100

	// Zeller's Congruence formula
	dayOfWeek := (day + ((13 * (month + 1)) / 5) + k + (k / 4) + (j / 4) - (2 * j)) % 7

	// Adjust for Sunday (0-based) to Saturday (6-based)
	if dayOfWeek < 0 {
		dayOfWeek += 7
	}

	// Map the result to the corresponding day
	days := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	return days[dayOfWeek]
}
