package main

import (
	"commons/reader"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var totalSum int
	lines, err := reader.ReadLines("../assets/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, m := range matches {
			m = strings.TrimLeft(m, "mul(")
			m = strings.TrimRight(m, ")")
			arr := strings.Split(m, ",")
			left, err := strconv.Atoi(arr[0])
			if err != nil {
				log.Fatalf("error converting left: %s", err)
			}
			right, err := strconv.Atoi(arr[1])
			if err != nil {
				log.Fatalf("error converting right: %s", err)
			}
			totalSum += left * right
		}
	}
	log.Println(totalSum)
}
