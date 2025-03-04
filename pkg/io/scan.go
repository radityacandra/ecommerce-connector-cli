package io

import (
	"bufio"
	"io"
	"strconv"
)

func scanBytes(r io.Reader) []byte {
	reader := bufio.NewReader(r)
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		return nil
	}
	bytes = bytes[:len(bytes)-1]
	return bytes
}

func ScanString(r io.Reader) string {
	bytes := scanBytes(r)

	return string(bytes)
}

func ScanInt(r io.Reader) int {
	bytes := scanBytes(r)
	digit, _ := strconv.Atoi(string(bytes))

	return digit
}
