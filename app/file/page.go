package file

import (
	"fmt"
	cellUtils "github/com/codecrafters-io/sqlite-starter-go/app/file/cells"
	"github/com/codecrafters-io/sqlite-starter-go/app/utils"
	"os"
)

type PageType int

const (
	InteriorIndexPage PageType = iota
	InteriorTablePage
	LeafIndexPage
	LeafTablePage
)

func (t PageType) IsInterior() bool {
	return t == InteriorIndexPage || t == InteriorTablePage
}

func (t PageType) IsLeaf() bool {
	return t == LeafIndexPage || t == LeafTablePage
}

func (t PageType) GetHeaderSize() uint {
	if t.IsInterior() {
		return 12
	}
	return 8
}

type PageHeader struct {
	Type                PageType
	StartFreeBlock      uint16
	NumberCells         uint16
	CellContentOffset   uint
	FragmentedFreeBytes uint8
	RightMostPointer    uint32
}

// Reads the page header from the given file at the given start index.
// Returns the parsed page header, index of the first byte after the header and an error if one occurred.
func ParsePageHeader(file *os.File, startIndex uint) (*PageHeader, uint, error) {
	buf := utils.ReadFile(file, startIndex, 8)
	header := PageHeader{}

	// Get the type of the page
	switch buf[0] {
	case 2:
		header.Type = InteriorIndexPage
	case 5:
		header.Type = LeafIndexPage
	case 10:
		header.Type = InteriorTablePage
	case 13:
		header.Type = LeafTablePage
	default:
		return nil, 0, fmt.Errorf("invalid page header: unknown page type %d", buf[0])
	}

	header.StartFreeBlock = utils.ReadUInt16(buf[1:3])
	header.NumberCells = utils.ReadUInt16(buf[3:5])

	// The cell content offset is interpreted as 65536 if it is 0
	header.CellContentOffset = uint(utils.ReadUInt16(buf[5:7]))
	if header.CellContentOffset == 0 {
		header.CellContentOffset = 65536
	}

	header.FragmentedFreeBytes = utils.ReadUInt8(buf[7:8])

	// Read the right-most pointer if the page is an interior index or table page
	if header.Type.IsInterior() {
		buf := utils.ReadFile(file, startIndex+8, 4)
		header.RightMostPointer = utils.ReadUInt32(buf)
		startIndex += 4 // Increase the start index by 4 bytes to skip the right-most pointer
	}

	return &header, startIndex + 8, nil
}

type Page struct {
	Header  *PageHeader
	ITCells []cellUtils.InteriorTableCell
	LTCells []cellUtils.LeafTableCell
}

func ParsePage(file *os.File, startIndex uint, isHeader bool) (*Page, error) {
	pageHeaderIndex := startIndex
	if isHeader {
		pageHeaderIndex += 100
	}

	header, nextIndex, err := ParsePageHeader(file, pageHeaderIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse page header for the given page: %v", err)
	}

	cellOffsets := make([]uint16, header.NumberCells)
	for i := 0; i < int(header.NumberCells); i++ {
		cellOffsets[i] = utils.ReadUInt16(
			utils.ReadFile(file, nextIndex+(uint(i)*2), 2),
		)
	}

	page := Page{
		Header: header,
	}

	switch header.Type {
	case InteriorIndexPage:
		cells := make([]cellUtils.InteriorTableCell, header.NumberCells)
		for i, offset := range cellOffsets {
			cell := cellUtils.InteriorTableCell{}
			cell.Parse(file, startIndex+uint(offset))
			cells[i] = cell
		}
		page.ITCells = cells

	case LeafTablePage:
		cells := make([]cellUtils.LeafTableCell, header.NumberCells)
		for i, offset := range cellOffsets {
			cell := cellUtils.LeafTableCell{}
			cell.Parse(file, startIndex+uint(offset))
			cells[i] = cell
		}
		page.LTCells = cells

	default:
		return nil, fmt.Errorf("unsupported page type %d", header.Type)
	}

	return &page, nil
}
