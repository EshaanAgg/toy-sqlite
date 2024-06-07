package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testEval struct {
	name       string
	shouldErr  bool
	inputExpr  Expr
	outputExpr Expr
}

func TestCallEval(t *testing.T) {
	testCases := []testEval{
		{
			name: "UPPER - Valid Call",
			inputExpr: getCallExpr(
				"UPPER",
				getStringExpr("hello"),
			),
			outputExpr: getStringExpr("HELLO"),
			shouldErr:  false,
		},
		{
			name: "UPPER - Multiple Arguments",
			inputExpr: getCallExpr(
				"UPPER",
				getStringExpr("hello"),
				getStringExpr("world"),
			),
			shouldErr: true,
		},
		{
			name: "UPPER - Invalid Argument Type",
			inputExpr: getCallExpr(
				"UPPER",
				getIntExpr(10),
			),
			shouldErr: true,
		},
		{
			name: "LOWER - Valid Call",
			inputExpr: getCallExpr(
				"LOWER",
				getStringExpr("HELLO"),
			),
			outputExpr: getStringExpr("hello"),
			shouldErr:  false,
		},
		{
			name: "LOWER - Multiple Arguments",
			inputExpr: getCallExpr(
				"LOWER",
				getStringExpr("hello"),
				getStringExpr("world"),
			),
			shouldErr: true,
		},
		{
			name: "LOWER - Invalid Argument Type",
			inputExpr: getCallExpr(
				"LOWER",
				getIntExpr(10),
			),
			shouldErr: true,
		},
		{
			name: "CONCAT - Valid Call",
			inputExpr: getCallExpr(
				"CONCAT",
				getStringExpr("hello"),
				getStringExpr("world"),
				getStringExpr("!!!"),
			),
			outputExpr: getStringExpr("helloworld!!!"),
			shouldErr:  false,
		},
		{
			name:      "CONCAT - No Arguments",
			inputExpr: getCallExpr("CONCAT"),
			shouldErr: true,
		},
		{
			name: "CONCAT - Invalid Argument Type",
			inputExpr: getCallExpr(
				"CONCAT",
				getIntExpr(10),
			),
			shouldErr: true,
		},
		{
			name: "Nesting #1 => UPPER(CONCAT('hello', 'world'))",
			inputExpr: getCallExpr(
				"UPPER",
				getCallExpr(
					"CONCAT",
					getStringExpr("hello"),
					getStringExpr("world"),
				),
			),
			outputExpr: getStringExpr("HELLOWORLD"),
			shouldErr:  false,
		},
		{
			name: "Nesting #2 => CONCAT(UPPER('hello'), LOWER('WORLD'))",
			inputExpr: getCallExpr(
				"CONCAT",
				getCallExpr(
					"UPPER",
					getStringExpr("hello"),
				),
				getCallExpr(
					"LOWER",
					getStringExpr("WORLD"),
				),
			),
			outputExpr: getStringExpr("HELLOworld"),
			shouldErr:  false,
		},
		{
			name: "Nesting #3 => UPPER(LOWER(CONCAT('hello', 'world')))",
			inputExpr: getCallExpr(
				"UPPER",
				getCallExpr(
					"LOWER",
					getCallExpr(
						"CONCAT",
						getStringExpr("hello"),
						getStringExpr("world"),
					),
				),
			),
			outputExpr: getStringExpr("HELLOWORLD"),
			shouldErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.inputExpr.evalCall()
			if tc.shouldErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.outputExpr.value, tc.inputExpr.value)
			assert.Equal(t, tc.outputExpr.ValueType, tc.inputExpr.ValueType)
			assert.Equal(t, 0, len(tc.inputExpr.operands))
			assert.Equal(t, "", tc.inputExpr.operator)
		})
	}
}
