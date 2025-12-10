package reader

import (
	"bufio"
	"errors"
	"io"
)

func ReadLines(r io.Reader, handlerFunc ReadLineHandler) error {
	reader := bufio.NewReader(r)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		err = handlerFunc(line)
		if err != nil {
			return err
		}
	}

	return nil
}

type ReadLineHandler func(line []byte) error
