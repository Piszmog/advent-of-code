package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	total1 := 0
	total2 := 0

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
		largest := make([]position, 12)

		for i, b := range line {
			num := byteToInt(b)

			if num > left && i < len(line)-1 {
				left = num
				right = byteToInt(line[i+1])
			} else if num > right {
				right = num
			}

			if num > largest[0].value && i < len(line)-11 {
				update(largest, 0, i, num)
			} else if num > largest[1].value && i < len(line)-10 {
				update(largest, 1, i, num)
			} else if num > largest[2].value && i < len(line)-9 {
				update(largest, 2, i, num)
			} else if num > largest[3].value && i < len(line)-8 {
				update(largest, 3, i, num)
			} else if num > largest[4].value && i < len(line)-7 {
				update(largest, 4, i, num)
			} else if num > largest[5].value && i < len(line)-6 {
				update(largest, 5, i, num)
			} else if num > largest[6].value && i < len(line)-5 {
				update(largest, 6, i, num)
			} else if num > largest[7].value && i < len(line)-4 {
				update(largest, 7, i, num)
			} else if num > largest[8].value && i < len(line)-3 {
				update(largest, 8, i, num)
			} else if num > largest[9].value && i < len(line)-2 {
				update(largest, 9, i, num)
			} else if num > largest[10].value && i < len(line)-1 {
				update(largest, 10, i, num)
			} else if num > largest[11].value {
				update(largest, 11, i, num)
			}
		}

		total1 += (left*10 + right)
		total2 += combine(largest)
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}

func byteToInt(v byte) int {
	return int(v - '0')
}

func update(arr []position, pos int, idx int, val int) {
	existingVal := arr[pos].value

	arr[pos].index = idx
	arr[pos].value = val

	if existingVal > 0 {
		zeroValues(arr, pos, idx)
	}
}

func zeroValues(arr []position, pos int, invalidIndex int) {
	for i := pos; i < len(arr); i++ {
		p := arr[i]
		if p.index < invalidIndex {
			p.index = 0
			p.value = 0
			arr[i] = p
		}
	}
}

func combine(arr []position) int {
	total := 0
	for i := len(arr) - 1; i >= 0; i-- {
		total += arr[i].value * int(math.Pow10(len(arr)-i-1))
	}
	return total
}

type position struct {
	index int
	value int
}
