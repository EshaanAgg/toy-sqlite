package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnaryEval(t *testing.T) {
	testCases := []testEval{
		// NOT
		{
			name: "Logical NOT for Bool",
			inputExpr: getUnaryExpr(
				"NOT",
				getBoolExpr(true),
			),
			outputExpr: getBoolExpr(false),
		},
		{
			name: "Logical NOT for Bool",
			inputExpr: getUnaryExpr(
				"NOT",
				getBoolExpr(false),
			),
			outputExpr: getBoolExpr(true),
		},
		{
			name: "Logical NOT for Integer",
			inputExpr: getUnaryExpr(
				"NOT",
				getIntExpr(10),
			),
			shouldErr: true,
		},
		{
			name: "Logical NOT for Float",
			inputExpr: getUnaryExpr(
				"NOT",
				getFloatExpr(10.5),
			),
			shouldErr: true,
		},
		{
			name: "Logical NOT for String",
			inputExpr: getUnaryExpr(
				"NOT",
				getStringExpr("hello"),
			),
			shouldErr: true,
		},

		// NEG
		{
			name: "Arithmetic NEG for Integer",
			inputExpr: getUnaryExpr(
				"-",
				getIntExpr(10),
			),
			outputExpr: getIntExpr(-10),
		},
		{
			name: "Arithmetic NEG for Float",
			inputExpr: getUnaryExpr(
				"-",
				getFloatExpr(10.5),
			),
			outputExpr: getFloatExpr(-10.5),
		},
		{
			name: "Arithmetic NEG for Bool",
			inputExpr: getUnaryExpr(
				"-",
				getBoolExpr(true),
			),
			shouldErr: true,
		},
		{
			name: "Arithmetic NEG for String",
			inputExpr: getUnaryExpr(
				"-",
				getStringExpr("hello"),
			),
			shouldErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.inputExpr.Eval()
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
