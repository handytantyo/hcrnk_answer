package golang

func viralAdvertising(n int32) int32 {
	if n <= 1 {
		return 2
	}

	var shared, liked, cumulative int32
	shared, liked, cumulative = 5, 2, 2
	for i := int32(1); i < n; i++ {
		shared = liked * 3
		liked = shared / 2
		cumulative += liked
	}
	return cumulative
}
