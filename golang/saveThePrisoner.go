package golang

func saveThePrisoner(n int32, m int32, s int32) int32 {

	result := (s + m - 1) % n

	if result == 0 {
		result = n
	}

	return result

}
