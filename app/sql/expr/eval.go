package expr

import "fmt"

// Evaluates the expression provided so that it's type
// can be narrowed to one of the BASIC types (INTEGER, STRING, BOOL, FLOAT)
func (e *Expr) Eval() error {
	switch e.ValueType {
	case "INTEGER", "STRING", "BOOL", "FLOAT":
		return nil

	case "UNARY":
		err := e.evalUnaryOp()
		if err != nil {
			return fmt.Errorf("failed to evaluate unary operation: %v", err)
		}
		return nil

	case "BINARY":
		err := e.evalBinaryOp()
		if err != nil {
			return fmt.Errorf("failed to evaluate binary operation: %v", err)
		}
		return nil

	case "CALL":
		err := e.evalCall()
		if err != nil {
			return fmt.Errorf("failed to evaluate call: %v", err)
		}
		return nil

	default:
		return fmt.Errorf("unexpected expression type: %s", e.ValueType)
	}
}
