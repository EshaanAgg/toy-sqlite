package dot

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/defs"
	"github/com/codecrafters-io/sqlite-starter-go/app/file/tables"
	"strings"
)

func Tables(commandData *defs.CommandData) {
	tableNames := make([]string, 0)

	schemaTable := tables.GetSQLiteSchema(commandData.DatabaseFile)
	for _, schemaItem := range schemaTable {
		if schemaItem.Type == "table" && !strings.HasPrefix(schemaItem.Name, "sqlite_") {
			tableNames = append(tableNames, schemaItem.Name)
		}
	}

	fmt.Println(strings.Join(tableNames, "\t"))
}
