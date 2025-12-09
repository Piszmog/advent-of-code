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
