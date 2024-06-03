package dot

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/defs"
	"github/com/codecrafters-io/sqlite-starter-go/app/file/tables"
	"strings"
)

func Tables(commandData *defs.CommandData) {
	tableNames := tables.GetSQLiteSchema(commandData.DatabaseFile).GetTableNames()

	fmt.Println(strings.Join(tableNames, "\t"))
}
