package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var totalPoints uint

	file, err := os.Open("../2.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalPoints += evaluate(string(scanner.Text()[0]),
			string(scanner.Text()[2]))
	}

	fmt.Printf("part 2: %d\n", totalPoints)
}

func evaluate(opponent string, outcome string) uint {
	var ret uint
	switch opponent {
	case "A":
		if outcome == "X" {
			ret = 3
			break
		} else if outcome == "Y" {
			ret = 4
			break
		} else {
			ret = 8
			break
		}
	case "B":
		if outcome == "X" {
			ret = 1
			break
		} else if outcome == "Y" {
			ret = 5
			break
		} else {
			ret = 9
			break
		}
	case "C":
		if outcome == "X" {
			ret = 2
			break
		} else if outcome == "Y" {
			ret = 6
			break
		} else {
			ret = 7
			break
		}
	}
	return ret
}
