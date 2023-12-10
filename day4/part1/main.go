package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cardIdWinningHave = regexp.MustCompile(`Card\s+\d+:(.*)\|(.*)`)

func computeCardPoints(line string) int {
	match := cardIdWinningHave.FindStringSubmatch(line)
	if len(match) != 3 {
		return 0
	}

	haveSlice := strings.Fields(strings.TrimSpace(match[2]))

	// Create a hashset of winning numbers
	winning := make(map[int]struct{})
	for _, v := range strings.Fields(strings.TrimSpace(match[1])) {
		n, _ := strconv.Atoi(v)
		winning[n] = struct{}{}
	}
	//fmt.Printf("Winning: %v\n", winning)

	// Check each number we have against the winning numbers
	count := 0
	for _, v := range haveSlice {
		n, _ := strconv.Atoi(v)
		if _, ok := winning[n]; ok {
			count++
		}
	}
	//fmt.Printf("Count: %v\n", count)

	points := int(math.Pow(2, float64(count-1)))
	//fmt.Printf("Points for card: %v\n", points)

	return points
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
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		points := computeCardPoints(line)
		sum += points
	}

	fmt.Printf("Total points: %v\n", sum)
}
