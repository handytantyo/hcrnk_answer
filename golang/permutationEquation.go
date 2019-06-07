package golang

func permutationEquation(p []int32) []int32 {
	result := make([]int32, len(p))

	// making map for finding map index
	mapIndex := make(map[int32]int32)
	for i := int32(0); i < int32(len(p)); i++ {
		mapIndex[p[i]] = i + 1
	}

	for i := int32(0); i < int32(len(p)); i++ {
		temp := mapIndex[mapIndex[i+1]]
		result[i] = temp
	}

	return result
}
