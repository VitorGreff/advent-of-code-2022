package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var counter uint = 0

	file, err := os.Open("../4.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counter += checkPair(scanner.Text())
	}
	fmt.Println(counter)
}

func checkPair(line string) uint {
	v := strings.Split(line, ",")
	c1 := strings.Split(v[0], "-")
	c2 := strings.Split(v[1], "-")

	return evaluate(c1, c2)
}

func evaluate(c1 []string, c2 []string) uint {
	n1, err1 := strconv.Atoi(c1[0])
	n2, err2 := strconv.Atoi(c1[1])
	n3, err3 := strconv.Atoi(c2[0])
	n4, err4 := strconv.Atoi(c2[1])
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		log.Fatal("Error converting values")
	}

	if n1 <= n3 && n2 >= n4 {
		return 1
	} else if n3 <= n1 && n4 >= n2 {
		return 1
	}
	return 0
}
