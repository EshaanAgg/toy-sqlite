package dot

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/defs"
)

func HandleDotCommand(commandData *defs.CommandData) {
	switch commandData.Command {
	case ".dbinfo":
		DBInfo(commandData)

	default:
		fmt.Println("Unknown command: ", commandData.Command)
	}
}
