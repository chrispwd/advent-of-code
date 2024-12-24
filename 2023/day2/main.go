package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var validation = map[string]int{
	"R": 12,
	"G": 13,
	"B": 14,
}

func main() {
	gameLog, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println(asciiArt())
	fmt.Println("Sum of Valid Game IDs:", gameSum(gameLog, validation))
	fmt.Println("Sum of Powers of Minimum cube set:", powerSum(gameLog))
}

// Part 1
func gameSum(gameLog []string, validation map[string]int) int {
	gameAcc := 0
	valid := false

	for _, line := range gameLog {
		reId := regexp.MustCompile(`[0-9]+`) // game Ids
		var gId, _ = strconv.Atoi(reId.FindString(line))
		reRgb := regexp.MustCompile(`(\w+[;]*)+`) // Split the sets by semicolon
		var gRgb []string = reRgb.FindAllString(line, -1)[2:]
		var gameSets [][]string
		var lastGameIdx int

		for i, game := range gRgb {
			if string(game[len(game)-1]) == ";" {
				gameSets = append(gameSets, gRgb[lastGameIdx:i+1])
				lastGameIdx = i + 1
			}
			if i == len(gRgb)-1 {
				gameSets = append(gameSets, gRgb[lastGameIdx:])
			}
		}

		outer:
			for i, set := range gameSets {
				_ = i
				for k, pull := range set {
					colorRe := regexp.MustCompile(`(red|green|blue)`)
					color := colorRe.FindString(string(pull))
					if len(color) > 0 { // on a color
						amt, _ := strconv.Atoi(set[k-1]) // such that set[k-1] = 1 in [1 green]
						valid = isValidGame(amt, color, validation)
						if !valid {
							break outer
						}
					}
				}

			}
			
		if valid {
			gameAcc += gId
		}
	}
	
	return gameAcc
}

// Part 2
func powerSum(gameLog []string) int {
	powerAcc := 0

	for _, line := range gameLog {		
		reRgb := regexp.MustCompile(`(\w+[;]*)+`) // Split the sets by semicolon
		var fullGame []string = reRgb.FindAllString(line, -1)[2:]
		var gameSets [][]string
		var lastGameIdx int

		for i, gameSet := range fullGame { // split each set into subslice [[s1;] [s2;] [s3]]
			if string(gameSet[len(gameSet)-1]) == ";" {
				gameSets = append(gameSets, fullGame[lastGameIdx:i+1])
				lastGameIdx = i + 1
			}
			if i == len(fullGame)-1 {
				gameSets = append(gameSets, fullGame[lastGameIdx:])
			}
		}
		
		var maxR int
		var maxG int
		var maxB int
		for j, game := range gameSets { // [ [1 b; 2 g; 5 r] [3 b; 2 r] [...] ]
			_ = j
			for k, set := range game { // [1 b; 2 g; 5 r]
				_ = k
				
				if k == len(game)-1 {
					continue
				}
				
				colorRe := regexp.MustCompile(`(red|green|blue)`)
				color := colorRe.FindString(string(set))
				if len(color) < 1 { // if not a color, so next element (k+1) will be color
					currVal, _ := strconv.Atoi(set)
					switch colorRe.FindString(game[k+1]) {
					case "red":
						maxR = getMaxVal(maxR, currVal)
					case "green":
						maxG = getMaxVal(maxG, currVal)
					case "blue":
						maxB = getMaxVal(maxB, currVal)
					}
				}
			}
			
		}
		
		power := maxR * maxG * maxB
		powerAcc += power
	}
	
	return powerAcc
}

func getMaxVal(currMax int, contender int) int {
	if currMax < contender {
		return contender
	} else {
		return currMax
	}
}

func isValidGame(num int, color string,  validation map[string]int) bool {
	validity := false
	switch color {
	case "red":
		if num <= validation["R"] {
			validity = true
		}
	case "green":
		if num <= validation["G"] {
			validity = true
		}
	case "blue":
		if num <= validation["B"] {
			validity = true
		}
	}
	return validity
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
