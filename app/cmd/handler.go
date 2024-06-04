package cmd

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/cmd/dot"
	"github/com/codecrafters-io/sqlite-starter-go/app/cmd/sel"
	"github/com/codecrafters-io/sqlite-starter-go/app/defs"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql/stmt"
	"strings"
)

func HandleCommand(commandData *defs.CommandData) {
	if strings.HasPrefix(commandData.Command, ".") {
		dot.HandleDotCommand(commandData)
		return
	}

	cmdLexer := sql.NewLexer(commandData.Command)
	selectStmt, err := stmt.ParseSelectStatement(cmdLexer)
	if err != nil {
		fmt.Println("Error parsing command: ", err)
		fmt.Println(commandData.Command)
		fmt.Printf("%s^\n", strings.Repeat(" ", cmdLexer.CurPos()))
		return
	}

	sel.HandleSelectCommand(commandData, selectStmt)
}
