package expr

import "fmt"

func (e *Expr) ExpectType(expectedType string) error {
	if e.ValueType != expectedType {
		return fmt.Errorf("expected type %s, got %s", expectedType, e.ValueType)
	}
	return nil
}

func (e *Expr) evalOperands() error {
	for i := 0; i < len(e.operands); i++ {
		err := e.operands[i].Eval()
		if err != nil {
			return fmt.Errorf("failed to evaluate operand %d: %v", i, err)
		}
	}
	return nil
}

func getStringExpr(value string) Expr {
	return Expr{
		ValueType: "STRING",
		value:     value,
	}
}

func getIntExpr(value int) Expr {
	return Expr{
		ValueType: "INTEGER",
		value:     fmt.Sprintf("%d", value),
	}
}

func getCallExpr(functionName string, operands ...Expr) Expr {
	return Expr{
		ValueType:    "CALL",
		functionName: functionName,
		operands:     operands,
	}
}
