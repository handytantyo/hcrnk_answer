package golang

func catAndMouse(x int32, y int32, z int32) string {
	stepCatA := x - z
	if stepCatA < 0 {
		stepCatA *= -1
	}
	stepCatB := y - z
	if stepCatB < 0 {
		stepCatB *= -1
	}

	if stepCatA < stepCatB {
		return "Cat A"
	} else if stepCatB < stepCatA {
		return "Cat B"
	} else {
		return "Mouse C"
	}

}
