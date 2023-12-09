package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func startsWith(s, sub string) bool {
	return len(s) >= len(sub) && s[:len(sub)] == sub
}

func startsWithSpelledOutDigit(line string, i int) (bool, int) {
	for key, value := range digits {
		if startsWith(line[i:], key) {
			return true, value
		}
	}
	return false, 0
}

func getCalibrationValue(line string) int {
	var first, last int

	// Find the first digit
	for i := 0; i < len(line); i++ {
		if '0' <= line[i] && line[i] <= '9' {
			first = int(line[i] - '0')
			break
		}

		if starts, digit := startsWithSpelledOutDigit(line, i); starts {
			first = digit
			break
		}
	}

	// Find the last digit
	for i := len(line) - 1; i >= 0; i-- {
		if '0' <= line[i] && line[i] <= '9' {
			last = int(line[i] - '0')
			break
		}

		if starts, digit := startsWithSpelledOutDigit(line, i); starts {
			last = digit
			break
		}
	}

	return first*10 + last
}

func main() {
	// Open the file
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Scan each line
	sum := 0
	for scanner.Scan() {
		// Do something with the line
		line := scanner.Text()
		calibrationValue := getCalibrationValue(line)
		sum += calibrationValue
	}
	fmt.Println("sum:", sum)
}
