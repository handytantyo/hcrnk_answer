package golang

func birthday(s []int32, d int32, m int32) int32 {
	var counter int32
	if int(m) > len(s) {
		return counter
	}

	var maxIter int = len(s) - int(m) + 1
	for i := 0; i < maxIter; i++ {
		var sum int32
		for j := i; j < (i + int(m)); j++ {
			sum += s[j]
		}
		if sum == d {
			counter += 1
		}
	}

	return counter
}
