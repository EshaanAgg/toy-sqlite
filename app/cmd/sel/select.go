package sel

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/defs"
	"github/com/codecrafters-io/sqlite-starter-go/app/file/tables"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql/stmt"
	"strings"
)

func HandleSelectCommand(cmdData *defs.CommandData, selectStmt *stmt.SelectStatement) {
	tableData, err := getTableData(cmdData, selectStmt)
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}

	// Handle the COUNT clause
	if len(selectStmt.Fields) == 1 && selectStmt.Fields[0].Type == "COUNT" {
		count, err := getCount(tableData, selectStmt.Fields)
		if err != nil {
			fmt.Println("ERR: ", err)
			return
		}

		fmt.Println(count)
		return
	}

	fmt.Println("Not implemented")
}

// Returns all the data associated with the table in the database file
func getTableData(cmdData *defs.CommandData, selectStmt *stmt.SelectStatement) (*tables.Table, error) {
	// Get the table schema
	tableSchema, err := getTableSchema(selectStmt.Table, cmdData.DatabaseFile)
	if err != nil {
		return nil, fmt.Errorf("error getting table root page: %w", err)
	}

	// Parse the table's CREATE TABLE statement to get the column definitions
	lexer := sql.NewLexer(tableSchema.SQL)
	createStatement, err := stmt.ParseCreateTableStatement(lexer)
	if err != nil {
		return nil, fmt.Errorf("error parsing CREATE TABLE statement: %w", err)
	}

	tableData, err := tables.NewTableFromRootPage(
		tableSchema.RootPage-1,
		cmdData.DatabaseFile,
		uint(cmdData.Header.PageSize),
		createStatement.TableName,
		createStatement.Columns,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting table data: %w", err)
	}

	return tableData, nil
}

// Returns the number of rows in the table according to the COUNT clause in the SELECT statement
func getCount(tableData *tables.Table, fields []stmt.Field) (uint, error) {
	if len(fields) != 1 {
		return 0, fmt.Errorf("expected exactly one field in COUNT statement, got %d", len(fields))
	}

	field := fields[0]

	if field.Name == "ALL" {
		return uint(len(tableData.Records)), nil
	}

	if field.Name == "COLS" {
		colNames := strings.Split(field.Metadata, ",")
		colIndices, err := getAllColIndices(tableData, colNames)
		if err != nil {
			return 0, err
		}

		count := 0
		for _, row := range tableData.Records {
			toCount := false
			for _, colIndex := range colIndices {
				if !row[colIndex].Null {
					toCount = true
					break
				}
			}

			if toCount {
				count++
			}
		}

		return uint(count), nil
	}

	return 0, fmt.Errorf("invalid field type in COUNT statement: %s", field.Type)
}
