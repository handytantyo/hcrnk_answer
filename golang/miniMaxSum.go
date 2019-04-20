package golang

import "fmt"

func miniMaxSum(arr []int32) {
	var sumTotal uint32
	for _, val := range arr {
		sumTotal += uint32(val)
	}
	fmt.Println(sumTotal)
	// find max and min
	var min, max uint32
	min = sumTotal - uint32(arr[0])
	max = sumTotal - uint32(arr[0])
	for _, val := range arr {
		if max < (sumTotal - uint32(val)) {
			max = sumTotal - uint32(val)
		} else if min > (sumTotal - uint32(val)) {
			min = sumTotal - uint32(val)
		}
	}

	fmt.Println(min, max)

}
