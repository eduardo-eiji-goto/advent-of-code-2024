package main

import (
	"bufio"
	"fmt"
	"os"
)

//* Day 04: Ceres Search
//* https://adventofcode.com/2024/day/4

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	//for i := range input {
	//	for j := range input[i] {
	//		fmt.Print(string(input[i][j]), " ")
	//	}
	//	fmt.Println("")
	//}
	//fmt.Println("---")

	var result int

	for i := range input {
		for j := range input[i] {
			if string(input[i][j]) == "X" {
				result += checkAddress(input, Address{i, j})
			}
		}
	}

	fmt.Println("XMAS Count:", result)
}

type Address struct {
	row int
	col int
}

type Direction struct {
	rowGrowth int
	colGrowth int
}

func checkAddress(matrix []string, address Address) int {
	var r int

	// Get Verticals
	if address.row > 2 {
		if checkDirection(matrix, address, Direction{-1, 0}) == "XMAS" {
			r += 1
		}
	}

	if address.row < len(matrix)-3 {
		if checkDirection(matrix, address, Direction{+1, 0}) == "XMAS" {
			r += 1
		}
	}

	// Get Horizontals
	if address.col > 2 {
		if checkDirection(matrix, address, Direction{0, -1}) == "XMAS" {
			r += 1
		}
	}

	if address.col < len(matrix[0])-3 {
		if checkDirection(matrix, address, Direction{0, +1}) == "XMAS" {
			r += 1
		}
	}

	// Get Diagonals
	if address.row > 2 && address.col > 2 {
		if checkDirection(matrix, address, Direction{-1, -1}) == "XMAS" {
			r += 1
		}
	}

	if address.row > 2 && address.col < len(matrix[0])-3 {
		if checkDirection(matrix, address, Direction{-1, +1}) == "XMAS" {
			r += 1
		}
	}

	if address.row < len(matrix)-3 && address.col > 2 {
		if checkDirection(matrix, address, Direction{+1, -1}) == "XMAS" {
			r += 1
		}
	}

	if address.row < len(matrix)-3 && address.col < len(matrix[0])-3 {
		if checkDirection(matrix, address, Direction{+1, +1}) == "XMAS" {
			r += 1
		}
	}

	return r
}

func checkDirection(matrix []string, address Address, direction Direction) string {
	var r string

	for i := range 4 {
		valueRow := address.row + (i * direction.rowGrowth)
		valueCol := address.col + (i * direction.colGrowth)

		r += string(matrix[valueRow][valueCol])
	}

	return r
}

func readInput(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var r []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return r, nil
}
