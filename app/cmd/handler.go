package cmd

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/cmd/dot"
	"github/com/codecrafters-io/sqlite-starter-go/app/defs"
	"strings"
)

func HandleCommand(commandData *defs.CommandData) {
	if strings.HasPrefix(commandData.Command, ".") {
		dot.HandleDotCommand(commandData)
		return
	}

	fmt.Println("Unknown command: ", commandData.Command)
}
