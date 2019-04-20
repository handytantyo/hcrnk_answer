package golang

import "fmt"

func plusMinus(arr []int32) {
	penyebut := len(arr)
	var positif, negatif, zero int32

	for _, val := range arr {
		if val > 0 {
			positif++
		} else if val < 0 {
			negatif++
		} else if val == 0 {
			zero++
		}
	}

	// print all result
	fmt.Printf("%.4f\n", float32(positif)/float32(penyebut))
	fmt.Printf("%.4f\n", float32(negatif)/float32(penyebut))
	fmt.Printf("%.4f\n", float32(zero)/float32(penyebut))

}
