package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cardIdWinningHave = regexp.MustCompile(`Card\s+(\d+):(.*)\|(.*)`)

type Card struct {
	id      int
	matches int
}

func parseCard(line string) *Card {
	match := cardIdWinningHave.FindStringSubmatch(line)
	if len(match) != 4 {
		return nil
	}

	id, _ := strconv.Atoi(match[1])

	// Create a hashset of winning numbers
	winning := make(map[int]struct{})
	for _, v := range strings.Fields(strings.TrimSpace(match[2])) {
		n, _ := strconv.Atoi(v)
		winning[n] = struct{}{}
	}

	// Count matching winning numbers
	matches := 0
	for _, v := range strings.Fields(strings.TrimSpace(match[3])) {
		n, _ := strconv.Atoi(v)
		if _, ok := winning[n]; ok {
			matches++
		}
	}

	return &Card{
		id:      id,
		matches: matches,
	}
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
	cards := make([]*Card, 0)
	cardsIdToCount := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		card := parseCard(line)
		if card == nil {
			log.Fatal("Failed to parse card")
		}
		cardsIdToCount[card.id]++
		cards = append(cards, card)
	}

	// Clone cards
	for _, card := range cards {
		if card.matches == 0 {
			continue
		}
		for n := 0; n < cardsIdToCount[card.id]; n++ {
			for i := 1; i <= card.matches; i++ {
				cardId := card.id + i
				cardsIdToCount[cardId]++
			}
		}
	}

	var sum int
	for _, count := range cardsIdToCount {
		sum += count
	}
	fmt.Printf("sum: %v\n", sum)
}
