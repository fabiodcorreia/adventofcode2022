package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalScore1 uint64
	var totalScore2 uint64

	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), " ")
		a := actionToScore1(lineSplit[0])
		b := actionToScore1(lineSplit[1])

		totalScore1 += calScore(a, b)
		bAction := actionToScore2(a, b)
		println(a, bAction)
		totalScore2 += calScore(a, bAction)
	}

	println(totalScore1)
	println(totalScore2)

}

func actionToScore1(action string) uint64 {
	if action == "A" || action == "X" {
		return 1
	}
	if action == "B" || action == "Y" {
		return 2
	}
	if action == "C" || action == "Z" {
		return 3
	}
	return 0
}

func actionToScore2(a, b uint64) uint64 {
	// If b is 2 we need to draw so we return the same as a
	if b == 2 {
		return a
	}

	// if b is 3 we need to lose
	if b == 1 {
		switch a {
		case 1:
			return 3
		case 2:
			return 1
		default:
			return 2
		}
	}

	// At the point the only option is to win
	// if a is 1 or 2 we return the next value
	if a < 3 {
		return a + 1
	}

	// if a is 3 (max value) we return 1 that wins 3
	return 1
}

func calScore(a, b uint64) uint64 {
	if a == b {
		return b + 3
	}

	if (a == 1 && b == 3) || (a == 2 && b == 1) || (a == 3 && b == 2) {
		return b + 0
	}

	return b + 6
}
