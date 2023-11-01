package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var totalPoints uint
	var totalPoints2 uint

	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalPoints += evaluate(string(scanner.Text()[0]),
			string(scanner.Text()[2]))
		totalPoints2 += evaluate2(string(scanner.Text()[0]),
			string(scanner.Text()[2]))
	}

	fmt.Printf("part 1: %d\n", totalPoints)
	fmt.Printf("part 2: %d", totalPoints2)
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

func evaluate2(opponent string, outcome string) uint {
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
