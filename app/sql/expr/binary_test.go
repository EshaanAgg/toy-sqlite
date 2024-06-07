package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryEval(t *testing.T) {
	testCases := []testEval{
		{
			name: "Arithmetic Addition for Integer",
			inputExpr: getBinaryExpr(
				"+",
				getIntExpr(10),
				getIntExpr(20),
			),
			outputExpr: getIntExpr(30),
		},
		{
			name: "Arithmetic Addition for Float",
			inputExpr: getBinaryExpr(
				"+",
				getFloatExpr(10.5),
				getFloatExpr(20.5),
			),
			outputExpr: getFloatExpr(31.0),
		},
		{
			name: "Arithmetic Addition for Mixed",
			inputExpr: getBinaryExpr(
				"+",
				getIntExpr(10),
				getFloatExpr(20.5),
			),
			shouldErr: true,
		},
		// TODO: Add the same for -, *, /

		{
			name: "Logical AND",
			inputExpr: getBinaryExpr(
				"AND",
				getBoolExpr(true),
				getBoolExpr(false),
			),
			outputExpr: getBoolExpr(false),
		},
		{
			name: "Logical OR",
			inputExpr: getBinaryExpr(
				"OR",
				getBoolExpr(true),
				getBoolExpr(false),
			),
			outputExpr: getBoolExpr(true),
		},
		{
			name: "Logical Invalid #1",
			inputExpr: getBinaryExpr(
				"AND",
				getIntExpr(10),
				getIntExpr(20),
			),
			shouldErr: true,
		},
		{
			name: "Logical Invalid #2",
			inputExpr: getBinaryExpr(
				"OR",
				getStringExpr("hello"),
				getBoolExpr(true),
			),
			shouldErr: true,
		},
		{
			name: "Logical Invalid #3",
			inputExpr: getBinaryExpr(
				"AND",
				getBoolExpr(true),
				getFloatExpr(10.5),
			),
			shouldErr: true,
		},

		{
			name: "Comaprision > for Integer",
			inputExpr: getBinaryExpr(
				">",
				getIntExpr(10),
				getIntExpr(20),
			),
			outputExpr: getBoolExpr(false),
		},
		{
			name: "Comaprision > for Float",
			inputExpr: getBinaryExpr(
				">",
				getFloatExpr(10.5),
				getFloatExpr(20.5),
			),
			outputExpr: getBoolExpr(false),
		},
		{
			name: "Comaprision > for Mixed",
			inputExpr: getBinaryExpr(
				">",
				getIntExpr(10),
				getFloatExpr(20.5),
			),
			shouldErr: true,
		},
		// TODO: Add the same for <, >=, <=

		{
			name: "Comaprision = for Integer",
			inputExpr: getBinaryExpr(
				"=",
				getIntExpr(10),
				getIntExpr(20),
			),
			outputExpr: getBoolExpr(false),
		},
		{
			name: "Comaprision = for Float",
			inputExpr: getBinaryExpr(
				"=",
				getFloatExpr(10.5),
				getFloatExpr(20.5),
			),
			outputExpr: getBoolExpr(false),
		},
		{
			name: "Comaprision = for Bool",
			inputExpr: getBinaryExpr(
				"=",
				getBoolExpr(true),
				getBoolExpr(true),
			),
			outputExpr: getBoolExpr(true),
		},
		{
			name: "Comaprision = for Mixed",
			inputExpr: getBinaryExpr(
				"=",
				getIntExpr(10),
				getFloatExpr(10.0),
			),
			shouldErr: true,
		},

		{
			name: "Comaprision != for Integer",
			inputExpr: getBinaryExpr(
				"!=",
				getIntExpr(10),
				getIntExpr(20),
			),
			outputExpr: getBoolExpr(true),
		},
		{
			name: "Comaprision != for Float",
			inputExpr: getBinaryExpr(
				"!=",
				getFloatExpr(10.5),
				getFloatExpr(20.5),
			),
			outputExpr: getBoolExpr(true),
		},
		{
			name: "Comaprision != for Bool",
			inputExpr: getBinaryExpr(
				"!=",
				getBoolExpr(true),
				getBoolExpr(true),
			),
			outputExpr: getBoolExpr(false),
		},
		{
			name: "Comaprision != for Mixed",
			inputExpr: getBinaryExpr(
				"!=",
				getIntExpr(10),
				getFloatExpr(10.0),
			),
			shouldErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.inputExpr.evalBinaryOp()
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
