package cells

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/utils"
	"github/com/codecrafters-io/sqlite-starter-go/app/utils/encoding"
	"os"
	"strings"
)

type LeafTableCell struct {
	TotalPayloadBytes uint
	RowID             uint
	Payload           []byte
	FirstOverflowPage uint32
	Rows              [][]encoding.RecordField
}

func (cell *LeafTableCell) Parse(file *os.File, startIndex uint) {
	cell.TotalPayloadBytes, startIndex = encoding.ReadVarInt(file, startIndex)
	cell.RowID, startIndex = encoding.ReadVarInt(file, startIndex)

	// TODO: Parse the payload and overflow pages
	// For now, we just read the payload and ignore the overflow pages
	cell.Payload = utils.ReadFile(file, startIndex, int(cell.TotalPayloadBytes))
	cell.Rows = encoding.ReadPayload(cell.Payload)
	cell.FirstOverflowPage = 0
}

func (cell *LeafTableCell) Debug() string {
	rowStrings := make([]string, len(cell.Rows))
	for rowIndex, row := range cell.Rows {
		recordStrings := make([]string, len(row))
		for recordIndex, record := range row {
			recordStrings[recordIndex] = record.Debug()
		}
		rowStrings[rowIndex] = fmt.Sprintf("[%s]", strings.Join(recordStrings, ", "))
	}

	return fmt.Sprintf(
		`LeafTableCell {
	RowID: %d,
	Records: %s,
}`,
		cell.RowID,
		fmt.Sprintf("(%s)", strings.Join(rowStrings, "\n")),
	)
}
