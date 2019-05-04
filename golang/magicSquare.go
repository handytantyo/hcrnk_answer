package golang

func formingMagicSquare(s [][]int32) int32 {
	var minimumDelta int32
	combinationMagicNumbers := [][][]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}

	// do looping to find the minimum delta
	var tempDelta int32
	for checker, combinationMagicNumber := range combinationMagicNumbers {
		for i, row := range combinationMagicNumber {
			for j, column := range row {
				temp := column - s[i][j]
				if temp < 0 {
					temp *= -1
				}

				tempDelta += temp
			}
		}

		if checker == 0 {
			minimumDelta = tempDelta
		} else {
			if minimumDelta > tempDelta {
				minimumDelta = tempDelta
			}
		}

		tempDelta = 0
	}

	return minimumDelta

}
