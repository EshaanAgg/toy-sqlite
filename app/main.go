package main

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/cmd"
	"github/com/codecrafters-io/sqlite-starter-go/app/defs"
	"github/com/codecrafters-io/sqlite-starter-go/app/file"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <database-file-path> <command>")
		os.Exit(1)
	}

	databaseFilePath := os.Args[1]
	command := os.Args[2]

	// Open the database file
	dbFile, err := os.Open(databaseFilePath)
	if err != nil {
		fmt.Println("Failed to open database file: ", err)
		os.Exit(1)
	}

	// Parse the database header
	header, err := file.ParseDatabaseHeader(dbFile)
	if err != nil {
		fmt.Println("Failed to parse database header: ", err)
		os.Exit(1)
	}

	// Create the command data and execute the command
	commandData := defs.CommandData{
		DatabaseFile: dbFile,
		Command:      strings.Trim(command, " "),
		Header:       header,
	}

	cmd.HandleCommand(&commandData)
}
