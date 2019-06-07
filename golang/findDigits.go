package golang

func findDigits(n int32) int32 {
	var counter, tempN int32
	tempN = n

	for tempN > 0 {
		mod := tempN % 10
		tempN = tempN / 10

		if mod != 0 && (n%mod == 0) {
			counter++
		}
	}

	return counter

}
