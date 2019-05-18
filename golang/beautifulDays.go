package golang

import (
	"math"
)

func beautifulDays(i int32, j int32, k int32) int32 {
	var result int32

	for number := i; number <= j; number++ {
		reverse := reverseNumber(number)
		total := number - reverse
		if total < 0 {
			total *= -1
		}

		if total%k == 0 {
			result++
		}
	}

	return result
}

func reverseNumber(number int32) (result int32) {
	temp := []int32{}
	for number > 0 {
		temp = append(temp, number%10)
		number /= 10
	}

	// calculate the reverse slice to reverse int
	power := float64(len(temp) - 1)
	for _, number := range temp {
		result = result + (number * int32(math.Pow(10, power)))
		power--
	}

	return
}
