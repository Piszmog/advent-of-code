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

	total := 0

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatalln(err)
		}

		left := 0
		right := 0

		for i, b := range line {
			num := byteToInt(b)

			if num > left && i < len(line)-1 {
				left = num
				right = byteToInt(line[i+1])
			} else if num > right {
				right = num
			}
		}

		total += (left*10 + right)
	}

	fmt.Println("Part 1:", total)
}

func byteToInt(v byte) int {
	return int(v - '0')
}
