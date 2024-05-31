package defs

import (
	"github/com/codecrafters-io/sqlite-starter-go/app/file"
	"os"
)

// Main struct that stores all the data related to the command being executed
type CommandData struct {
	DatabaseFile *os.File
	ReadIndex    int
	Command      string
	Header       *file.DatabaseHeader
}
