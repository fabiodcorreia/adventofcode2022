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
	// Read the file line by line and sum until we find a blank line
	scanner := bufio.NewScanner(file)

	var counter uint64
	result := make([]uint64, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// After a blank line store the result and start sum from 0 again
			result = append(result, counter)
			counter = 0
			continue
		}

		var cal uint64
		cal, err = strconv.ParseUint(line, 10, 32)
		if err != nil {
			fmt.Printf("error converting calories to uint64: %v", err)
			return
		}

		counter += cal
		// Repeate the process until EOF
	}

	// Store the last one after EOF
	result = append(result, counter)

	if scanner.Err() != nil {
		fmt.Printf("error during the file scan: %v", err)
		return
	}

	// Find the max of the slide for the example is 24000
	fmt.Println(slices.Max(result)) //67450

	// For the second part we need to find the top 3 elfs

	// Sort the list, grab the bottom 3 because it sorts ascending and sum
	slices.Sort(result)
	var top3Cal uint64
	top3Cal += result[len(result)-1]
	top3Cal += result[len(result)-2]
	top3Cal += result[len(result)-3]

	fmt.Println(top3Cal) //199357
}
