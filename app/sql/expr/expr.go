package expr

import (
	"fmt"
)

type Expr struct {
	value        string
	ValueType    string // "STRING", "INT", "BOOL", "FLOAT", "BINARY", "CALL"
	functionName string // only used when ValueType is "CALL"
	operands     []Expr
}

// Debug print for Expr
// String() method is automatically called when fmt.Print is called on an Expr object
func (e *Expr) String() string {
	return fmt.Sprintf(`Expr{value: %s, ValueType: %s, operands: %v}`, e.value, e.ValueType, e.operands)
}
