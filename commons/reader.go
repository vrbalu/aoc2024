package reader

import (
	"bufio"
	"os"
)

// ReadLines returns array of lines read from a file that is given as an input.
// Error will be returned in case of failed reading - e.g. file not found
func ReadLines(path string) ([]string, error) {
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
