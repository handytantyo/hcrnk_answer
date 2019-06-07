package golang

func jumpingOnClouds(c []int32, k int32) int32 {
	health := int32(100)

	for i := int32(k); i < int32(len(c))+k; i += k {
		if i >= int32(len(c)) {
			health = health - 1 - (c[0] * 2)
			break

		} else {
			health = health - 1 - (c[i] * 2)
		}

	}

	return health

}
