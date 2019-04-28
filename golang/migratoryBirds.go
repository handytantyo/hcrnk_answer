package golang

func migratoryBirds(arr []int32) int32 {
	mapBirds := make(map[int32]int32)

	// count every kind of bird
	for _, bird := range arr {
		if mapBirds[bird] == 0 {
			mapBirds[bird] = 1
		} else {
			mapBirds[bird] += 1
		}
	}

	// find max number bird
	var maxKindOfBird int32
	var result int32
	for key, value := range mapBirds {
		if maxKindOfBird < value {
			result = key
			maxKindOfBird = value
		} else if maxKindOfBird == value {
			if result > key {
				result = key
			}
		}
	}

	return result
}
