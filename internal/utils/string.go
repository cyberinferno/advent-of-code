package utils

import "strings"

func FindAllStringPositions(mainString, subString string) []int {
	var positions []int
	startIndex := 0

	for {
		index := strings.Index(mainString[startIndex:], subString)
		if index == -1 {
			break
		}

		positions = append(positions, startIndex+index)
		startIndex += index + len(subString)
	}

	return positions
}
