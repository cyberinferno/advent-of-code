package main

import (
	"flag"
	"fmt"
	"github.com/cyberinferno/advent-of-code/internal/aoc2023day1"
	"os"
)

var funcMap map[string]map[string]func(input string)

func main() {
	var year string
	var day string
	var input string
	flag.StringVar(&year, "year", "2023", "Year to run")
	flag.StringVar(&day, "day", "1", "Day to run")
	flag.StringVar(&input, "input", "", "Input file to run")
	flag.Parse()
	if _, ok := funcMap[year]; !ok {
		fmt.Println("Year not found")
		return
	}

	if _, ok := funcMap[year][day]; !ok {
		fmt.Println("Day not found")
		return
	}

	if input == "" {
		input = fmt.Sprintf("internal/aoc%sday%s/input.txt", year, day)
	}

	fmt.Println("Input file:", input)
	fileContents, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading input file")
		return
	}

	fmt.Printf("Running solutions for year %s day %s\n", year, day)
	funcMap[year][day](string(fileContents))
}

func init() {
	funcMap = map[string]map[string]func(input string){
		"2023": {
			"1": aoc2023day1.Run,
		},
	}
}
