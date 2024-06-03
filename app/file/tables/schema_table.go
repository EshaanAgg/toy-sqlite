package tables

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/file"
	"os"
)

type SchemaItem struct {
	Type      string
	Name      string
	TableName string
	RootPage  uint
	SQL       string
}

func GetSQLiteSchema(dbFile *os.File) []SchemaItem {
	rootPage, err := file.ParsePage(dbFile, 0, true)
	if err != nil {
		fmt.Println("Failed to parse root page: ", err)
		os.Exit(1)
	}

	schemaItems := make([]SchemaItem, 0)

	for _, cell := range rootPage.LTCells {
		for _, row := range cell.Rows {
			if len(row) != 5 {
				panic(fmt.Sprintf("Expected 5 columns in schema table record, but got %d in cell %s", len(row), cell.Debug()))
			}

			schemaItem := SchemaItem{
				Type:      row[0].ValString,
				Name:      row[1].ValString,
				TableName: row[2].ValString,
				RootPage:  uint(row[3].ValInt),
				SQL:       row[4].ValString,
			}

			schemaItems = append(schemaItems, schemaItem)
		}
	}

	return schemaItems
}
