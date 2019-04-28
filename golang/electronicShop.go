package golang

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var maximumCost int32 = -1

	for _, keyboard := range keyboards {
		for _, drive := range drives {
			spend := keyboard + drive
			if b >= spend {
				if maximumCost < spend {
					maximumCost = spend
				}
			}
		}
	}

	return maximumCost
}
