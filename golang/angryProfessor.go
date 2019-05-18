package golang

import "fmt"

func angryProfessor(k int32, a []int32) string {
	var result string

	// calculate the on time student
	var ontimeStudent int32
	for _, student := range a {
		if student <= 0 {
			ontimeStudent++
		}
	}

	// compare the threshold
	if ontimeStudent >= k {
		result = fmt.Sprintf("NO")
	} else {
		result = fmt.Sprintf("YES")
	}

	return result
}
