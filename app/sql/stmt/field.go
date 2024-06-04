package stmt

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql"
	"strings"
)

type Field struct {
	Name     string
	Type     string // "COLUMN", "COUNT", "ALL"
	Metadata string
}

func ParseFields(l *sql.Lexer) ([]Field, error) {
	fields := make([]Field, 0)

	tok, err := l.PeekToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get next token while parsing field: %v", err)
	}

	// *
	if tok.Type == sql.ASTERISK {
		l.NextToken()
		fields = append(fields, Field{
			Name: "ALL",
			Type: "ALL",
		})
		return fields, nil
	}

	// COUNT(*) or COUNT(column1, column2, ...)
	if tok.Type == sql.COUNT {
		l.NextToken()
		err = l.ExpectToken(sql.OPEN_PAREN, "Parsing COUNT fields")
		if err != nil {
			return nil, err
		}

		tok, err = l.PeekToken()
		if err != nil {
			return nil, fmt.Errorf("failed to get next token while parsing COUNT fields: %v", err)
		}
		if tok.Type == sql.ASTERISK {
			l.NextToken()
			fields = append(fields,
				Field{
					Name:     "ALL",
					Type:     "COUNT",
					Metadata: "*",
				},
			)
		} else {
			columns, err := parseColumnNames(l)
			if err != nil {
				return nil, fmt.Errorf("failed to parse columns: %v", err)
			}
			fields = append(fields,
				Field{
					Name:     "COLS",
					Type:     "COUNT",
					Metadata: strings.Join(columns, ","),
				})
		}

		err = l.ExpectToken(sql.CLOSE_PAREN, "Parsing COUNT fields")
		if err != nil {
			return nil, err
		}

		return fields, nil
	}

	// column1, column2, ...
	columns, err := parseColumnNames(l)
	if err != nil {
		return nil, fmt.Errorf("failed to parse columns: %v", err)
	}
	for _, column := range columns {
		fields = append(fields, Field{
			Name: column,
			Type: "COLUMN",
		})
	}

	return fields, nil
}

// Parses comma sperated columns names
func parseColumnNames(l *sql.Lexer) ([]string, error) {
	columns := make([]string, 0)

	tok, err := l.PeekToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get next token while parsing columns: %v", err)
	}

	for tok.Type == sql.OBJ {
		columns = append(columns, tok.Value)
		l.NextToken()

		tok, err = l.PeekToken()
		if err != nil {
			return nil, fmt.Errorf("failed to get next token while parsing columns: %v", err)
		}

		if tok.Type == sql.COMMA {
			l.NextToken()
			tok, err = l.PeekToken()
			if err != nil {
				return nil, fmt.Errorf("failed to get next token while parsing columns: %v", err)
			}
		}
	}

	return columns, nil
}
