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

type Round struct {
	red, green, blue int
}

type Game struct {
	id     int
	rounds []*Round
}

func (g Game) IsValid() bool {
	for _, round := range g.rounds {
		if round.red > 12 || round.green > 13 || round.blue > 14 {
			return false
		}
	}
	return true
}

func (g Game) Power() int {
	var maximum Round

	for _, round := range g.rounds {
		if round.red > maximum.red {
			maximum.red = round.red
		}
		if round.green > maximum.green {
			maximum.green = round.green
		}
		if round.blue > maximum.blue {
			maximum.blue = round.blue
		}
	}
	return maximum.red * maximum.green * maximum.blue
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
	var sum, power int
	for scanner.Scan() {
		// Do something with the line
		line := scanner.Text()
		game := scanGame(line)
		if game.IsValid() {
			sum += game.id
		}
		power += game.Power()
	}
	fmt.Printf("Sum of valid games: %d\n", sum)
	fmt.Printf("Power of valid games: %d\n", power)

}

var scanGameIdAndData = regexp.MustCompile(`Game\s+(\d+):\s+(.*)`)
var scanColorCountAndName = regexp.MustCompile(`(\d+)\s+(.*)`)

func scanRound(data string) *Round {
	round := &Round{}
	colors := strings.Split(data, ",")
	for _, color := range colors {
		color = strings.TrimSpace(color)

		match := scanColorCountAndName.FindStringSubmatch(color)
		if len(match) != 3 {
			log.Fatal("Invalid color line:", color)
			return nil
		}

		// Parse color count
		count, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal("Invalid color count:", match[1])
			return nil
		}

		switch match[2] {
		case "red":
			round.red = count
		case "green":
			round.green = count
		case "blue":
			round.blue = count
		default:
			log.Fatal("Invalid color name:", match[2])
		}
	}
	return round
}

func scanGame(line string) *Game {
	// Scan the game id and data
	matchGame := scanGameIdAndData.FindStringSubmatch(line)
	if len(matchGame) != 3 {
		log.Fatal("Invalid game line:", line)
		return nil
	}

	// Parse the game id
	id, err := strconv.Atoi(matchGame[1])
	if err != nil {
		log.Fatal("Invalid game id:", matchGame[1])
		return nil
	}

	// Parse the rounds
	var rounds []*Round
	for _, roundData := range strings.Split(matchGame[2], ";") {
		round := scanRound(roundData)
		rounds = append(rounds, round)
	}

	// Create the game
	game := &Game{
		id:     id,
		rounds: rounds,
	}
	return game
}
