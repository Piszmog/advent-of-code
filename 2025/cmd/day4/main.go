package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	rolls := make(map[roll]struct{})

	row := 0
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatalln(err)
		}

		for col, v := range line {
			switch v {
			case '@':
				rolls[roll{row: row, col: col}] = struct{}{}
			}
		}
		row++
	}

	fmt.Println("Part 1:", part1(rolls))
	fmt.Println("Part 2:", part2(rolls))
}

func part1(rolls map[roll]struct{}) int {
	result := 0
	for r := range rolls {
		count := countNeighbors(r, rolls)
		if count < 4 {
			result++
		}
	}
	return result
}

func part2(rolls map[roll]struct{}) int {
	rollsToCheck := rolls
	result := 0
	for {
		leftovers, taken := nextGeneration(rollsToCheck)
		if taken == 0 {
			break
		}
		rollsToCheck = leftovers
		result += taken
	}
	return result
}

func nextGeneration(rolls map[roll]struct{}) (map[roll]struct{}, int) {
	taken := 0
	newRolls := make(map[roll]struct{})

	for r := range rolls {
		count := countNeighbors(r, rolls)
		if count < 4 {
			taken++
		} else {
			newRolls[r] = struct{}{}
		}
	}
	return newRolls, taken
}

func countNeighbors(r roll, rolls map[roll]struct{}) int {
	count := 0
	for row := -1; row <= 1; row++ {
		for col := -1; col <= 1; col++ {
			if row == 0 && col == 0 {
				continue
			}
			if _, ok := rolls[roll{row: r.row + row, col: r.col + col}]; ok {
				count++
			}
		}
	}
	return count
}

type roll struct {
	row int
	col int
}
