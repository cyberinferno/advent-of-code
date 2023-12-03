package aoc2023day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	inputArray := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	result := SolvePart1(inputArray)
	assert.Equal(t, "142", result)
}

func TestSolvePart2(t *testing.T) {
	inputArray := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	result := SolvePart2(inputArray)
	assert.Equal(t, "281", result)
}
