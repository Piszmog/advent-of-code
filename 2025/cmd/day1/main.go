package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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

	part1Count := 0
	part2Count := 0
	newNum := 50
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatalln(err)
		}

		dir, dirCount, err := parseLine(string(line))
		if err != nil {
			log.Fatalln(err)
		}

		// 100 is a full rotation, so we can take the differnce to what to actually work with
		if dirCount >= 100 {
			fmt.Println(string(line), dirCount/100)
			part2Count += dirCount / 100
			dirCount = dirCount % 100
		}

		originalNum := newNum
		switch dir {
		case directionLeft:
			if newNum >= dirCount {
				newNum -= dirCount
			} else {
				newNum = 100 - (dirCount - newNum)
			}
		case directionRight:
			newNum += dirCount
			if newNum > 99 {
				newNum = newNum - 100
			}
		default:
			log.Fatalln("unknown direction: ", dir)
		}

		if newNum == 0 {
			part1Count++
			part2Count++
		} else if dir == directionLeft && originalNum < newNum && originalNum != 0 {
			part2Count++
		} else if dir == directionRight && newNum < originalNum && originalNum != 0 {
			part2Count++
		}
	}

	fmt.Println("Part 1:", part1Count)
	fmt.Println("Part 2:", part2Count)
}

func parseLine(line string) (direction, int, error) {
	dir := line[0]
	num, err := strconv.Atoi(line[1:])

	return direction(dir), num, err
}

type direction byte

const (
	directionLeft  direction = 'L'
	directionRight direction = 'R'
)
