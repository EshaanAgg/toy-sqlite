package stmt_test

import (
	"github/com/codecrafters-io/sqlite-starter-go/app/sql"
	"github/com/codecrafters-io/sqlite-starter-go/app/sql/stmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseCreateTable struct {
	desc  string
	input string
	want  *stmt.CreateTableStatement
}

func TestCreateTableStatement(t *testing.T) {
	testCases := []testCaseCreateTable{
		{
			desc: "Simple CREATE TABLE statement",
			input: `CREATE TABLE table_name (
                column1 INTEGER,
                column2 TEXT
            )`,
			want: &stmt.CreateTableStatement{
				TableName: "table_name",
				Columns: []stmt.Column{
					{Name: "column1", Type: "INTEGER"},
					{Name: "column2", Type: "TEXT"},
				},
			},
		},

		{
			desc: "CREATE TABLE statement with primary key",
			input: `CREATE TABLE table_name (
                column1 INTEGER PRIMARY KEY,
                column2 TEXT
            )`,
			want: &stmt.CreateTableStatement{
				TableName: "table_name",
				Columns: []stmt.Column{
					{Name: "column1", Type: "INTEGER", PrimaryKey: true},
					{Name: "column2", Type: "TEXT"},
				},
			},
		},

		{
			desc: "CREATE TABLE statement with autoincrement",
			input: `CREATE TABLE table_name (
                column1 INTEGER PRIMARY KEY AUTOINCREMENT,
                column2 TEXT
            )`,
			want: &stmt.CreateTableStatement{
				TableName: "table_name",
				Columns: []stmt.Column{
					{Name: "column1", Type: "INTEGER", PrimaryKey: true, AutoIncrement: true},
					{Name: "column2", Type: "TEXT"},
				},
			},
		},

		{
			desc: "CREATE TABLE statement with AUTOINCREMENT followed by PRIMARY KEY",
			input: `CREATE TABLE table_name (
                column1 INTEGER AUTOINCREMENT PRIMARY KEY,
                column2 TEXT
            )`,
			want: &stmt.CreateTableStatement{
				TableName: "table_name",
				Columns: []stmt.Column{
					{Name: "column1", Type: "INTEGER", PrimaryKey: true, AutoIncrement: true},
					{Name: "column2", Type: "TEXT"},
				},
			},
		},

		{
			desc: "CREATE TABLE statement with MULTIPLE columns and MULTIPLE AUTOINCREMENTING keys",
			input: `CREATE TABLE table_name (
                column1 INTEGER PRIMARY KEY AUTOINCREMENT,
                column2 TEXT,
                column3 INTEGER AUTOINCREMENT
            )`,
			want: &stmt.CreateTableStatement{
				TableName: "table_name",
				Columns: []stmt.Column{
					{Name: "column1", Type: "INTEGER", PrimaryKey: true, AutoIncrement: true},
					{Name: "column2", Type: "TEXT"},
					{Name: "column3", Type: "INTEGER", AutoIncrement: true},
				},
			},
		},

		{
			desc: "CREATE TABLE statement with SINGLE column",
			input: `CREATE TABLE table_name (
                column1 INTEGER
            )`,
			want: &stmt.CreateTableStatement{
				TableName: "table_name",
				Columns: []stmt.Column{
					{Name: "column1", Type: "INTEGER"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			lexer := sql.NewLexer(tc.input)

			got, err := stmt.ParseCreateTableStatement(lexer)
			assert.NoError(t, err)

			assert.Equal(t, tc.want.TableName, got.TableName)
			assert.Equal(t, len(tc.want.Columns), len(got.Columns))
			for i := range tc.want.Columns {
				assert.Equal(t, tc.want.Columns[i].Name, got.Columns[i].Name)
				assert.Equal(t, tc.want.Columns[i].Type, got.Columns[i].Type)
				assert.Equal(t, tc.want.Columns[i].PrimaryKey, got.Columns[i].PrimaryKey)
				assert.Equal(t, tc.want.Columns[i].AutoIncrement, got.Columns[i].AutoIncrement)
			}
		})
	}
}
