package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
		return
	}
	defer file.Close()

	for i := 0; i < b.N; i++ {
		result := make([]uint64, 0, 10)
		part1(file, &result)
	}
}

func BenchmarkPart2(b *testing.B) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
		return
	}
	defer file.Close()

	result := make([]uint64, 0, 10)
	part1(file, &result)

	for i := 0; i < b.N; i++ {
		part2(result)
	}

}

//go test -bench=. -count 1 -benchmem
