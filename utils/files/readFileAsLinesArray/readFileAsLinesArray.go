package readFileAsLinesArray

import (
	"os"
	"log"
	"bufio"
)

func Exec(filename string) []string {
    result := []string{}

    file, err := os.Open(filename)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
 
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
 
	file.Close()

    return result
}
