package main

import (
	"bufio"
	"day01/calculator"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readLines("../assets/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	leftArray, rightArray := sortLinesInArrays(lines)
	distance := calculator.CalculateDistance(leftArray[:], rightArray[:])
	fmt.Printf("total distance between points: %v\n", distance)
	similiarity := calculator.CalculateSimiliarity(leftArray[:], rightArray[:])
	fmt.Printf("total similiarity of arrays: %v\n", similiarity)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func sortLinesInArrays(lines []string) ([1000]int, [1000]int) {
	var leftArray, rightArray [1000]int
	for i, line := range lines {
		splittedLine := strings.Split(line, "   ")
		if len(splittedLine) != 2 {
			log.Fatalf("lines are not 2 line long, the lenght %v", len(splittedLine))
		}

		leftNumber, err := strconv.Atoi(strings.TrimSpace(splittedLine[0]))
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		rightNumber, err := strconv.Atoi(strings.TrimSpace(splittedLine[1]))
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		leftArray[i] = leftNumber
		rightArray[i] = rightNumber
	}
	return leftArray, rightArray
}
