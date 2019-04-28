package golang

func getTotalX(a []int32, b []int32) int32 {
	// find the maximum number of a
	maxNumberA := a[0]
	for _, number := range a {
		if maxNumberA < number {
			maxNumberA = number
		}
	}

	// find the minimum number of slice b
	minNumber := b[0]
	for _, number := range b {
		if minNumber > number {
			minNumber = number
		}
	}

	// make sure if maxNumberA always smaller than minNumber
	if maxNumberA > minNumber {
		return int32(0)
	}

	// find all number less than equal of minimum number
	KPK := findKPK(a)
	step1 := make([]int32, 0)
	var Factors int32 = KPK
	for {
		if Factors <= minNumber {
			step1 = append(step1, Factors)
			Factors += KPK
		} else {
			break
		}
	}

	// find all number that is factor from slice b
	step2 := make([]int32, 0)
	var checker int
	for _, factor := range step1 {
		checker = 0
		for _, numberB := range b {
			if numberB%factor != 0 {
				break
			} else if numberB%factor == 0 {
				checker++
			}
		}
		if checker == len(b) {
			step2 = append(step2, factor)
		}
	}

	return int32(len(step2))
}

func findKPK(numbers []int32) int32 {
	var maxNumber int32 = numbers[0]
	for _, number := range numbers {
		if maxNumber < number {
			maxNumber = number
		}
	}

	var multiplier int32 = 1
	var KPK int32 = maxNumber
	var checker bool
	for {
		for i := 0; i < len(numbers); i++ {
			if KPK%numbers[i] != 0 {
				break
			} else if KPK%numbers[i] == 0 && i == (len(numbers)-1) {
				checker = true
			}
		}

		if checker == true {
			break
		}

		multiplier++
		KPK = maxNumber * multiplier
	}

	return KPK
}
