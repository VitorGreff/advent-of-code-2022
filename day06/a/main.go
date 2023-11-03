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
	for i := 0; i < len(line)-4; i++ {
		if line[i] != line[i+1] && line[i] != line[i+2] && line[i] != line[i+3] && line[i+1] != line[i+2] && line[i+1] != line[i+3] && line[i+2] != line[i+3] {
			return i + 4
		}
	}
	return 0
}
