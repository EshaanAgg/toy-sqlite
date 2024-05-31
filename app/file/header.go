package file

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/utils"
	"os"
)

type DatabaseHeader struct {
	PageSize               uint16
	WriteFormat            uint8
	ReadFormat             uint8
	ReservedBytes          uint8
	FileChangeCounter      uint32
	NumberPages            uint32
	FirstFreelistTrunkPage uint32
	TotalFreelistPages     uint32
	SchemaCookie           uint32
	SchemaFormatNumber     uint32
	DefaultCacheSize       uint32
	LargestRootBTreePage   uint32
	TextEncoding           uint32
	UserVersion            uint32
	IncrementalVacuumMode  bool
	ApplicationID          uint32
	VersionValidFor        uint32
	SQLiteVersion          string
}

func ParseDatabaseHeader(databaseFile *os.File) (*DatabaseHeader, error) {
	var data = make([]byte, 100)
	_, err := databaseFile.ReadAt(data, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to read database file: %v", err)
	}

	if len(data) < 100 {
		return nil, fmt.Errorf("invalid database file header: should be at least 100 bytes, is %d bytes", len(data))
	}

	if string(data[0:16]) != "SQLite format 3\000" {
		return nil, fmt.Errorf("invalid database file header: not a SQLite database file")
	}

	header := DatabaseHeader{}

	// Parse the database header fields
	header.PageSize = utils.ReadUInt16(data[16:18])
	header.WriteFormat = utils.ReadUInt8(data[18:19])
	header.ReadFormat = utils.ReadUInt8(data[19:20])
	header.ReservedBytes = utils.ReadUInt8(data[20:21])

	// Handle payload values
	maxPayloadFraction := utils.ReadUInt8(data[21:22])
	if maxPayloadFraction != 64 {
		return nil, fmt.Errorf("invalid database file header: unsupported max payload fraction %d", maxPayloadFraction)
	}
	minPayloadFraction := utils.ReadUInt8(data[22:23])
	if minPayloadFraction != 32 {
		return nil, fmt.Errorf("invalid database file header: unsupported min payload fraction %d", minPayloadFraction)
	}
	leafPayloadFraction := utils.ReadUInt8(data[23:24])
	if leafPayloadFraction != 32 {
		return nil, fmt.Errorf("invalid database file header: unsupported leaf payload fraction %d", leafPayloadFraction)
	}

	// Continue parsing the database header fields
	header.FileChangeCounter = utils.ReadUInt32(data[24:28])
	header.NumberPages = utils.ReadUInt32(data[28:32])
	header.FirstFreelistTrunkPage = utils.ReadUInt32(data[32:36])
	header.TotalFreelistPages = utils.ReadUInt32(data[36:40])
	header.SchemaCookie = utils.ReadUInt32(data[40:44])

	// Check for valid schema format number
	header.SchemaFormatNumber = utils.ReadUInt32(data[44:48])
	if header.SchemaFormatNumber < 1 || header.SchemaFormatNumber > 4 {
		return nil, fmt.Errorf("invalid database file header: unsupported schema format number %d", header.SchemaFormatNumber)
	}

	header.DefaultCacheSize = utils.ReadUInt32(data[48:52])
	header.LargestRootBTreePage = utils.ReadUInt32(data[52:56])
	header.TextEncoding = utils.ReadUInt32(data[56:60])
	header.UserVersion = utils.ReadUInt32(data[60:64])
	header.IncrementalVacuumMode = utils.ReadUInt32(data[64:68]) != 0
	header.ApplicationID = utils.ReadUInt32(data[68:72])

	// Check that the bits 72-91 are all zero (reserved for expansion)
	for i := 72; i < 92; i++ {
		if data[i] != 0 {
			return nil, fmt.Errorf("invalid database file header: reserved bytes are not zero")
		}
	}

	header.VersionValidFor = utils.ReadUInt32(data[92:96])
	header.SQLiteVersion = string(data[96:100])

	return &header, nil
}
