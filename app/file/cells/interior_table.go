package cells

import (
	"github/com/codecrafters-io/sqlite-starter-go/app/utils"
	"github/com/codecrafters-io/sqlite-starter-go/app/utils/encoding"
	"os"
)

type InteriorTableCell struct {
	LeftChildPage uint32
	RowID         uint
}

func (cell *InteriorTableCell) Parse(file *os.File, startIndex uint) {
	cell.LeftChildPage = utils.ReadUInt32(
		utils.ReadFile(file, startIndex, 4),
	)

	intKey, _ := encoding.ReadVarInt(file, startIndex+4)
	cell.RowID = intKey
}
