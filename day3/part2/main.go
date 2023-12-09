package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
	matrix := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []byte(line))
	}

	// Gears: coordinates -> slice of numbers touching the gear
	gears := map[[2]int][]int{}

	// Scan the matrix character by character
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			fmt.Printf("%c", matrix[r][c])
		}
		fmt.Printf("    ")
		number, included := 0, false
		gearR, gearC := 0, 0
		for c := 0; c < len(matrix[r]); c++ {
			// We're only interested in digits
			if matrix[r][c] < '0' || matrix[r][c] > '9' {
				if number > 0 && included {
					fmt.Printf("%d(%d:%d) ", number, gearR, gearC)
					gears[[2]int{gearR, gearC}] = append(gears[[2]int{gearR, gearC}], number)
				}
				number, included = 0, false
				continue
			}

			// Convert the current character to a digit and computer the number as we go
			digit := int(matrix[r][c] - '0')
			number *= 10
			number += digit

			// check if the number is included in the matrix
			if !included {
				included, gearR, gearC = checkIfGear(matrix, r, c)
			}
		}

		if number > 0 && included {
			fmt.Printf("%d (%d:%d)", number, gearR, gearC)
			gears[[2]int{gearR, gearC}] = append(gears[[2]int{gearR, gearC}], number)
		}

		fmt.Println()
	}

	var sum int
	for _, v := range gears {
		if len(v) != 2 {
			continue
		}
		product := 1
		for _, n := range v {
			product *= n
		}
		sum += product
	}
	fmt.Printf("Sum: %d\n", sum)
}

func checkIfGear(matrix [][]byte, r int, c int) (bool, int, int) {
	// check all positions around the current position to see if it "touches" a gear
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) {
				continue
			}
			if matrix[i][j] == '*' {
				return true, i, j
			}
		}
	}
	return false, 0, 0
}
