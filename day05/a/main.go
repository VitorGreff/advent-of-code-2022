package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack []string

var stacks []stack
var aux int = 0

func (s *stack) push(v string) {
	*s = append(*s, v)
}

func (s *stack) pop() string {
	l := len(*s)
	if l == 0 {
		return ""
	}
	value := (*s)[l-1]
	*s = (*s)[:l-1]
	return value
}

func main() {
	var path string = "../5.txt"
	checkStackNumber(path)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening the file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && scanner.Text() != "" {
		if string(scanner.Text()[1]) != "1" {
			for i := 1; i < len(scanner.Text())+1; i += 4 {
				readFileToStack(&stacks[aux], string(scanner.Text()[i]))
				aux++
			}
			aux = 0
		}
	}
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		index1, err1 := strconv.Atoi(string(s[1]))
		index2, err2 := strconv.Atoi(string(s[3]))
		index3, err3 := strconv.Atoi(string(s[5]))
		if err1 != nil || err2 != nil || err3 != nil {
			log.Fatal("Error parsing instructions")
		}
		evaluateInstructions(index1, index2, index3)
	}
	fmt.Println(stacks)
}

func readFileToStack(stack *stack, str string) {
	if str != " " {
		*stack = append([]string{str}, *stack...)
	}
}

func checkStackNumber(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening the file")
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()
	for scanner.Scan() && scanner.Text() != "" {
		if line := string(scanner.Text()); string(line[1]) == "1" {
			n, err := strconv.Atoi(string(line[len(line)-2]))
			instantiateStacks(n)
			if err != nil {
				log.Fatal("Error declaring the number of stacks")
			}
		}
	}
}

func instantiateStacks(n int) {
	for i := 0; i < n; i++ {
		stacks = append(stacks, stack{})
	}
}

func evaluateInstructions(inst1, inst2, inst3 int) {
	for i := 0; i < inst1; i++ {
		stacks[inst3-1].push(stacks[inst2-1].pop())
	}
}
