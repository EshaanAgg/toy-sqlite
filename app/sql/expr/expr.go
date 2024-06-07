package expr

import (
	"fmt"
)

type Expr struct {
	value     string
	ValueType string // "STRING", "INTEGER", "BOOL", "FLOAT", "BINARY", "CALL"
	operator  string // only used when ValueType is "CALL" or "BINARY"
	operands  []Expr
}

// Debug print for Expr
// String() method is automatically called when fmt.Print is called on an Expr object
func (e *Expr) String() string {
	return fmt.Sprintf(`Expr{value: %s, ValueType: %s, operands: %v}`, e.value, e.ValueType, e.operands)
}
