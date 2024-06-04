package stmt

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql"
)

type SelectStatement struct {
	Fields []Field
	Table  string
}

func ParseSelectStatement(l *sql.Lexer) (*SelectStatement, error) {
	err := l.ExpectToken(sql.SELECT, "Parsing SELECT statement")
	if err != nil {
		return nil, err
	}

	fields, err := ParseFields(l)
	if err != nil {
		return nil, err
	}

	err = l.ExpectToken(sql.FROM, "Parsing SELECT statement")
	if err != nil {
		return nil, err
	}

	table, err := l.NextToken()
	if err != nil {
		return nil, err
	}
	if table.Type != sql.OBJ {
		return nil, fmt.Errorf("expected table name, got %s", table.Value)
	}

	return &SelectStatement{
		Fields: fields,
		Table:  table.Value,
	}, nil
}
