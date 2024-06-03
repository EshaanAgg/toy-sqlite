package dot

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/defs"
	"github/com/codecrafters-io/sqlite-starter-go/app/file/tables"
)

func DBInfo(cmdData *defs.CommandData) {
	h := cmdData.Header

	text_encoding_map := map[uint32]string{
		1: "utf8",
		2: "utf16le",
		3: "utf16be",
	}

	fmt.Println("database page size:", h.PageSize)
	fmt.Println("write format:", h.WriteFormat)
	fmt.Println("read format:", h.ReadFormat)
	fmt.Println("reserved bytes:", h.ReservedBytes)
	fmt.Println("file change counter:", h.FileChangeCounter)
	fmt.Println("database page count:", h.NumberPages)
	fmt.Println("freelist page count:", h.TotalFreelistPages)
	fmt.Println("schema cookie:", h.SchemaCookie)
	fmt.Println("schema format:", h.SchemaFormatNumber)
	fmt.Println("default cache size:", h.DefaultCacheSize)
	fmt.Println("autovacuum top root:", h.LargestRootBTreePage)
	fmt.Printf("text encoding: %d (%s)\n", h.TextEncoding, text_encoding_map[h.TextEncoding])
	fmt.Println("user version:", h.UserVersion)
	fmt.Println("application id:", h.ApplicationID)

	schemaTable := tables.GetSQLiteSchema(cmdData.DatabaseFile)
	numberTables := 0
	for _, schemaItem := range schemaTable {
		if schemaItem.Type == "table" {
			numberTables++
		}
	}
	fmt.Println("number of tables:", numberTables)
}
