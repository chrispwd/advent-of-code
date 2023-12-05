package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	document, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println(calibrateDocument(document))
}

// TODO :: Finish Part 2
func calibrateDocument(doc []string) int {
	var acc int = 0
	_ = acc
	for i, line := range doc {
		_ = i
		re := regexp.MustCompile(`([0-9]){1}`)
		nums := re.FindAllString(line, -1)
		num, err := strconv.Atoi(fmt.Sprintf("%v%v", nums[0], nums[len(nums)-1]))
		if err != nil {
			continue
		}
		acc += num
	}
	return acc
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
