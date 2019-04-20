package golang

import "fmt"

func staircase(n int32) {
	for i := 0; i < int(n); i++ {
		// space
		for space := (int(n) - 1); space > i; space-- {
			fmt.Printf(" ")
		}

		for hash := 0; hash <= i; hash++ {
			fmt.Printf("#")
		}

		fmt.Println()
	}

}
