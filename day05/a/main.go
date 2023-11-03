package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type stack []string

var stack1 = stack{}
var stack2 = stack{}
var stack3 = stack{}

var stacks = map[int]*stack{
	1: &stack1,
	2: &stack2,
	3: &stack3,
}

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
	file, err := os.Open("../5.txt")
	if err != nil {
		log.Fatal("Error opening the file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && scanner.Text() != "" {
		for string(scanner.Text()[1]) != "1" {
			if scanner.Text()[1] != ' ' {
				readFileToStack(&stack1, string(scanner.Text()[1]))
			}
			if scanner.Text()[5] != ' ' {
				readFileToStack(&stack2, string(scanner.Text()[5]))
			}
			if scanner.Text()[9] != ' ' {
				readFileToStack(&stack3, string(scanner.Text()[9]))
			}
			scanner.Scan()
		}
	}
	for scanner.Scan() {
		index1, err1 := strconv.Atoi(string(scanner.Text()[5]))
		index2, err2 := strconv.Atoi(string(scanner.Text()[12]))
		index3, err3 := strconv.Atoi(string(scanner.Text()[17]))
		if err1 != nil || err2 != nil || err3 != nil {
			log.Fatal("Error parsing instructions")
		}
		evaluateInstructions(index1, index2, index3)
	}
	fmt.Printf("End of first stack: %s\n", stack1[len(stack1)-1])
	fmt.Printf("End of second stack: %s\n", stack2[len(stack2)-1])
	fmt.Printf("End of third stack: %s", stack3[len(stack3)-1])

}

func evaluateInstructions(inst1, inst2, inst3 int) {
	for i := 0; i < inst1; i++ {
		stacks[inst3].push(stacks[inst2].pop())
	}
}

func readFileToStack(stack *stack, str string) {
	*stack = append([]string{str}, *stack...)
}
