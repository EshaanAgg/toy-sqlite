package sel

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/file/tables"
	"os"
)

func getTableRootPage(tableName string, dbFile *os.File) (uint, error) {
	schema := tables.GetSQLiteSchema(dbFile)

	for _, table := range schema.Tables {
		if table.TableName == tableName && table.Type == "table" {
			return table.RootPage, nil
		}
	}

	return 0, fmt.Errorf("table not found: %s in the database", tableName)
}
