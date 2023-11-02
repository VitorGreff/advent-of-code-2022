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

	file, err := os.Open("../3.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += int(calculatePriority(findEqual(separateCompartments(scanner.Text()))))
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func separateCompartments(fullCompartment string) (c1 string, c2 string) {
	var size int = len(fullCompartment) / 2
	return fullCompartment[:size], fullCompartment[size:]
}

func findEqual(c1 string, c2 string) rune {
	for _, char := range c1 {
		for _, char2 := range c2 {
			if char == char2 {
				return char
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
