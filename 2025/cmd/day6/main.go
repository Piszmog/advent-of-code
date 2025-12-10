package main

import (
	"aoc-2025/internal/reader"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	var part1Data [][]int
	var part1 []int
	err = reader.ReadLines(f, func(line []byte) error {
		parts := bytes.Fields(line)

		row := make([]int, 0, len(parts))
		for col, p := range parts {
			switch p[0] {
			case '+':
				res := 0
				for _, data := range part1Data {
					res += data[col]
				}
				part1 = append(part1, res)
			case '*':
				res := 0
				for _, data := range part1Data {
					if res == 0 {
						res = data[col]
					} else {
						res *= data[col]
					}
				}
				part1 = append(part1, res)
			default:
				num, atoiErr := strconv.Atoi(string(p))
				if atoiErr != nil {
					return atoiErr
				}
				row = append(row, num)
			}
		}

		part1Data = append(part1Data, row)
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	part1Result := 0
	for _, v := range part1 {
		part1Result += v
	}

	fmt.Println("Part 1:", part1Result)
}
