package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1, 7}
	selectionSort(arr)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] <= arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
