package golang

import "fmt"

// source https://www.geeksforgeeks.org/factorial-large-number/
func extraLongFactorials(n int32) {
	array := make([]int32, 500)

	array[0] = 1
	arraySize := int32(1)

	for i := int32(2); i <= n; i++ {
		arraySize = multiply(i, array, arraySize)
	}

	for i := arraySize - 1; i >= 0; i-- {
		fmt.Printf("%d", array[i])
	}
	fmt.Println()

}

func multiply(x int32, array []int32, arraySize int32) int32 {
	var carry int32

	for i := int32(0); i < arraySize; i++ {
		prod := array[i]*x + carry

		array[i] = prod % 10
		carry = prod / 10
	}

	for carry != 0 {
		array[arraySize] = carry % 10
		carry /= 10
		arraySize++
	}

	return arraySize
}
