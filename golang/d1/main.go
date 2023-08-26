package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	// Load the data file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
		return
	}
	defer file.Close()

	result := make([]uint64, 0, 10)

	maxCal, err := part1(file, &result)
	if err != nil {
		fmt.Printf("error on part 1: %v", err)
		return
	}

	fmt.Println(maxCal)        //67450
	fmt.Println(part2(result)) //199357
}

// Find the max of the slide for the example is 24000
func part1(file *os.File, result *[]uint64) (uint64, error) {
	// Read the file line by line and sum until we find a blank line
	scanner := bufio.NewScanner(file)

	var counter uint64

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// After a blank line store the result and start sum from 0 again
			*result = append(*result, counter)
			counter = 0
			continue
		}

		var cal uint64
		cal, err := strconv.ParseUint(line, 10, 32)
		if err != nil {
			return 0, fmt.Errorf("error converting calories to uint64: %w", err)
		}

		counter += cal
		// Repeate the process until EOF
	}

	// Store the last one after EOF
	*result = append(*result, counter)

	if scanner.Err() != nil {
		return 0, fmt.Errorf("error during the file scan: %w", scanner.Err())
	}
	return slices.Max(*result), nil
}

// For the second part we need to find the top 3 elfs
func part2(result []uint64) uint64 {
	// Sort the list, grab the bottom 3 because it sorts ascending and sum
	var top3Cal uint64
	slices.Sort(result)
	idx := len(result)
	top3Cal += result[idx-1]
	top3Cal += result[idx-2]
	top3Cal += result[idx-3]
	return top3Cal
}
