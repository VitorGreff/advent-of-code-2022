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

	fmt.Printf("part 1: %d\n", totalPoints)
}

func evaluate(opponent string, player string) uint {
	var ret uint
	switch opponent {
	case "A":
		if player == "X" {
			ret = 4
			break
		} else if player == "Y" {
			ret = 8
			break
		} else {
			ret = 3
			break
		}
	case "B":
		if player == "X" {
			ret = 1
			break
		} else if player == "Y" {
			ret = 5
			break
		} else {
			ret = 9
			break
		}
	case "C":
		if player == "X" {
			ret = 7
			break
		} else if player == "Y" {
			ret = 2
			break
		} else {
			ret = 6
			break
		}
	}
	return ret
}
