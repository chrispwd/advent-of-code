package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	document, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	calibrateDocument(document)
}

// TODO :: Figure out solution
func calibrateDocument(doc []string) {
	var acc int = 0
	// var numStr string
	// var fdi int // first digit index
	// var ldi int // last digit index
	_ = acc
	for i, line := range doc {
		_ = i
		for j, ch := range line {
			_ = j
			s := string(ch)
			num, err := strconv.Atoi(s)
			if err != nil {
				//	fdi++
				continue
			}
			// if fdi < j {
			// 	fdi = j
			// }
			// switch num {

			// }
			fmt.Println(num)
		}
	}
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
