package utils

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
