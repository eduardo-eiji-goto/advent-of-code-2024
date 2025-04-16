package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//* Day 02: Red-Nosed Reports
//* https://adventofcode.com/2024/day/2

func main() {
	reports, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Safe Reports:", getSafeReportCount(reports))
	fmt.Println("Safe Reports (Problem Dampener):", getSafeReportCountDampener(reports))
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

func getSafeReportCountDampener(reports [][]int) int {
	var safeReportCount int

	for i := range reports {
		reportCurrent := reports[i]
		isReportSafe, violations := validateReport(reportCurrent)

		if isReportSafe != true {
			violationPossibilities := []int{0, len(reportCurrent) - 1}
			violationPossibilities = append(violationPossibilities, violations...)

			for _, violationId := range violationPossibilities {
				var reportDampener []int
				reportDampener = append(reportDampener, reportCurrent[:violationId]...)
				reportDampener = append(reportDampener, reportCurrent[violationId+1:]...)

				isReportSafeDampener, _ := validateReport(reportDampener)

				if isReportSafeDampener == true {
					isReportSafe = true
				}
			}
		}

		if isReportSafe == true {
			safeReportCount += 1
		}
	}

	return safeReportCount
}

func isItemValid(report []int, index int) bool {
	differencePrevious := report[index] - report[index-1]
	differencePreviousAbs := math.Abs(float64(differencePrevious))
	if differencePreviousAbs < 1 || differencePreviousAbs > 3 {
		return false
	}

	differenceNext := report[index+1] - report[index]
	differenceNextAbs := math.Abs(float64(differenceNext))
	if differenceNextAbs < 1 || differenceNextAbs > 3 {
		return false
	}

	if math.Signbit(float64(differencePrevious)) != math.Signbit(float64(differenceNext)) {
		return false
	}

	return true
}

func validateReport(report []int) (bool, []int) {
	isValid := true
	var violations []int

	for i := range report {
		if i > 0 && i < len(report)-1 {
			if isItemValid(report, i) == false {
				violations = append(violations, i)
			}
		}
	}

	if len(violations) > 0 {
		isValid = false
	}

	return isValid, violations
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
