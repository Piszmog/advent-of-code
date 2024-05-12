package utils

import (
	"bufio"
	"encoding/csv"
	"io"
)

func ReadLines(r io.Reader, lineChan chan<- string) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lineChan <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	close(lineChan)
	return nil
}

func ReadCSV(r io.Reader, lineChan chan<- []string) error {
	reader := csv.NewReader(r)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		lineChan <- line
	}

	close(lineChan)
	return nil
}
