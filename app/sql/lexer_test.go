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
			input: "SELECT * FROM table_name",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table_name"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "SELECT statement with WHERE clause",
			input: "SELECT * FROM table_name WHERE column = 10",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table_name"},
				{Type: sql.WHERE, Value: "WHERE"},
				{Type: sql.OBJ, Value: "column"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.INT, Value: "10"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "SELECT statement with WHERE clause and LOGICAL operators",
			input: "SELECT * FROM table_name WHERE column = 10 AND column2 = 20 OR column3 = 30",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table_name"},
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
			desc: "SELECT statement with WHERE clause, all ARITHMETIC operators, PARENTHESIS and FLOATS",
			input: `
				SELECT * FROM table_name 
				WHERE 
					(column = 10 AND column2 != 20) 
					OR (column3 > 30.0 AND column4 >= 40.124) 
					OR (column5 < 50.10 AND column6 <= 60)
				`,
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table_name"},
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
			input: "SELECT * FROM table_name WHERE column = 'value'",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table_name"},
				{Type: sql.WHERE, Value: "WHERE"},
				{Type: sql.OBJ, Value: "column"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.STRING, Value: "value"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "COUNT statement",
			input: "SELECT COUNT(*) FROM table_name",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.COUNT, Value: "COUNT"},
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table_name"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "COUNT statement with WHERE clause",
			input: "SELECT COUNT(column1, column2) FROM table_name WHERE column = 10",
			want: []sql.Token{
				{Type: sql.SELECT, Value: "SELECT"},
				{Type: sql.COUNT, Value: "COUNT"},
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.OBJ, Value: "column1"},
				{Type: sql.COMMA, Value: ","},
				{Type: sql.OBJ, Value: "column2"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.FROM, Value: "FROM"},
				{Type: sql.OBJ, Value: "table_name"},
				{Type: sql.WHERE, Value: "WHERE"},
				{Type: sql.OBJ, Value: "column"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.INT, Value: "10"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "CREATE TABLE statement with MULTIPLE column and PRIMARY KEY",
			input: "CREATE TABLE table_name (column1 INTEGER PRIMARY KEY AUTOINCREMENT, column2 TEXT)",
			want: []sql.Token{
				{Type: sql.CREATE, Value: "CREATE"},
				{Type: sql.TABLE, Value: "TABLE"},
				{Type: sql.OBJ, Value: "table_name"},
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.OBJ, Value: "column1"},
				{Type: sql.INTEGER, Value: "INTEGER"},
				{Type: sql.PRIMARY, Value: "PRIMARY"},
				{Type: sql.KEY, Value: "KEY"},
				{Type: sql.AUTOINCREMENT, Value: "AUTOINCREMENT"},
				{Type: sql.COMMA, Value: ","},
				{Type: sql.OBJ, Value: "column2"},
				{Type: sql.TEXT, Value: "TEXT"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "CREATE TABLE statement with SINGLE column and NO PRIMARY KEY",
			input: "CREATE TABLE table_name (column1 INTEGER)",
			want: []sql.Token{
				{Type: sql.CREATE, Value: "CREATE"},
				{Type: sql.TABLE, Value: "TABLE"},
				{Type: sql.OBJ, Value: "table_name"},
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.OBJ, Value: "column1"},
				{Type: sql.INTEGER, Value: "INTEGER"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "Logical Operators",
			input: "AND OR NOT",
			want: []sql.Token{
				{Type: sql.AND, Value: "AND"},
				{Type: sql.OR, Value: "OR"},
				{Type: sql.NOT, Value: "NOT"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "Arithmetic Operators",
			input: "+ - / % *",
			want: []sql.Token{
				{Type: sql.ADD, Value: "+"},
				{Type: sql.SUB, Value: "-"},
				{Type: sql.DIV, Value: "/"},
				{Type: sql.MOD, Value: "%"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "Arithmetic Operators without space",
			input: "+-/*%",
			want: []sql.Token{
				{Type: sql.ADD, Value: "+"},
				{Type: sql.SUB, Value: "-"},
				{Type: sql.DIV, Value: "/"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.MOD, Value: "%"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "Sample Expression",
			input: "(1+2)*(3-4)/5%6 AND UPPER(COLUMN1=10 OR COLUMN2=20)",
			want: []sql.Token{
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.INT, Value: "1"},
				{Type: sql.ADD, Value: "+"},
				{Type: sql.INT, Value: "2"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.ASTERISK, Value: "*"},
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.INT, Value: "3"},
				{Type: sql.SUB, Value: "-"},
				{Type: sql.INT, Value: "4"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.DIV, Value: "/"},
				{Type: sql.INT, Value: "5"},
				{Type: sql.MOD, Value: "%"},
				{Type: sql.INT, Value: "6"},
				{Type: sql.AND, Value: "AND"},
				{Type: sql.FUNC, Value: "UPPER"},
				{Type: sql.OPEN_PAREN, Value: "("},
				{Type: sql.OBJ, Value: "COLUMN1"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.INT, Value: "10"},
				{Type: sql.OR, Value: "OR"},
				{Type: sql.OBJ, Value: "COLUMN2"},
				{Type: sql.EQ, Value: "="},
				{Type: sql.INT, Value: "20"},
				{Type: sql.CLOSE_PAREN, Value: ")"},
				{Type: sql.EOF, Value: ""},
			},
		},

		{
			desc:  "Function names",
			input: "UPPER LOWER CONCAT",
			want: []sql.Token{
				{Type: sql.FUNC, Value: "UPPER"},
				{Type: sql.FUNC, Value: "LOWER"},
				{Type: sql.FUNC, Value: "CONCAT"},
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
