package aoc2023day2

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MaxRed   = 12
	MaxBlue  = 14
	MaxGreen = 13
)

type Cubes struct {
	Red   int
	Blue  int
	Green int
}

func Run(input string) {
	inputArray := strings.Split(input, "\n")
	fmt.Println("Part 1 Solution:", SolvePart1(inputArray))
	fmt.Println("Part 2 Solution:", SolvePart2(inputArray))
}

func SolvePart1(input []string) string {
	parsedData := parseInput(input)
	result := 0
	for gameId, cubes := range parsedData {
		if cubes.Red <= MaxRed && cubes.Blue <= MaxBlue && cubes.Green <= MaxGreen {
			result += gameId
		}
	}

	return fmt.Sprintf("%d", result)
}

func SolvePart2(input []string) string {
	parsedData := parseInput(input)
	result := 0
	for _, cubes := range parsedData {
		result += cubes.Red * cubes.Blue * cubes.Green
	}

	return fmt.Sprintf("%d", result)
}

func parseInput(input []string) map[int]Cubes {
	parsedData := map[int]Cubes{}
	for _, line := range input {
		game := strings.Split(line, ":")
		gameId := 0
		_, err := fmt.Sscanf(game[0], "Game %d", &gameId)
		if err != nil {
			continue
		}

		cubes := Cubes{}
		gameData := strings.Split(strings.TrimSpace(game[1]), ";")
		for _, gameDataLine := range gameData {
			colorData := strings.Split(strings.TrimSpace(gameDataLine), ",")
			for _, color := range colorData {
				colorArr := strings.Split(strings.TrimSpace(color), " ")
				if len(colorArr) != 2 {
					continue
				}

				colorCount, err := strconv.Atoi(colorArr[0])
				if err != nil {
					continue
				}

				switch strings.TrimSpace(colorArr[1]) {
				case "red":
					if colorCount > cubes.Red {
						cubes.Red = colorCount
					}

				case "blue":
					if colorCount > cubes.Blue {
						cubes.Blue = colorCount
					}

				case "green":
					if colorCount > cubes.Green {
						cubes.Green = colorCount
					}
				}
			}
		}

		parsedData[gameId] = cubes
	}

	return parsedData
}
