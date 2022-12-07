package main

import (
	"advent-of-code-2022/utils/files/readFileAsLinesArray"
	"fmt"
	"strings"
)

type ExecutedCommand struct {
	command string
	output  []string
}

func newExecutedCommand(command string, output []string) *ExecutedCommand {
	var executedCommand = ExecutedCommand{command, output}

	return &executedCommand
}

func parseExecutedCommendsFromLines(lines []string) []ExecutedCommand {
	var executedCommands = []ExecutedCommand{}

	var lastCommand = ""
	var lastCOmmandOutput = []string{}

	for _, line := range lines {
		if strings.HasPrefix(line, "$ ") {
			if lastCommand != "" {
				executedCommands = append(executedCommands, *newExecutedCommand(lastCommand, lastCOmmandOutput))

				lastCommand = ""
				lastCOmmandOutput = []string{}
			}

			lastCommand = line
		} else {
			lastCOmmandOutput = append(lastCOmmandOutput, line)
		}
	}

	executedCommands = append(executedCommands, *newExecutedCommand(lastCommand, lastCOmmandOutput))

	return executedCommands
}

func calculateFolderSizes(executedCommands []ExecutedCommand) map[string]int {
	var currentFolder = "/"

	var folderSizes = make(map[string]int)
	var folderParentFolder = make(map[string]string)

	for _, executedCommand := range executedCommands {

		if strings.HasPrefix(executedCommand.command, "$ cd ") {
			var destination = strings.Replace(executedCommand.command, "$ cd ", "", 1)

			if destination == "/" {
				currentFolder = "/"
			} else if destination == ".." {
				currentFolder = folderParentFolder[currentFolder]
			} else {
				currentFolder = strings.Join([]string{currentFolder, destination}, "/")
			}
		}

		if strings.HasPrefix(executedCommand.command, "$ ls") {
			for _, fileString := range executedCommand.output {
				var fileStringParts = strings.Fields(fileString)

				var dirOrSize = fileStringParts[0]
				var dorOrFileName = fileStringParts[1]

				if dirOrSize == "dir" {
					var folderPath = strings.Join([]string{currentFolder, dorOrFileName}, "/")
					folderParentFolder[folderPath] = currentFolder
				} else {
					var size = 0
					fmt.Sscan(dirOrSize, &size)

					var temporaryParentFolder = currentFolder

					for parentFolderExists := true; parentFolderExists; temporaryParentFolder, parentFolderExists = folderParentFolder[temporaryParentFolder] {
						folderSizes[temporaryParentFolder] += size
					}
				}
			}
		}
	}

	return folderSizes
}

func main() {
	var lines = readFileAsLinesArray.Exec("input.txt")

	var executedCommands = parseExecutedCommendsFromLines(lines)
	var folderSizes = calculateFolderSizes(executedCommands)

	var maxFolderSize = 100000
	var sumOfFolderSizeLessthanMexFolderSize = 0

	var filesystemSize = 70000000
	var neededSizeForUpdate = 30000000
	var currentOccupiedSize = folderSizes["/"]

	var smallestEgibleFolderSize = filesystemSize

	for _, folderSize := range folderSizes {
		if folderSize <= maxFolderSize {
			sumOfFolderSizeLessthanMexFolderSize += folderSize
		}

		var freeSpaceAfterDelete = filesystemSize - currentOccupiedSize + folderSize

		if freeSpaceAfterDelete > neededSizeForUpdate {
			if folderSize < smallestEgibleFolderSize {
				smallestEgibleFolderSize = folderSize
			}
		}
	}

	fmt.Println("PART1: ", sumOfFolderSizeLessthanMexFolderSize)
	fmt.Println("PART2: ", smallestEgibleFolderSize)
}
