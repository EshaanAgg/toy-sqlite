package cells

import (
	"github/com/codecrafters-io/sqlite-starter-go/app/utils"
	"github/com/codecrafters-io/sqlite-starter-go/app/utils/encoding"
	"os"
)

type LeafTableCell struct {
	TotalPayloadBytes uint
	RowID             uint
	Payload           []byte
	FirstOverflowPage uint32
}

func (cell *LeafTableCell) Parse(file *os.File, startIndex uint) {
	cell.TotalPayloadBytes, startIndex = encoding.ReadVarInt(file, startIndex)
	cell.RowID, startIndex = encoding.ReadVarInt(file, startIndex)

	// TODO: Parse the payload and overflow pages
	// For now, we just read the payload and ignore the overflow pages
	cell.Payload = utils.ReadFile(file, startIndex, int(cell.TotalPayloadBytes))
	cell.FirstOverflowPage = 0
}
