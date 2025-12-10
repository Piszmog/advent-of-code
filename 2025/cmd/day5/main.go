package main

import (
	"bufio"
	"bytes"
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

	part1 := 0
	var ids []idRange
	startAvailableIDs := false
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatalln(err)
		}

		if len(line) > 0 && !startAvailableIDs {
			parts := bytes.SplitN(line, []byte("-"), 2)
			left, err := strconv.Atoi(string(parts[0]))
			if err != nil {
				log.Fatalln(err)
			}
			right, err := strconv.Atoi(string(parts[1]))
			if err != nil {
				log.Fatalln(err)
			}

			update := false
			for i, id := range ids {
				if left > id.start && left < id.end && right > id.end {
					ids[i].end = right
					update = true
					break
				} else if left < id.start && right > id.start && right < id.end {
					ids[i].start = left
					update = true
					break
				}
			}
			if !update {
				ids = append(ids, idRange{start: left, end: right})
			}

		} else if len(line) == 0 && !startAvailableIDs {
			startAvailableIDs = true
		} else if len(line) > 0 && startAvailableIDs {
			num, err := strconv.Atoi(string(line))
			if err != nil {
				log.Fatalln(err)
			}

			for _, id := range ids {
				if num >= id.start && num <= id.end {
					part1++
					break
				}
			}
		}
	}

	fmt.Println("Part 1:", part1)
}

type idRange struct {
	start int
	end   int
}
