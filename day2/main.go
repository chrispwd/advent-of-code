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

// const R int = 12
// const G int = 13
// const B int = 14

func main() {
	gameLog, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println(asciiArt())
	fmt.Println("Sum of Valid Game IDs:", gameSum(gameLog, validation))
}

// Part 1
func gameSum(gameLog []string, validation map[string]int) int {
	// TODO :: you are here
	gameAcc := 0
	valid := false

	for _, line := range gameLog {
		// DEBUG
		fmt.Println(line)
		
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
		}

		outer:
		for _, set := range gameSets {
			for k, pull := range set {
				colorRe := regexp.MustCompile(`(red|green|blue)`)
				color := colorRe.FindString(string(pull))
				if len(color) > 0 { // on a color
					colorAmt, _ := strconv.Atoi(set[k-1]) // such that set[k-1] = 1 in [1 green]
					valid = isValidGame(colorAmt, color, validation)
					if !valid {
						break outer
					}
				}
			}

		}
		fmt.Printf("%d red, %d green, %d blue - %v\n", validation["R"], validation["G"], validation["B"], valid)
		// TODO :: Figure out why this is not adding up correctly
		if valid {
			fmt.Println("Adding ID", gId)
			gameAcc += gId
		}
	}
	
	return gameAcc
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
