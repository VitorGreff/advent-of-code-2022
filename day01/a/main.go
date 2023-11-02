package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var aux, greater, counter int = 0, 0, 0

	file, err := os.Open("../1.txt")
	if err != nil {
		log.Fatal("Error reading the txt file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			value, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal("Error converting text")
			}
			aux += value
		} else {
			if aux > greater {
				greater = aux
			}
			aux = 0
			counter++
		}
	}
	fmt.Printf("Elf carrying more calories: %d\nTotal calories: %d\n", counter, greater)
}
