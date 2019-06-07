package golang

func hurdleRace(k int32, height []int32) int32 {
	var result int32

	// find max height
	var maxHeight int32 = height[0]
	for _, val := range height {
		if maxHeight < val {
			maxHeight = val
		}
	}

	result = maxHeight - k
	if result < 0 {
		result = 0
	}

	return result

}
