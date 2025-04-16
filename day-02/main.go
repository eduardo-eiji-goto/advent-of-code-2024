package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Safe Reports:", getSafeReportCount(reports))
}

func getSafeReportCount(reports [][]int) int {
	var safeReportCount int

	for i := range reports {
		reportCurrent := reports[i]

		differenceViolation := false
		directionInversion := false
		var directionPrevious bool

		for j := range reportCurrent {
			if j > 0 {
				differenceCurrent := reportCurrent[j] - reportCurrent[j-1]
				directionCurrent := math.Signbit(float64(differenceCurrent))

				differenceCurrent = int(math.Abs(float64(differenceCurrent)))

				if differenceCurrent < 1 || differenceCurrent > 3 {
					differenceViolation = true
				}

				if j > 1 && directionCurrent != directionPrevious {
					directionInversion = true
				}

				directionPrevious = directionCurrent
			}
		}

		if differenceViolation == false && directionInversion == false {
			safeReportCount += 1
		}
	}

	return safeReportCount
}

func readInput(name string) ([][]int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var r [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var values []int

		valuesString := strings.Split(scanner.Text(), " ")
		for i := range valuesString {
			v, err := strconv.Atoi(valuesString[i])
			if err != nil {
				return nil, err
			}
			values = append(values, v)
		}

		r = append(r, values)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return r, nil
}
