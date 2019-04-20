package golang

func birthdayCakeCandles(ar []int32) int32 {
	// find max
	var count int32
	var max int32 = ar[0]
	for _, val := range ar {
		if max < val {
			max = val
			count = 0
		}

		if max == val {
			count++
		}
	}

	return count
}
