package stmt_test

import (
	"github/com/codecrafters-io/sqlite-starter-go/app/sql"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql/stmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseField struct {
	name  string
	query string
	want  []stmt.Field
}

func TestField(t *testing.T) {
	testCases := []testCaseField{
		{
			name:  "All fields",
			query: "*",
			want: []stmt.Field{
				{Name: "ALL", Type: "ALL"},
			},
		},

		{
			name:  "Single column",
			query: "column1",
			want: []stmt.Field{
				{Name: "column1", Type: "COLUMN"},
			},
		},

		{
			name:  "Multiple columns",
			query: "column1, column2, column3",
			want: []stmt.Field{
				{Name: "column1", Type: "COLUMN"},
				{Name: "column2", Type: "COLUMN"},
				{Name: "column3", Type: "COLUMN"},
			},
		},

		{
			name:  "COUNT for records",
			query: "COUNT(*)",
			want: []stmt.Field{
				{Name: "COUNT", Type: "COUNT", Metadata: "*"},
			},
		},

		{
			name:  "COUNT for columns",
			query: "COUNT(column1, column2)",
			want: []stmt.Field{
				{Name: "COUNT", Type: "COUNT", Metadata: "column1,column2"},
			},
		},

		{
			name:  "COUNT for one column",
			query: "COUNT(column1)",
			want: []stmt.Field{
				{Name: "COUNT", Type: "COUNT", Metadata: "column1"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := sql.NewLexer(tc.query)
			got, err := stmt.ParseFields(l)

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)

			nextToken, err := l.NextToken()
			assert.NoError(t, err)
			assert.Equal(t, sql.EOF, nextToken.Type)
		})
	}
}
