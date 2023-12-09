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

	// Scan the matrix character by character
	var sum int
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			fmt.Printf("%c", matrix[r][c])
		}
		fmt.Printf("    ")
		number, included := 0, false
		for c := 0; c < len(matrix[r]); c++ {
			// We're only interested in digits
			if matrix[r][c] < '0' || matrix[r][c] > '9' {
				if number > 0 && included {
					fmt.Printf("%d ", number)
					sum += number
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
				included = checkIfIncluded(matrix, r, c)
			}
		}

		if number > 0 && included {
			fmt.Printf("%d ", number)
			sum += number
		}

		fmt.Println()
	}

	fmt.Printf("Sum: %d\n", sum)
}

func checkIfIncluded(matrix [][]byte, r int, c int) bool {
	// check all positions around the current position to see if it "touches" a special character
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) {
				continue
			}
			if isSymbol(matrix[i][j]) {
				return true
			}
		}
	}
	return false
}

func isSymbol(b byte) bool {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return false
	case '.':
		return false
	}
	return true
}
