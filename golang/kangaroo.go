package golang

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v2 >= v1 {
		return "NO"
	}

	multiple := int32(1)

	for {
		kangarooAPosition := x1 + (v1 * multiple)
		kangarooBPosition := x2 + (v2 * multiple)

		if kangarooAPosition > kangarooBPosition {
			return "NO"
		} else if kangarooAPosition == kangarooBPosition {
			break
		}
		multiple++
	}

	return "YES"

}
