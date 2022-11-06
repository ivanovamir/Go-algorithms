package main

import "testing"

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(length, searchingNumber)
	}
}

func BenchmarkSimpleSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleSearch(length, searchingNumber)
	}
}
