package sel

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/file/tables"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql/stmt"
	"os"
	"strings"
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

// Prints the data supplied in accordance to the types defined by the columns.
// It is assumed that the length of each row would be equal to the length of columns, otherwise would panic.
func DisplayRecords(data [][]tables.Record, colDefs []stmt.Column) {
	for _, row := range data {
		items := make([]string, len(row))
		if len(row) != len(colDefs) {
			panic(fmt.Sprintf("The length of the supplied row %v and column definitions %v is not equal", row, colDefs))
		}

		for ind, item := range row {
			if item.Null {
				items[ind] = "null"
				continue
			}

			switch colDefs[ind].Type {
			case "INTEGER":
				items[ind] = fmt.Sprintf("%d", item.IntVal)
			case "TEXT":
				items[ind] = item.StrVal
			default:
				panic(fmt.Sprintf("Unimplemented type of column found in the table: %s for column %v", colDefs[ind].Type, colDefs[ind]))
			}
		}

		fmt.Println(strings.Join(items, "|"))
	}
}

// Returns a slice of the elements at the indices 'idx' from the 'slice'
func getIndices[T any](slice []T, idx []int) []T {
	newSlice := make([]T, 0)
	for _, i := range idx {
		newSlice = append(newSlice, slice[i])
	}
	return newSlice
}
