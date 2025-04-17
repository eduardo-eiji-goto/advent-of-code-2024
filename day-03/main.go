package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

//* Day 03: Mull It Over
//* https://adventofcode.com/2024/day/3

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Sum Valid Mults:", getSumValidMultiplications(input))
}

func getSumValidMultiplications(input string) int {
	r, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	validMultIds := r.FindAllStringSubmatchIndex(input, -1)

	var sum int
	for _, v := range validMultIds {
		r, _ = regexp.Compile("[0-9]+")
		multValues := r.FindAllString(input[v[0]:v[1]], -1)

		a, _ := strconv.Atoi(multValues[0])
		b, _ := strconv.Atoi(multValues[1])
		sum += a * b
	}

	return sum
}

func readInput(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r += scanner.Text()
	}

	return r, nil
}
