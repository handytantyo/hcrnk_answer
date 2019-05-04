package golang

import (
	"fmt"
)

func bonAppetit(bill []int32, k int32, b int32) {
	// calculate total sharing
	var totalSharing int32
	for index, b := range bill {
		if index != int(k) {
			totalSharing += b
		}
	}

	// get each person that have to pay
	totalSharing = totalSharing / 2
	result := b - totalSharing
	if result == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(result)
	}

}
