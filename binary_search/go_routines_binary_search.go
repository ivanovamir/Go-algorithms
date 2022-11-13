package main

import (
	"fmt"
	"time"
)

type BinarySearchStruct struct {
	Index    int
	Attempts int
}

type SimpleSearchStruct struct {
	Index    int
	Attempts int
}

func main() {

	t := time.Now()

	ch1 := make(chan BinarySearchStruct)
	ch2 := make(chan SimpleSearchStruct)
	var binarySearchResult, binarySearchAttempts, simpleSearchResult, simpleSearchAttempts int

	var array []int
	length := 10000000 //length of array
	searchingNumber := 10000000
	for i := 0; i <= length; i++ {
		array = append(array, i)
	}

	go BinarySearchGoRoutine(array, searchingNumber, ch1)
	go SimpleSearchGoRoutine(array, searchingNumber, ch2)

	for x := range ch1 {
		binarySearchResult = x.Index
		binarySearchAttempts = x.Attempts
	}
	for x := range ch2 {
		simpleSearchResult = x.Index
		simpleSearchAttempts = x.Attempts
	}

	if binarySearchResult == 0 && binarySearchAttempts == 0 || simpleSearchResult == 0 && simpleSearchAttempts == 0 {
		fmt.Printf("Number: %d not found", searchingNumber)
		return
	}

	//This benchmark is not unambiguously accurate,
	//a full-fledged benchmark is implemented in the
	//first two files of this directory

	fmt.Printf("Binary search results:\nNumber: %d | total attempts: %d | Latency: %d ns\n", array[binarySearchResult], binarySearchAttempts, time.Since(t).Nanoseconds())

	fmt.Printf("Simple search results:\nNumber: %d | total attempts: %d | Latency: %d ns\n", array[simpleSearchResult], simpleSearchAttempts, time.Since(t).Nanoseconds())
}

func BinarySearchGoRoutine(array []int, number int, ch1 chan BinarySearchStruct) {
	low := 0               //first index
	high := len(array) - 1 //last index
	attempts := 0          //attempts count
	result := new(BinarySearchStruct)
	for low <= high {
		mid := (low + high) / 2
		guess := array[mid]
		if guess == number {
			result.Index = mid
			result.Attempts = attempts
			ch1 <- *result
		}
		if guess > number {
			high = mid - 1
		} else {
			low = mid + 1
		}
		attempts += 1
	}
	close(ch1)
}

func SimpleSearchGoRoutine(array []int, number int, ch2 chan SimpleSearchStruct) {
	attempts := 0 //attempts count
	result := new(SimpleSearchStruct)
	for _, y := range array {
		if array[y] == number {
			result.Index = array[y]
			result.Attempts = attempts
			ch2 <- *result
		}
		attempts += 1
	}
	close(ch2)
}
