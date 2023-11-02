package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	var sum int
	var counter uint = 0
	var slice []string

	file, err := os.Open("../3.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
		counter++
		if counter == 3 {
			sum += calculatePriority(findEqual(slice[0], slice[1], slice[2]))
			slice = []string{}
			counter = 0
		}
	}
	fmt.Printf("Part 2: %d\n", sum)
}

func findEqual(c1 string, c2 string, c3 string) rune {
	for _, char := range c1 {
		for _, char2 := range c2 {
			for _, char3 := range c3 {
				if char == char2 && char == char3 {
					return char
				}
			}
		}
	}
	log.Fatal("Error finding equal")
	return 0
}

func calculatePriority(char rune) int {
	if unicode.IsUpper(char) {
		return int(unicode.ToLower(char)) - 96 + 26
	}
	return int(unicode.ToLower(char)) - 96
}
