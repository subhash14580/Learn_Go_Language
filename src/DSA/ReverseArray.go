package main

import (
	"fmt"
)

func main() {
	fmt.Println("Reverse an array")

	arr := []int{1, 2, 3, 4, 5}

	leng := len(arr)

	result := make([]int, leng)

	for i := 0; i < leng; i++ {
		result[i] = arr[leng-1-i]
	}

	fmt.Println(result)
}
