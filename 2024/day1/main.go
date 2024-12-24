package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filename = "input.txt"

func main() {

	l1, l2 := loadLists(filename)

	part1Ans := part1(l1, l2)
	part2Ans := part2(l1, l2)
	fmt.Printf("Part 1: total distance is %v\n", part1Ans)
	fmt.Printf("Part 2: total similarity score is %v\n", part2Ans)
}

func part2(l1 []int, l2 []int) int {

	score := 0
	for _, h := range l1 {
		factor := 0
		for _, k := range l2 {
			if h == k {
				factor++
			}
		}
		score += (h * factor)
	}

	return score
}

func part1(l1 []int, l2 []int) int {

	totalDistance := 0

	for i := 0; i < len(l1); i++ {
		// fmt.Printf("l1: %v l2: %v\n", l1[i], l2[i])
		diff := l1[i] - l2[i]
		if diff < 0 {
			diff = diff * -1
		}
		totalDistance += diff
	}

	return totalDistance
}

func loadLists(filename string) ([]int, []int) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var list1 []int
	var list2 []int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Printf("Error parsing line '%s': %v\n", line, err)
			continue
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Printf("Error parsing line '%s': %v\n", line, err)
			continue
		}

		list1 = append(list1, x)
		list2 = append(list2, y)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while scanning: %v\n", err)
	}

	sort.Ints(list1)
	sort.Ints(list2)
	return list1, list2
}
