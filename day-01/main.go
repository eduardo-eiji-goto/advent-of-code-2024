package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

//* Day 01: Historian Hysteria
//* https://adventofcode.com/2024/day/1

func main() {
	a, b, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Total Distance:", getTotalDistance(a, b))

}

func readInput(name string) ([]int, []int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var a []int
	var b []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineValues := strings.Split(scanner.Text(), "   ")

		aValue, err := strconv.Atoi(lineValues[0])
		if err != nil {
			return nil, nil, err
		}

		bValue, err := strconv.Atoi(lineValues[1])
		if err != nil {
			return nil, nil, err
		}

		a = append(a, aValue)
		b = append(b, bValue)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return a, b, nil
}

func getTotalDistance(a []int, b []int) int {
	aSorted := make([]int, len(a))
	bSorted := make([]int, len(b))

	// Copying the slices in case we need the original ones intact in the future
	copy(aSorted, a)
	copy(bSorted, b)

	slices.Sort(aSorted)
	slices.Sort(bSorted)

	var totalDistance int
	for i, v := range aSorted {
		// NOTE: Weird ass conversions to get the difference as int
		totalDistance = totalDistance + int(math.Abs(float64(v-bSorted[i])))
	}

	return totalDistance
}
