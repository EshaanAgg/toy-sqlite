package sql_test

import (
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getAllTokens(input string) ([]sql.Token, error) {
	lexer := sql.NewLexer(input)
	var tokens []sql.Token
	for {
		token, err := lexer.NextToken()
		if err != nil {
			return nil, fmt.Errorf("failed to get next token: %v", err)
		}

		tokens = append(tokens, token)
		if token.Type == sql.EOF {
			break
		}
	}

	return tokens, nil
}

type testCaseLexer struct {
	desc  string
	input string
	want  []sql.Token
}

func TestLexer(t *testing.T) {
	testCases := []testCaseLexer{
		{
			desc:  "Simple SELECT statement",
			input: "SELECT * FROM table",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "SELECT statement with WHERE clause",
			input: "SELECT * FROM table WHERE column = 10",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table"},
				{Type: sql.WHERE, Value: "WHERE"},
				{Type: sql.OBJ, Value: "column"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.INT, Value: "10"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "SELECT statement with WHERE clause and LOGICAL operators",
			input: "SELECT * FROM table WHERE column = 10 AND column2 = 20 OR column3 = 30",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table"},
				{Type: sql.WHERE, Value: "WHERE"},
				{Type: sql.OBJ, Value: "column"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.INT, Value: "10"},
				{Type: sql.AND, Value: "AND"},
				{Type: sql.OBJ, Value: "column2"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.INT, Value: "20"},
				{Type: sql.OR, Value: "OR"},
				{Type: sql.OBJ, Value: "column3"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.INT, Value: "30"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "SELECT statement with WHERE clause, all ARITHMETIC operators, PARENTHESIS and FLOATS",
			input: "SELECT * FROM table WHERE (column = 10 AND column2 != 20) OR (column3 > 30.0 AND column4 >= 40.124) OR (column5 < 50.10 AND column6 <= 60)",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table"},
				{Type: sql.WHERE, Value: "WHERE"},
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.OBJ, Value: "column"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.INT, Value: "10"},
				{Type: sql.AND, Value: "AND"},
				{Type: sql.OBJ, Value: "column2"},
				{Type: sql.NEQ, Value: "!="},
				{Type: sql.INT, Value: "20"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.OR, Value: "OR"},
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.OBJ, Value: "column3"},
				{Type: sql.GT, Value: ">"},
				{Type: sql.FLOAT, Value: "30.0"},
				{Type: sql.AND, Value: "AND"},
				{Type: sql.OBJ, Value: "column4"},
				{Type: sql.GTE, Value: ">="},
				{Type: sql.FLOAT, Value: "40.124"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.OR, Value: "OR"},
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.OBJ, Value: "column5"},
				{Type: sql.LT, Value: "<"},
				{Type: sql.FLOAT, Value: "50.10"},
				{Type: sql.AND, Value: "AND"},
				{Type: sql.OBJ, Value: "column6"},
				{Type: sql.LTE, Value: "<="},
				{Type: sql.INT, Value: "60"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "SELECT statement with WHERE clause with STRING values",
			input: `SELECT * FROM table WHERE column = "value"`,
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table"},
				{Type: sql.WHERE, Value: "WHERE"},
				{Type: sql.OBJ, Value: "column"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.STRING, Value: "value"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "COUNT statement",
			input: "SELECT COUNT(*) FROM table",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.COUNT, Value: "COUNT"},
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "COUNT statement with WHERE clause",
			input: "SELECT COUNT(column1, column2) FROM table WHERE column = 10",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.COUNT, Value: "COUNT"},
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.OBJ, Value: "column1"},
				{Type: sql.COMMA, Value: ","},
				{Type: sql.OBJ, Value: "column2"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table"},
				{Type: sql.WHERE, Value: "WHERE"},
				{Type: sql.OBJ, Value: "column"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.INT, Value: "10"},
				{Type: sql.EOF, Value: ""},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := getAllTokens(tc.input)
			assert.NoError(t, err)

			assert.Equal(t, len(tc.want), len(got))
			for i := range tc.want {
				assert.Equal(t, tc.want[i].Type, got[i].Type)

				// For float values, convert to float64 and compare
				if tc.want[i].Type == sql.FLOAT {
					gotFloat, err := strconv.ParseFloat(got[i].Value, 64)
					assert.NoError(t, err)
					wantFloat, err := strconv.ParseFloat(tc.want[i].Value, 64)
					assert.NoError(t, err)

					assert.InDelta(t, wantFloat, gotFloat, 0.0001)
				} else {
					assert.Equal(t, tc.want[i].Value, got[i].Value)
				}
			}
		})
	}
}
