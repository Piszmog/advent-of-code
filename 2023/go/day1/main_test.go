package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	scanner := bufio.NewScanner(strings.NewReader(input))

	actual := 0

	for scanner.Scan() {
		lineVal, err := readLinePart1(string(scanner.Text()))
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		actual += lineVal
	}

	if actual != 142 {
		t.Errorf("Expected 142, got %d", actual)
	}
}

func TestPart2(t *testing.T) {
	input := `two1nine
			eightwothree
			abcone2threexyz
			xtwone3four
			4nineeightseven2
			zoneight234
			7pqrstsixteen`
	scanner := bufio.NewScanner(strings.NewReader(input))

	actual := 0

	for scanner.Scan() {
		lineVal, err := readLinePart2(string(scanner.Text()))
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		actual += lineVal
	}

	if actual != 281 {
		t.Errorf("Expected 281, got %d", actual)
	}
}
