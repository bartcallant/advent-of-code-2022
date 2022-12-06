package main

import (
	"advent-of-code-2022/utils/files/readFileAsLinesArray"
	"fmt"
)

func findIndexOfUniqueDistinctCharacters(line string, count int) int {
	var result = -1

	var lineAsRune = []rune(line)

	var goBack = count - 1

	for i := goBack; i < len(lineAsRune); i++ {
		var chars = lineAsRune[i-goBack : i+1]

		var charsMap = map[rune]rune{}

		for _, char := range chars {
			var _, exists = charsMap[char]

			if !exists {
				charsMap[char] = char
			}
		}

		if len(charsMap) == count {
			result = i + 1

			break
		}
	}

	return result
}

func main() {
	var lines = readFileAsLinesArray.Exec("input.txt")

	var part1Results = []int{}
	var part2Results = []int{}

	for _, line := range lines {
		part1Results = append(part1Results, findIndexOfUniqueDistinctCharacters(line, 4))
		part2Results = append(part2Results, findIndexOfUniqueDistinctCharacters(line, 14))
	}

	fmt.Println("PART1: ", part1Results)
	fmt.Println("PART2: ", part2Results)
}
