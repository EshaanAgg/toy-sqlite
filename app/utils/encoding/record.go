package encoding

import (
	"encoding/binary"
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/utils"
	"math"
)

type ItemType int

const (
	Int ItemType = iota
	Double
	String
	Null
)

type RecordField struct {
	ContentSize uint
	ValInt      int64
	ValString   string
	ValDouble   float64
	Type        ItemType

	serialType uint
}

func (r *RecordField) Debug() string {
	switch r.Type {
	case Int:
		return fmt.Sprintf("%d", r.ValInt)
	case Double:
		return fmt.Sprintf("%f", r.ValDouble)
	case String:
		return r.ValString
	case Null:
		return "NULL"
	default:
		return "Unknown"
	}

}

func newRecordFielFromSerialType(serialType uint) RecordField {
	record := RecordField{
		serialType: serialType,
	}

	switch serialType {
	case 0:
		record.Type = Null

	case 1:
		record.Type = Int
		record.ContentSize = 1

	case 2:
		record.Type = Int
		record.ContentSize = 2

	case 3:
		record.Type = Int
		record.ContentSize = 3

	case 4:
		record.Type = Int
		record.ContentSize = 4

	case 5:
		record.Type = Int
		record.ContentSize = 6

	case 6:
		record.Type = Int
		record.ContentSize = 8

	case 7:
		record.Type = Double
		record.ContentSize = 8

	case 8:
		record.Type = Int
		record.ContentSize = 0
		record.ValInt = 0

	case 9:
		record.Type = Int
		record.ContentSize = 0
		record.ValInt = 1

	case 10, 11:
		panic("Reserved for internal SQLite use. Should not be used in user databases.")

	default:
		record.Type = String
		if serialType%2 == 0 {
			record.ContentSize = (serialType - 12) / 2
		} else {
			record.ContentSize = (serialType - 13) / 2
		}
	}

	return record
}

func (r *RecordField) readContent(content []byte) {
	if uint(len(content)) != r.ContentSize {
		panic(fmt.Sprintf("Expected content size of %d but got %d while parsing records from cell content.", r.ContentSize, len(content)))
	}

	switch r.Type {
	case Int:
		if r.ContentSize != 0 {
			r.ValInt = utils.ReadTwosComplement(content)
		}
	case Double:
		r.ValDouble = math.Float64frombits(binary.BigEndian.Uint64(content))
	case String:
		r.ValString = string(content)
	case Null:
		// Do nothing

	default:
		panic(fmt.Sprintf("Unknown record type %d while parsing records from cell content.", r.Type))
	}
}

func parseRecords(content []byte, startIndex uint) ([]RecordField, uint) {
	// Read the header size
	index := startIndex
	headerSize, offset := ReadVarIntFromBytes(content[index:])
	index += offset

	// Read the header
	records := make([]RecordField, 0)
	for index < headerSize {
		serialType, offset := ReadVarIntFromBytes(content[index:])
		index += offset
		records = append(records, newRecordFielFromSerialType(serialType))
	}

	// Read the content
	for recordIndex, record := range records {
		contentData := content[index : index+record.ContentSize]
		records[recordIndex].readContent(contentData)
		index += record.ContentSize
	}

	return records, index
}

func ReadPayload(payload []byte) [][]RecordField {
	rows := make([][]RecordField, 0)

	index := uint(0)
	for index < uint(len(payload)) {
		records, offset := parseRecords(payload, index)
		rows = append(rows, records)
		index += offset
	}

	return rows
}
