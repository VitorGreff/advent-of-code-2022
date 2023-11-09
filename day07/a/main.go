package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	total  int
	leaves *[]dir
}

func (t *dir) calculateSize() int {
	for i := 0; i < len(*t.leaves); i++ {
		t.total += (*t.leaves)[i].total
	}
	return t.total
}

var position string = "/"

func main() {
	directories := make(map[string]*dir)
	file, err := os.Open("../7.txt")
	if err != nil {
		log.Fatal("Error reading the file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	External:
		if strings.Split(scanner.Text(), " ")[0] == "$" {
			if strings.Split(scanner.Text(), " ")[1] == "cd" {
				dirName := strings.Split(scanner.Text(), " ")[2]
				if dirName != ".." {
					directories[dirName] = &dir{}
					position = dirName
				}
			} else if strings.Split(scanner.Text(), " ")[1] == "ls" {
				for scanner.Scan() {
					if strings.Split(scanner.Text(), " ")[0] == "$" {
						goto External
					}
					if strings.Split(scanner.Text(), " ")[0] != "dir" {
						size, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
						directories[position].total += size
						fmt.Printf("Diretorio %s recebe %d\n", position, size)
					}
				}
			}
		}
	}
	fmt.Println(directories["a"].total)
	fmt.Println(directories["e"].total)
	fmt.Println(directories["d"].total)
}
