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
	games, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println(asciiArt())
	fmt.Println("Sum of Valid Game IDs:", gameSum(games, validation))
}

// Part 1
func gameSum(games []string, validation map[string]int) int {
	// TODO :: you are here
	gameAcc := 0
	for _, line := range games {

		// capture gameID (quick regex?)
		re := regexp.MustCompile(`[0-9]+`)
		var gId, _ = strconv.Atoi(re.FindString(line))
		// Split the sets by semicolon
		// Split the colors
		// Compare if colcount > constants, skip line

		// DEBUG
		gameAcc += gId
		fmt.Println(line)
	}
	return gameAcc
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
