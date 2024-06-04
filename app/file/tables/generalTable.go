package tables

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/file"
	"os"
)

type Column struct {
	Name          string
	Type          string
	PrimaryKey    bool
	AutoIncrement bool
}

type Record struct {
	IntVal    int64
	StrVal    string
	DoubleVal float64
}

type Table struct {
	Name    string
	Columns []Column
	Records [][]Record
}

func NewTableFromRootPage(rootPage uint, dbFile *os.File, pageSize uint, tableName string, colDefs []Column) (*Table, error) {
	page, err := file.ParsePage(dbFile, rootPage*pageSize, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing root page for table: %w", err)
	}

	rows := make([][]Record, 0)

	for _, cell := range page.LTCells {
		for _, r := range cell.Rows {
			row := make([]Record, 0)
			for _, c := range r {
				row = append(row, Record{
					IntVal:    c.ValInt,
					StrVal:    c.ValString,
					DoubleVal: c.ValDouble,
				})
			}
			rows = append(rows, row)
		}
	}

	return &Table{
		Name:    tableName,
		Columns: colDefs,
		Records: rows,
	}, nil
}
