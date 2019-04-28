package golang

import "fmt"

func dayOfProgrammer(year int32) string {
	dayOfMonths := []int32{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	days := int32(256)
	months := 0

	// except for 1918
	if year == 1918 {
		dayOfMonths[1] = 15
		goto calculate
	}

	// in normal condition
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) || year == 1700 || year == 1800 || year == 1900 {
		// leap year
		dayOfMonths[1] = 29
	}

	// calculate the day
calculate:
	for i := 0; i < 12; i++ {
		days -= dayOfMonths[i]

		if days <= 0 {
			days += dayOfMonths[i]
			months = i + 1
			break
		}
	}

	// convert it into string
	var stringDays, stringMonths string
	stringDays = fmt.Sprintf("%d", days)
	stringMonths = fmt.Sprintf("%d", months)
	if days < 10 {
		stringDays = "0" + stringDays
	}
	if months < 10 {
		stringMonths = "0" + stringMonths
	}

	return fmt.Sprintf("%s.%s.%d", stringDays, stringMonths, year)

}
