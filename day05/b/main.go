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

func (s *stack) pushSlice(v []string) {
	*s = append(*s, v...)
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
	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}
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
			instantiateStacks(len(strings.Split(line, " ")) / 3)
		}
	}
}

func instantiateStacks(n int) {
	for i := 0; i < n; i++ {
		stacks = append(stacks, stack{})
	}
}
func evaluateInstructions(inst1, inst2, inst3 int) {
	var auxSlice []string
	for i := 0; i < inst1; i++ {
		if inst1 > 1 {
			auxSlice = append([]string{stacks[inst2-1].pop()}, auxSlice...)
		} else {
			stacks[inst3-1].push(stacks[inst2-1].pop())
		}
	}
	if inst1 > 1 {
		stacks[inst3-1].pushSlice(auxSlice)
	}
}
