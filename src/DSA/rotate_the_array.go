package main

import "fmt"

func main() {

	fmt.Println("Rotation of the array upto n times")

	rotationUptoTimes := 2

	arr := []int{1, 2, 3, 4, 5, 6, 7}

	lengt := len(arr)

	arr = reverse(arr, 0, rotationUptoTimes)
	arr = reverse(arr, rotationUptoTimes, lengt)

	arr = reverse(arr, 0, lengt)

}

func reverse(arr []int, start int, end int) []int {

	var temp int
	j := end - 1
	for i := start; i < j; i++ {
		temp = arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		j--
	}

	fmt.Println(arr)
	return arr
}
