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
	fmt.Println(asciiArt())
	fmt.Println("Calibration Reading:", calibrateDocument(document))
}

func calibrateDocument(doc []string) int {
	var acc int = 0
	_ = acc
	for i, line := range doc {
		_ = i
		strNums := unspell(line)
		intNums, err := strconv.Atoi(fmt.Sprintf("%v%v", strNums[0], strNums[len(strNums)-1]))
		if err != nil {
			continue
		}
		acc += intNums
	}
	return acc
}

func unspell(line string) []string {
	var raw []string
	re := regexp.MustCompile(`([0-9]|one|two|three|four|five|six|seven|eight|nine){1}`)
	for true {
		var match string = re.FindString(line)
		var loc []int = re.FindStringIndex(line)

		if loc == nil {
			break
		}
		switch match {
		case "one":
			raw = append(raw, "1")
			line = line[loc[1]-1:]
		case "two":
			raw = append(raw, "2")
			line = line[loc[1]-1:]
		case "three":
			raw = append(raw, "3")
			line = line[loc[1]-1:]
		case "four":
			raw = append(raw, "4")
			line = line[loc[1]:]
		case "five":
			raw = append(raw, "5")
			line = line[loc[1]-1:]
		case "six":
			raw = append(raw, "6")
			line = line[loc[1]:]
		case "seven":
			raw = append(raw, "7")
			line = line[loc[1]-1:]
		case "eight":
			raw = append(raw, "8")
			line = line[loc[1]-1:]
		case "nine":
			raw = append(raw, "9")
			line = line[loc[1]-1:]
		default:
			raw = append(raw, match)
			line = line[loc[1]:]
		}
	}
	return raw
}

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

func asciiArt() string {
	elf := `                           '-.
                         ,'   |
                       ,'     |
                     ,'       :
                   ,<\        \
                 ,' //         .
               ,' ,'/          :
             ,'  /,'           |
            /._,'/             |
           /.__,'              |
          /                    |
       .-'                     |
       '.                      |
         |'-._          ,      |
         ; -. '-.     ,'|      :
        / |@)\  |   ,', ;      '
       /        |  : /| |     /
      /         |  |: | |     :
    ,'          |  |'-' ;     |
   |   .      ,''._|   /-._   '
   '-.\ )   ,'       -'    -,-'
      |    :                :
      ;._                   |
      \__'.         )       |
      /_         ,-'        |
     ,'       ,-'           |
    (     _,-'\             |
     '._,'     :            | SSt
               '`
	return elf
}
