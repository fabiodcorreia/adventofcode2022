package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var counter uint32

	for scanner.Scan() {
		line := scanner.Text()
		elfs := strings.Split(line, ",")

		elf1 := strings.Split(elfs[0], "-")
		elf2 := strings.Split(elfs[1], "-")

		elf1s1, _ := strconv.ParseUint(elf1[0], 10, 32)
		elf1s2, _ := strconv.ParseUint(elf1[1], 10, 32)

		elf2s1, _ := strconv.ParseUint(elf2[0], 10, 32)
		elf2s2, _ := strconv.ParseUint(elf2[1], 10, 32)

		if elf1s2 >= elf2s1 && elf1s1 <= elf2s2 {
			counter++
		}
	}

	println(counter)
}
func part1() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var counter uint32

	for scanner.Scan() {
		line := scanner.Text()
		elfs := strings.Split(line, ",")

		elf1 := strings.Split(elfs[0], "-")
		elf2 := strings.Split(elfs[1], "-")

		elf1s1, _ := strconv.ParseUint(elf1[0], 10, 32)
		elf1s2, _ := strconv.ParseUint(elf1[1], 10, 32)

		elf2s1, _ := strconv.ParseUint(elf2[0], 10, 32)
		elf2s2, _ := strconv.ParseUint(elf2[1], 10, 32)

		if elf1s2 >= elf2s1 {
			if elf1s1 >= elf2s1 && elf1s2 <= elf2s2 {
				counter++
				continue
			}
			if elf2s1 >= elf1s1 && elf2s2 <= elf1s2 {
				counter++
				continue
			}
		}
	}

	println(counter)
}
