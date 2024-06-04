package stmt_test

import (
	"github/com/codecrafters-io/sqlite-starter-go/app/sql"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql/stmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseSelect struct {
	desc  string
	input string
	want  *stmt.SelectStatement
}

func TestSelectStatement(t *testing.T) {
	testCases := []testCaseSelect{
		{
			desc:  "Simple SELECT statement",
			input: "SELECT * FROM table",
			want: &stmt.SelectStatement{
				Fields: []stmt.Field{
					{Name: "ALL", Type: "ALL"},
				},
				Table: "table",
			},
		},

		{
			desc:  "SELECT statement with MULTIPLE columns",
			input: "SELECT column1, column2, column3 FROM table",
			want: &stmt.SelectStatement{
				Fields: []stmt.Field{
					{Name: "column1", Type: "COLUMN"},
					{Name: "column2", Type: "COLUMN"},
					{Name: "column3", Type: "COLUMN"},
				},
				Table: "table",
			},
		},

		{
			desc:  "SELECT statement with COUNT(*)",
			input: "SELECT COUNT(*) FROM table",
			want: &stmt.SelectStatement{
				Fields: []stmt.Field{
					{Name: "COUNT", Type: "COUNT", Metadata: "*"},
				},
				Table: "table",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			l := sql.NewLexer(tc.input)
			got, err := stmt.ParseSelectStatement(l)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			assert.Equal(t, tc.want, got)
		})
	}
}
