package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "regexp"
)

func main() {
	games, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println(asciiArt())
	fmt.Println("Sum of Valid Game IDs:", gameSum(games))
}

func gameSum(games []string) int {
	// TODO :: you are here
	return 0
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