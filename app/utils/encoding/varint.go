package encoding

import (
	"github/com/codecrafters-io/sqlite-starter-go/app/utils"
	"os"
)

// Reads a variable length integer from the given file at the given start index.
// Returns the parsed integer and the updated index after the integer.
func ReadVarInt(file *os.File, startIndex uint) (uint, uint) {
	byte := utils.ReadFile(file, startIndex, 1)[0]

	value := uint(byte & 0b01111111)
	hasMoreBytes := byte & 0b10000000
	startIndex++

	for hasMoreBytes > 0 {
		byte = utils.ReadFile(file, startIndex, 1)[0]
		readValue := uint(byte & 0b01111111)
		value = value<<7 | readValue
		hasMoreBytes = byte & 0b10000000
		startIndex++
	}

	return value, startIndex
}

func ReadVarIntFromBytes(bytes []byte) (uint, uint) {
	startIndex := uint(0)
	byte := bytes[startIndex]

	value := uint(byte & 0b01111111)
	hasMoreBytes := byte & 0b10000000
	startIndex++

	for hasMoreBytes > 0 {
		byte = bytes[startIndex]
		readValue := uint(byte & 0b01111111)
		value = value<<7 | readValue
		hasMoreBytes = byte & 0b10000000
		startIndex++
	}

	return value, startIndex
}
