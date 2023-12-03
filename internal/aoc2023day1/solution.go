package aoc2023day1

import (
	"fmt"
	"strconv"
	"strings"
)

func Run(input string) {
	inputArray := strings.Split(input, "\n")
	fmt.Println("Part 1 Solution:", SolvePart1(inputArray))
	fmt.Println("Part 2 Solution:", SolvePart2(inputArray))
}

func SolvePart1(input []string) string {
	sum := 0
	for _, line := range input {
		firstChar := ""
		lastChar := ""
		for _, char := range line {
			if char >= '0' && char <= '9' {
				if firstChar == "" {
					firstChar = string(char)
				} else {
					lastChar = string(char)
				}
			}
		}

		currentInt, err := strconv.Atoi(firstChar + lastChar)
		if err == nil {
			sum += currentInt
		}
	}

	return fmt.Sprintf("%d", sum)
}

func SolvePart2(input []string) string {
	sum := 0
	for _, line := range input {
		firstChar := ""
		firstCharPosition := -1
		lastChar := ""
		lastCharPosition := -1
		numberStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		numberStringToNumberMap := map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		}
		for index, char := range line {
			if char >= '0' && char <= '9' {
				if firstCharPosition == -1 {
					firstCharPosition = index
					firstChar = string(char)
				} else {
					lastCharPosition = index
					lastChar = string(char)
				}
			}
		}

		for _, numberString := range numberStrings {
			if strings.Contains(line, numberString) {
				position := strings.Index(line, numberString)
				if firstCharPosition == -1 || position < firstCharPosition {
					firstChar = numberStringToNumberMap[numberString]
					firstCharPosition = position
				}

				if position > lastCharPosition && position != firstCharPosition {
					lastChar = numberStringToNumberMap[numberString]
					lastCharPosition = position
				}
			}
		}

		for index, char := range line {
			if char >= '0' && char <= '9' {
				if index > lastCharPosition && index != firstCharPosition {
					lastChar = string(char)
					lastCharPosition = index
					break
				}
			}
		}

		if firstCharPosition == -1 && lastCharPosition == -1 {
			continue
		}

		current := firstChar
		if lastCharPosition != -1 {
			current += lastChar
		}

		currentInt, err := strconv.Atoi(current)
		if err == nil {
			sum += currentInt
		}
	}

	return fmt.Sprintf("%d", sum)
}
