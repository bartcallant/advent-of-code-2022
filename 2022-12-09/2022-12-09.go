package main

import (
	"advent-of-code-2022/utils/files/readFileAsLinesArray"
	"fmt"
	"strconv"
	"strings"
)

type Action struct {
	direction     string
	amountOfSteps int
}

type Point struct {
	id string
	x  int
	y  int
}

func buildPointId(x int, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func createPoint(x int, y int) Point {
	var id = buildPointId(x, y)

	return Point{id, x, y}
}

type PointDelta struct {
	x int
	y int
}

func buildPointDeltasForAction(action Action) []PointDelta {
	var pointDeltas = []PointDelta{}

	for i := 0; i < action.amountOfSteps; i++ {
		switch action.direction {
		case "L":
			pointDeltas = append(pointDeltas, PointDelta{x: -1, y: 0})
		case "R":
			pointDeltas = append(pointDeltas, PointDelta{x: 1, y: 0})
		case "U":
			pointDeltas = append(pointDeltas, PointDelta{x: 0, y: 1})
		case "D":
			pointDeltas = append(pointDeltas, PointDelta{x: 0, y: -1})
		default:
			fmt.Println("Found unknown direction: ", action.direction)
		}
	}

	return pointDeltas
}

func absolute(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func arePointsTouching(a Point, b Point) bool {
	var xEqual = a.x == b.x
	var yEqual = a.y == b.y

	if xEqual && yEqual {
		return true
	}

	var xDistance = absolute(a.x - b.x)
	var yDistance = absolute(a.y - b.y)

	if xDistance <= 1 && yDistance <= 1 {
		return true
	}

	return false
}

func displayTable(startingPosition Point, headPosition Point, tailPositions []Point) {
	fmt.Print("\n")

	for y := 20; y >= 0; y-- {
		for x := 0; x <= 25; x++ {
			if startingPosition.x == x && startingPosition.y == y {
				fmt.Print("s")

				continue
			}

			if headPosition.x == x && headPosition.y == y {
				fmt.Print("H")

				continue
			}

			var shouldBreak = false

			for pointIndex, point := range tailPositions {
				if point.x == x && point.y == y {
					fmt.Print(pointIndex + 1)

					shouldBreak = true
					break
				}
			}

			if shouldBreak {
				continue
			}

			fmt.Print(".")
		}
		fmt.Print("\n")
	}

	fmt.Print("\n")
}

func calculateNumberOfTailVisistedPoints(actions []Action, length int) int {
	var tailVisitedPoints = make(map[string]Point)

	var startingPoint = createPoint(11, 5)
	tailVisitedPoints[startingPoint.id] = startingPoint

	var points = []Point{
		createPoint(startingPoint.x, startingPoint.y),
	}

	for i := 0; i < length; i++ {
		points = append(points, createPoint(startingPoint.x, startingPoint.y))
	}

	for _, action := range actions {
		var pointDeltas = buildPointDeltasForAction(action)

		for _, pointDelta := range pointDeltas {
			points[0] = createPoint(points[0].x+pointDelta.x, points[0].y+pointDelta.y)

			for pointIndex := 1; pointIndex < len(points); pointIndex++ {
				var previousPoint = points[pointIndex-1]
				var currentPoint = points[pointIndex]

				if arePointsTouching(previousPoint, currentPoint) {
					break
				}

				var sameX = previousPoint.x == currentPoint.x
				var sameY = previousPoint.y == currentPoint.y

				if sameY {
					if previousPoint.x > points[pointIndex].x {
						points[pointIndex] = createPoint(points[pointIndex].x+1, points[pointIndex].y)
					} else {
						points[pointIndex] = createPoint(points[pointIndex].x-1, points[pointIndex].y)
					}
					continue
				}

				if sameX {
					if previousPoint.y > points[pointIndex].y {
						points[pointIndex] = createPoint(points[pointIndex].x, points[pointIndex].y+1)
					} else {
						points[pointIndex] = createPoint(points[pointIndex].x, points[pointIndex].y-1)
					}
					continue
				}

				if previousPoint.x > points[pointIndex].x {
					points[pointIndex] = createPoint(points[pointIndex].x+1, points[pointIndex].y)
				} else {
					points[pointIndex] = createPoint(points[pointIndex].x-1, points[pointIndex].y)
				}

				if previousPoint.y > points[pointIndex].y {
					points[pointIndex] = createPoint(points[pointIndex].x, points[pointIndex].y+1)
				} else {
					points[pointIndex] = createPoint(points[pointIndex].x, points[pointIndex].y-1)
				}
			}

			tailVisitedPoints[points[len(points)-1].id] = points[len(points)-1]
		}
		// displayTable(startingPosition, positions[0], positions[1:])
	}

	return len(tailVisitedPoints)
}

func main() {
	var lines = readFileAsLinesArray.Exec("input.txt")

	var actions = []Action{}

	for _, line := range lines {
		var parsedLine = strings.Fields(line)

		var direction = parsedLine[0]
		var amountOfSTeps, _ = strconv.ParseInt(parsedLine[1], 10, 0)

		actions = append(actions, Action{direction, int(amountOfSTeps)})
	}

	fmt.Println("PART1: ", calculateNumberOfTailVisistedPoints(actions, 1))
	fmt.Println("PART2: ", calculateNumberOfTailVisistedPoints(actions, 9))
}
