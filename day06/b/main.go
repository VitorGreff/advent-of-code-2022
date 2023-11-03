package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var answer int
	file, err := os.Open("../6.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		answer = checkSequence(scanner.Text())
	}
	fmt.Println(answer)
}

func checkSequence(line string) int {
	for i := 0; i < len(line); i++ {
		var aux = []byte{line[i]}
		for j := i + 1; j < len(line); j++ {
			if iterateThroughLine(aux, line[j]) {
				aux = append(aux, line[j])
			} else {
				break
			}
			if len(aux) == 14 {
				return j + 1
			}
		}
	}
	return 0
}

func iterateThroughLine(line []byte, char byte) bool {
	for _, b := range line {
		if char != b {
			continue
		} else {
			return false
		}
	}
	return true
}
