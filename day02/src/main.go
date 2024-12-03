package main

import (
	"commons/reader"
	"day02/counter"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines, err := reader.ReadLines("../assets/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	allReports := parseLinesToReports(lines)
	safe := counter.CountSafeReports(allReports)
	fmt.Printf("Number of safe reports %v\n", safe)
}

func parseLinesToReports(lines []string) [][]int {
	var allReports [][]int
	for _, line := range lines {
		splittedLine := strings.Split(line, " ")
		var levels = make([]int, len(splittedLine))
		for i, level := range splittedLine {
			levelInt, err := strconv.Atoi(level)
			if err != nil {
				log.Fatalf("level is not integer %s", level)
			}
			levels[i] = levelInt
		}
		allReports = append(allReports, levels)
	}
	return allReports
}
