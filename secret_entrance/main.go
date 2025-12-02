package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	currentPos := 50
	totalZeros := 0

	fmt.Println("Advent of Code 2025 Day 1, Part Two")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		direction := line[0]
		clicks, _ := strconv.Atoi(line[1:])

		zerosInMove, newPos := processRotation(currentPos, clicks, direction)

		totalZeros += zerosInMove
		currentPos = newPos
	}

	fmt.Printf("Final Password: %d\n", totalZeros)
}

func processRotation(startPos int, clicks int, direction byte) (int, int) {
	count := 0

	count += clicks / 100

	remainder := clicks % 100

	if direction == 'R' {
		if startPos+remainder >= 100 {
			count++
		}

		startPos = (startPos + remainder) % 100

	} else {
		if startPos > 0 && remainder >= startPos {
			count++
		}

		// Calculate new position (handle negative wrapping)
		startPos = (startPos - remainder) % 100
		if startPos < 0 {
			startPos += 100
		}
	}

	return count, startPos
}
