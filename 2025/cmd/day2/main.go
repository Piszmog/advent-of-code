package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var count uint = 0
	seq := bytes.SplitSeq(b, []byte(","))
	for part := range seq {
		ids := bytes.SplitN(bytes.TrimRight(part, "\n"), []byte("-"), 2)
		left, err := byteToUint(ids[0])
		if err != nil {
			log.Fatalln(err)
		}
		right, err := byteToUint(ids[1])
		if err != nil {
			log.Fatalln(err)
		}

		for i := left; i <= right; i++ {
			numStr := strconv.FormatUint(uint64(i), 10)
			length := len(numStr)
			if length%2 > 0 {
				continue
			}

			leftStr := numStr[0 : length/2]
			rightStr := numStr[length/2:]
			if leftStr == rightStr {
				count += i
			}
		}
	}

	fmt.Println("Part 1:", count)
}

func byteToUint(input []byte) (uint, error) {
	var result uint

	for _, v := range input {
		if v < '0' || v > '9' {
			return 0, fmt.Errorf("invalid digit: %c", v)
		}

		newResult := result*10 + uint(v-'0')
		if newResult < result {
			return 0, errors.New("overflow")
		}
		result = newResult
	}

	return result, nil
}
