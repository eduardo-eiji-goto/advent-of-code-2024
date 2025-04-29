package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
}

func readInput(name string) (map[int]int, [][]int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := scanner.Next()


	}

	if err := scanner.Error(); err != nil {
		return nil, nil, err
	}
	return [], err
}
