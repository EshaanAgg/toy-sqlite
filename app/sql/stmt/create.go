package stmt

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql"
)

type Column struct {
	Name          string
	Type          string
	PrimaryKey    bool
	AutoIncrement bool
}

type CreateTableStatement struct {
	TableName string
	Columns   []Column
}

func ParseCreateTableStatement(l *sql.Lexer) (*CreateTableStatement, error) {
	err := l.ExpectToken(sql.CREATE, "Parsing CREATE TABLE statement")
	if err != nil {
		return nil, err
	}

	err = l.ExpectToken(sql.TABLE, "Parsing CREATE TABLE statement")
	if err != nil {
		return nil, err
	}

	tableName, err := l.NextToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get table name: %v", err)
	}
	if tableName.Type != sql.OBJ {
		return nil, fmt.Errorf("expected table name, got %s", tableName.Value)
	}

	err = l.ExpectToken(sql.OPEN_PAREN, "Parsing CREATE TABLE statement")
	if err != nil {
		return nil, err
	}

	columns, err := parseColumns(l)
	if err != nil {
		return nil, err
	}

	err = l.ExpectToken(sql.CLOSE_PAREN, "Parsing CREATE TABLE statement")
	if err != nil {
		return nil, err
	}

	return &CreateTableStatement{
		TableName: tableName.Value,
		Columns:   columns,
	}, nil
}

// Parses a single column definition in a CREATE TABLE statement
func parseSingleColumn(l *sql.Lexer) (Column, error) {
	col := Column{}

	// Column name
	columnName, err := l.NextToken()
	if err != nil {
		return col, fmt.Errorf("failed to get column name: %v", err)
	}
	if columnName.Type != sql.OBJ {
		return col, fmt.Errorf("expected column name, got %s", columnName.Value)
	}
	col.Name = columnName.Value

	// Data type
	columnType, err := l.NextToken()
	if err != nil {
		return col, fmt.Errorf("failed to get column type: %v", err)
	}
	if columnType.Type != sql.INTEGER && columnType.Type != sql.TEXT {
		return col, fmt.Errorf("expected column type, got %s", columnType.Value)
	}
	col.Type = columnType.Value

	// Check for PRIMARY KEY or AUTOINCREMENT
	var primaryKey, autoIncrement bool
	for {
		token, err := l.PeekToken()
		if err != nil {
			return col, fmt.Errorf("failed to get next token: %v", err)
		}

		if token.Type == sql.PRIMARY {
			l.NextToken()
			err = l.ExpectToken(sql.KEY, "Parsing CREATE TABLE statement")
			if err != nil {
				return col, fmt.Errorf("expected KEY after primary: %v", err)
			}
			primaryKey = true
		} else if token.Type == sql.AUTOINCREMENT {
			l.NextToken()
			autoIncrement = true
		} else {
			break
		}
	}

	return Column{
		Name:          columnName.Value,
		Type:          columnType.Value,
		PrimaryKey:    primaryKey,
		AutoIncrement: autoIncrement,
	}, nil
}

// Parses comma separated columns definitions in a CREATE TABLE statement
func parseColumns(l *sql.Lexer) ([]Column, error) {
	columns := make([]Column, 0)

	for {
		col, err := parseSingleColumn(l)
		if err != nil {
			return nil, fmt.Errorf("failed to parse column: %v", err)
		}

		columns = append(columns, col)

		// Check if there are more columns
		token, err := l.PeekToken()
		if err != nil {
			return nil, fmt.Errorf("failed to get next token: %v", err)
		}

		if token.Type == sql.COMMA {
			l.NextToken()
			continue
		}

		break
	}

	return columns, nil
}
