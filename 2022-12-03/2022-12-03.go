package main

import (
	"fmt"

	"advent-of-code-2022/utils/arrays/reduceArray"
	"advent-of-code-2022/utils/files/readFileAsLinesArray"
)

func chunkRuneArrayByChunkSize[T any](input []T, chunkSize int) [][]T {
	var chunks [][]T

	for i := 0; i < len(input); i += chunkSize {
		var end = i + chunkSize

		if end > len(input) {
			end = len(input)
		}

		chunks = append(chunks, input[i:end])
	}

	return chunks
}

func chunkRuneArrayByNumberOfChunks[T any](input []T, numberOfChunks int) [][]T {
	var chunkSize = len(input) / numberOfChunks

	return chunkRuneArrayByChunkSize(input, chunkSize)
}

func itemExistsInArray(item rune, array []rune) bool {
	var result bool = false

	for _, arrayItem := range array {
		if arrayItem == item {
			result = true

			break
		}
	}

	return result
}

func generatePriorityForRune(character rune) int {
	var possibleLowercase = character - 96

	if possibleLowercase < 0 {
		var uppercase = possibleLowercase + 32 + 26

		return int(uppercase)
	}

	return int(possibleLowercase)
}

func part1(lines []string) {
	var result = 0

	for _, line := range lines {
		var runeArray = []rune(line)

		var lineInChunks = chunkRuneArrayByNumberOfChunks(runeArray, 2)

		// TODO: Make dynamoic?
		var first = lineInChunks[0]
		var second = lineInChunks[1]

		var duplicates = []rune{}

		for _, item := range first {
			if itemExistsInArray(item, second) {
				if !itemExistsInArray(item, duplicates) {
					duplicates = append(duplicates, item)
				}
			}
		}

		var duplicateSum = reduceArray.Exec(duplicates, func(acc int, current rune) int {
			return acc + generatePriorityForRune(current)
		}, int(0))

		result += duplicateSum
	}

	fmt.Println("PART1:", result)
}

func part2(lines []string) {
	var groupsSum int = 0

	var groups = chunkRuneArrayByChunkSize(lines, 3)

	for _, group := range groups {
		var possibleGroupIds = []rune{}

		var first = group[0]

		for _, character := range first {
			if itemExistsInArray(character, possibleGroupIds) {
				continue
			}

			var characterFoundInNumberOfBags = 1

			for i := 1; i < len(group); i++ {
				var bag = []rune(group[i])

				if itemExistsInArray(character, bag) {
					characterFoundInNumberOfBags += 1
				}
			}

			if characterFoundInNumberOfBags == len(group) {
				possibleGroupIds = append(possibleGroupIds, character)
			}
		}

		var possibleGroupIdsSum = reduceArray.Exec(possibleGroupIds, func(acc int, current rune) int {
			return acc + generatePriorityForRune(current)
		}, int(0))

		groupsSum += possibleGroupIdsSum
	}

	fmt.Println("PART2:", groupsSum)
}

func main() {
	var lines = readFileAsLinesArray.Exec("input.txt")

	part1(lines)
	part2(lines)
}
