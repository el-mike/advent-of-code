package common

import (
	"bufio"
	"os"
)

func GetFileScanner(filePath string) (*bufio.Scanner, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(f), nil
}
