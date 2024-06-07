package expr

import "fmt"

// Evaluates the expression provided so that it's type
// can be narrowed to one of the BASIC types (INT, STRING, BOOL, FLOAT)
func (e *Expr) Eval() error {
	switch e.ValueType {
	case "INT", "STRING", "BOOL", "FLOAT":
		return nil

	case "BINARY":
		// Evaluate the operands first

	case "CALL":
		err := e.evalCall()
		if err != nil {
			return fmt.Errorf("failed to evaluate call: %v", err)
		}
		return nil

	default:
		return fmt.Errorf("unexpected expression type: %s", e.ValueType)
	}

	return nil
}
