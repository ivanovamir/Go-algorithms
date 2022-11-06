package main

import (
	"fmt"
)

var (
	length          = 1000000 //length of array
	searchingNumber = 1000000
)

func main() {

	fmt.Printf("Attempts: %d\n", BinarySearch(length, searchingNumber))
	fmt.Printf("Attempts: %d\n", SimpleSearch(length, searchingNumber))
}

func BinarySearch(length, searchingNumber int) int {
	var array []int
	for i := 0; i <= length; i++ {
		array = append(array, i)
	}
	low := 0               //first index
	high := len(array) - 1 //last index
	attempts := 0          //attempts count
	for low <= high {
		mid := (low + high) / 2
		guess := array[mid]
		if guess == searchingNumber {
			return attempts
		}
		if guess > searchingNumber {
			high = mid - 1
		} else {
			low = mid + 1
		}
		attempts += 1
	}
	return 0
}

func SimpleSearch(length, searchingNumber int) int {
	var array []int
	for i := 0; i <= length; i++ {
		array = append(array, i)
	}
	attempts := 0 //attempts count
	for _, y := range array {
		if array[y] == searchingNumber {
			return attempts
		}
		attempts += 1
	}
	return 0
}
