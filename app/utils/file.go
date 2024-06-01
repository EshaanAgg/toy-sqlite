package utils

import (
	"fmt"
	"os"
)

func ReadFile(file *os.File, startIndex uint, length int) []byte {
	var buf = make([]byte, length)
	_, err := file.ReadAt(buf, int64(startIndex))
	if err != nil {
		panic(fmt.Errorf("failed to read the database file: %v", err))
	}

	return buf
}
