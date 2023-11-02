package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var aux, greater, counter int = 0, 0, 0
	var slice = []int{}

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
			slice = append(slice, aux)
			if aux > greater {
				greater = aux
			}
			aux = 0
			counter++
		}
	}
	fmt.Printf("Sum of the top 3 elfs: %d", findTop3Sum(slice))
}

func findTop3Sum(slice []int) int {
	var ret int
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	for _, value := range slice[:3] {
		ret += value
	}
	return ret
}
