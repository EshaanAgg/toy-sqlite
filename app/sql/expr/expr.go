package expr

import (
	"fmt"
)

type Expr struct {
	value     string
	ValueType string // STRING, INTEGER, BOOL, FLOAT, CALL, UNARY, BINARY
	operator  string // Only used when ValueType is not a basic type (STRING, INTEGER, BOOL, FLOAT)
	operands  []Expr
}

// Debug print for Expr
// String() method is automatically called when fmt.Print is called on an Expr object
func (e *Expr) String() string {
	return fmt.Sprintf(`Expr{value: %s, ValueType: %s, operands: %v}`, e.value, e.ValueType, e.operands)
}
