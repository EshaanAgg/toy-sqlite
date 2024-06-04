package tables

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/file"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql/stmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/utils/encoding"
	"os"
)

type Record struct {
	IntVal    int64
	StrVal    string
	DoubleVal float64
	Null      bool
}

type Table struct {
	Name     string
	Columns  []stmt.Column
	Records  [][]Record
	ColToIdx map[string]int
}

func NewTableFromRootPage(rootPage uint, dbFile *os.File, pageSize uint, tableName string, colDefs []stmt.Column) (*Table, error) {
	page, err := file.ParsePage(dbFile, rootPage*pageSize, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing root page for table: %w", err)
	}

	rows := make([][]Record, 0)

	for _, cell := range page.LTCells {
		if len(cell.Rows) != 1 {
			return nil, fmt.Errorf("expected exactly one row in cell %s", cell.Debug())
		}

		r := cell.Rows[0]
		if len(r) != len(colDefs) {
			return nil, fmt.Errorf("row length does not match column length in cell %s", cell.Debug())
		}

		row := make([]Record, len(colDefs))

		for ind, c := range r {
			// If the column is a INTEGER PRIMARY KEY, then the value is the RowID of the cell
			if colDefs[ind].Type == "INTEGER" && colDefs[ind].PrimaryKey {
				row[ind] = Record{
					IntVal: int64(cell.RowID),
				}
				continue
			}

			row[ind] = Record{
				IntVal:    c.ValInt,
				StrVal:    c.ValString,
				DoubleVal: c.ValDouble,
				Null:      c.Type == encoding.Null,
			}
		}

		rows = append(rows, row)
	}

	// Create a map of column names to their index in the row
	colToIdx := make(map[string]int)
	for i, col := range colDefs {
		colToIdx[col.Name] = i
	}

	return &Table{
		Name:     tableName,
		Columns:  colDefs,
		Records:  rows,
		ColToIdx: colToIdx,
	}, nil
}
