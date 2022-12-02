package main

import (
    "fmt"
	"log"
	"strings"

	"advent-of-code-2022/utils/files/readFileAsLinesArray"
)

func calculateHandScore(hand string) int {
	var scores = map[string]int{
		"rock":  1,
		"paper": 2,
		"scissors": 3,
	}

	score, scoreFound := scores[hand]

	if scoreFound != true {
		log.Fatalf("Unable to get score for hand %s", hand)
	}

	return score
}

func getResultScore(result string) int {
	scores := map[string]int{
		"lost":  0,
		"draw": 3,
		"win": 6,
	}

	score, ok := scores[result];

	if ok != true {
		log.Fatalf("Unable to generate score for %s", result)
	}

	return score
}

func determineResult(hand1 string, hand2 string) string {
	var resultMap = map[string]string{
		"rock+rock": "draw",
		"rock+paper": "win",
		"rock+scissors": "lost",
		"paper+rock": "lost",
		"paper+paper": "draw",
		"paper+scissors": "win",
		"scissors+rock": "win",
		"scissors+paper": "lost",
		"scissors+scissors": "draw",
	}

	var result, resultFound = resultMap[fmt.Sprintf("%s+%s", hand1, hand2)]

	if (resultFound != true) {
		log.Fatalf("No result found for round %s vs %s", hand1, hand2)
	}

	return result
}
 
func calculateRoundScorePart1(opponentHand string, ownHand string) int {
	opponentHandConverter := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	convertedOpponentHand, conversionOoponentHandWorked := opponentHandConverter[opponentHand]

	if conversionOoponentHandWorked != true {
		log.Fatalf("Unable to convert opponent hand %s", opponentHand)
	}

	ownHandConverter := map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	convertedOwnHand, conversionOwnHandWorked := ownHandConverter[ownHand]

	if conversionOwnHandWorked != true {
		log.Fatalf("Unable to convert own hand %s", ownHand)
	}

	result := determineResult(convertedOpponentHand, convertedOwnHand)

	return calculateHandScore(convertedOwnHand) + getResultScore(result)
}

func calculateRoundScorePart2(opponentHand string, desiredResult string) int {
	opponentHandConverter := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	convertedOpponentHand, conversionOoponentHandWorked := opponentHandConverter[opponentHand]

	if conversionOoponentHandWorked != true {
		log.Fatalf("Unable to convert opponent hand %s", opponentHand)
	}

	desiredResultConverter := map[string]string{
		"X": "lost",
		"Y": "draw",
		"Z": "win",
	}

	convertedDesiredResult, conversionDesiredResultWorked := desiredResultConverter[desiredResult]

	if conversionDesiredResultWorked != true {
		log.Fatalf("Unable to convert desired result %s", desiredResult)
	}

	desiredResultToHandConverter := map[string]string{
		determineResult(convertedOpponentHand, "rock"): "rock",
		determineResult(convertedOpponentHand, "paper"): "paper",
		determineResult(convertedOpponentHand, "scissors"): "scissors",
	}

	convertedDesiredResultHand, conversionDesiredResultHandWorked := desiredResultToHandConverter[convertedDesiredResult]

	if conversionDesiredResultHandWorked != true {
		log.Fatalf("Unable to determine hand for desired result %s", convertedDesiredResult)
	}

	return calculateHandScore(convertedDesiredResultHand) + getResultScore(convertedDesiredResult)
}

func main() {
	var lines = readFileAsLinesArray.Exec("input.txt")

	var scorePart1 = 0
	var scorePart2 = 0

	for _, line := range lines {
		var round = strings.Fields(line)

		var opponentHand = round[0]
		var ownHand = round[1]

		scorePart1 += calculateRoundScorePart1(opponentHand, ownHand)
		scorePart2 += calculateRoundScorePart2(opponentHand, ownHand)	
	}

	fmt.Println("PART1: Score", scorePart1)
	fmt.Println("PART2: Score", scorePart2)
}
