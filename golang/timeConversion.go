package golang

import (
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	// make sure lower case
	s = strings.ToLower(s)

	// if pm
	if strings.Contains(s, "pm") {
		// check if 12 pm
		if s[0:2] == "12" {
			return s[0 : len(s)-2]
		}

		temp, err := strconv.Atoi(s[0:2])
		if err != nil {
			return ""
		}

		return strconv.Itoa(temp+12) + s[2:len(s)-2]
	} else {
		temp, err := strconv.Atoi(s[0:2])
		if err != nil {
			return ""
		}

		if temp == 12 {
			return "00" + s[2:len(s)-2]
		}

	}

	return s[0 : len(s)-2]
}
