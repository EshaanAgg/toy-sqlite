package sel

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/defs"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql/stmt"
)

func HandleSelectCommand(cmdData *defs.CommandData, selectStmt *stmt.SelectStatement) {
	tableRootPage, err := getTableRootPage(selectStmt.Table, cmdData.DatabaseFile)
	if err != nil {
		fmt.Println("Error getting table root page: ", err)
		return
	}
}
