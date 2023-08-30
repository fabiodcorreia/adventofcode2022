package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}
func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var counter uint

	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		badgeItem := findTrip(line1, line2, line3)
		counter += runeToPriority(badgeItem)
	}
	println(counter)
}

func findTrip(ruck1, ruck2, ruck3 string) rune {
	for _, item := range ruck1 {
		if strings.ContainsRune(ruck2, item) && strings.ContainsRune(ruck3, item) {
			return item
		}
	}
	return rune(0)
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var counter uint

	for scanner.Scan() {
		line := scanner.Text()
		middle := len(line) / 2
		dup := findDup(line[:middle], line[middle:])
		counter += runeToPriority(dup)
	}
	println(counter)
}

func findDup(comp1, comp2 string) rune {
	for _, item := range comp1 {
		if strings.ContainsRune(comp2, item) {
			return item
		}
	}
	return rune(0)
}

func runeToPriority(r rune) uint {
	if 'a' <= r && r <= 'z' {
		return uint(r - 'a' + 1)
	} else if 'A' <= r && r <= 'Z' {
		return uint(r - 'A' + 27)
	}
	return 0
}
