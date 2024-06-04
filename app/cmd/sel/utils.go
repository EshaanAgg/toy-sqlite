package sel

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/file/tables"
	"os"
)

func getTableSchema(tableName string, dbFile *os.File) (tables.SchemaItem, error) {
	schema := tables.GetSQLiteSchema(dbFile)

	for _, table := range schema.Tables {
		if table.TableName == tableName && table.Type == "table" {
			return table, nil
		}
	}

	return tables.SchemaItem{}, fmt.Errorf("table not found: %s in the database", tableName)
}

// Returns the indices of the columns in the table data
func getAllColIndices(tableData *tables.Table, fields []string) ([]int, error) {
	colIndices := make([]int, len(fields))

	for i, field := range fields {
		colIndex, ok := tableData.ColToIdx[field]
		if !ok {
			return nil, fmt.Errorf("column %s not found in table", field)
		}
		colIndices[i] = colIndex
	}

	return colIndices, nil
}
