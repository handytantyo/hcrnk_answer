package golang

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var countApples, countOranges int32

	for _, apple := range apples {
		distance := int32(a + apple)
		if distance >= s && distance <= t {
			countApples++
		}
	}

	for _, orange := range oranges {
		distance := int32(b + orange)
		if distance >= s && distance <= t {
			countOranges++
		}
	}

	fmt.Println(countApples)
	fmt.Println(countOranges)

}
