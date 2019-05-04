package golang

func pickingNumbers(a []int32) int32 {
	mapA := make(map[int32]int32)
	for _, number := range a {
		// insert into map
		if _, ok := mapA[number]; ok {
			mapA[number] += 1
		} else {
			mapA[number] = 1
		}
	}

	var result, tempResult int32
	for k, v := range mapA {
		if mapA[k-1] >= mapA[k+1] {
			tempResult = mapA[k-1]
		} else {
			tempResult = mapA[k+1]
		}
		tempResult += v

		if result < tempResult {
			result = tempResult
		}
		tempResult = 0
	}

	return result
}
