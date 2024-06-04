package tables

import (
	"fmt"
	"strings"
)

func DebugRecordData(data [][]Record) string {
	records := make([]string, 0)
	for _, row := range data {
		items := make([]string, 0)
		for _, item := range row {
			items = append(items, item.Debug())
		}
		records = append(records, fmt.Sprintf("[%s]", strings.Join(items, ", ")))
	}

	return strings.Join(records, " ")
}
