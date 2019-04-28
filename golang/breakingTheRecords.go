package golang

func breakingRecords(scores []int32) []int32 {
	var minScore, maxScore int32
	var countMin, countMax int32

	minScore, maxScore = scores[0], scores[0]

	for _, score := range scores {
		if minScore > score {
			countMin++
			minScore = score
		} else if maxScore < score {
			countMax++
			maxScore = score
		}
	}

	return []int32{countMax, countMin}
}
