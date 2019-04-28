package golang

func pageCount(n int32, p int32) int32 {
	var firstPage int32 = 1
	lastPage := (n + 1) / 2
	if (n+1)%2 == 1 {
		lastPage += 1
	}
	targetPage := (p + 1) / 2
	if (p+1)%2 == 1 {
		targetPage += 1
	}

	var minimumSteps int32
	minimumSteps = targetPage - firstPage
	if minimumSteps > (lastPage - targetPage) {
		minimumSteps = lastPage - targetPage
	}

	return minimumSteps

}
