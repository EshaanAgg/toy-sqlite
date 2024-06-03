package utils

import (
	"encoding/binary"
	"fmt"
)

func ReadUInt8(data []byte) uint8 {
	if len(data) < 1 {
		panic("Not enough data to read uint8")
	}
	return data[0]
}

func ReadUInt16(data []byte) uint16 {
	if len(data) < 2 {
		panic("Not enough data to read uint16")
	}
	return uint16(data[0])<<8 | uint16(data[1])
}

func ReadUInt32(data []byte) uint32 {
	if len(data) < 4 {
		panic("Not enough data to read uint32")
	}

	val := uint32(0)
	for i := 0; i < 4; i++ {
		offset := 8 * (3 - i)
		val |= uint32(data[i]) << uint(offset)
	}
	return val
}

func ReadTwosComplement(b []byte) int64 {
	length := len(b)
	if length != 1 && length != 2 && length != 3 && length != 4 && length != 6 && length != 8 {
		panic(fmt.Sprintf("byte slice must be 1, 2, 3, 4, 6, or 8 bytes long, but got %d", length))
	}

	var result int64

	switch length {
	case 1:
		result = int64(int8(b[0]))

	case 2:
		result = int64(int16(binary.BigEndian.Uint16(b)))

	case 3:
		// Sign extend to 32 bits
		if b[0]&0b10000000 != 0 {
			b = append([]byte{0xFF}, b...)
		} else {
			b = append([]byte{0x00}, b...)
		}
		result = int64(int32(binary.BigEndian.Uint32(b)))

	case 4:
		result = int64(int32(binary.BigEndian.Uint32(b)))

	case 6:
		// Sign extend to 64 bits
		if b[0]&0b10000000 != 0 {
			b = append([]byte{0xFF, 0xFF}, b...)
		} else {
			b = append([]byte{0x00, 0x00}, b...)
		}
		result = int64(binary.BigEndian.Uint64(b))

	case 8:
		result = int64(binary.BigEndian.Uint64(b))

	default:
		panic(fmt.Sprintf("unreachable: invalid length %d", length))
	}

	return result
}
