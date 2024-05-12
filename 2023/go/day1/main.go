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
	runes := []rune(input)
	number := make([]rune, 2)
	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		if runes[i] >= '0' && runes[i] <= '9' && number[0] == 0 {
			number[0] = runes[i]
		}
		if runes[j] >= '0' && runes[j] <= '9' && number[1] == 0 {
			number[1] = runes[j]
		}
	}
	actual, err := strconv.Atoi(string(number[0]) + string(number[1]))
	if err != nil {
		panic(err)
	}
	return actual, nil
}

func readLinePart2(input string) (int, error) {
	firstNumber := ""
	firstNumberIndex := -1
	lastNumber := ""
	lastNumberIndex := -1

	for num, val := range numbers {
		if strings.Contains(input, num) {
			i := strings.Index(input, num)
			j := strings.LastIndex(input, num)
			if firstNumberIndex == -1 || i < firstNumberIndex {
				if firstNumberIndex > lastNumberIndex {
					lastNumber = firstNumber
					lastNumberIndex = firstNumberIndex
				}
				firstNumber = val
				firstNumberIndex = i
			}

			if lastNumberIndex == -1 || j > lastNumberIndex {
				lastNumber = val
				lastNumberIndex = j
			}
		}
	}

	actual, err := strconv.Atoi(firstNumber + lastNumber)
	if err != nil {
		return 0, err
	}
	return actual, nil
}

var numbers = map[string]string{
	"0":     "0",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
