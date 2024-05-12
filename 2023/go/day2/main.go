package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Piszmog/advent-2023/go/utils"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	part1, part2, err := solveParts(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func solveParts(r io.Reader) (int, int, error) {
	inputChan := make(chan string, 10)
	go utils.ReadLines(r, inputChan)

	part1Result := 0
	part2Result := 0
	for input := range inputChan {
		actual, err := readLinePart1(input)
		if err != nil {
			return 0, 0, err
		}
		part1Result += actual

		actual, err = readLinePart2(input)
		if err != nil {
			return 0, 0, err
		}
		part2Result += actual
	}
	return part1Result, part2Result, nil
}

func readLinePart1(input string) (int, error) {
	inputParts := strings.Split(input, ": ")
	if len(inputParts) != 2 {
		return 0, fmt.Errorf("invalid input: %s", input)
	}
	gameID := strings.Split(inputParts[0], " ")[1]

	gameParts := strings.Split(inputParts[1], "; ")

	for _, gamePart := range gameParts {
		colors := strings.Split(gamePart, ", ")
		for _, color := range colors {
			colorParts := strings.SplitN(color, " ", 2)
			i, err := strconv.Atoi(colorParts[0])
			if err != nil {
				return 0, fmt.Errorf("invalid input: %s", input)
			}
			switch colorParts[1] {
			case "red":
				if i > 12 {
					return 0, nil
				}
			case "blue":
				if i > 14 {
					return 0, nil
				}
			case "green":
				if i > 13 {
					return 0, nil
				}
			default:
				return 0, fmt.Errorf("invalid input: %s", input)
			}
		}
	}

	id, err := strconv.Atoi(gameID)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %s", input)
	}

	return id, nil
}

func readLinePart2(input string) (int, error) {
	inputParts := strings.Split(input, ": ")
	if len(inputParts) != 2 {
		return 0, fmt.Errorf("invalid input: %s", input)
	}

	gameParts := strings.Split(inputParts[1], "; ")

	redCount := 0
	blueCount := 0
	greenCount := 0
	for _, gamePart := range gameParts {
		colors := strings.Split(gamePart, ", ")
		for _, color := range colors {
			colorParts := strings.SplitN(color, " ", 2)
			i, err := strconv.Atoi(colorParts[0])
			if err != nil {
				return 0, fmt.Errorf("invalid input: %s", input)
			}
			switch colorParts[1] {
			case "red":
				if redCount == 0 {
					redCount = i
				} else if i > redCount {
					redCount = i
				}
			case "blue":
				if blueCount == 0 {
					blueCount = i
				} else if i > blueCount {
					blueCount = i
				}
			case "green":
				if greenCount == 0 {
					greenCount = i
				} else if i > greenCount {
					greenCount = i
				}
			default:
				return 0, fmt.Errorf("invalid input: %s", input)
			}
		}
	}

	return redCount * blueCount * greenCount, nil
}
