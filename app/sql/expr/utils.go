package expr

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func (e *Expr) ExpectType(expectedType string) error {
	if e.ValueType != expectedType {
		return fmt.Errorf("expected type %s, got %s", expectedType, e.ValueType)
	}
	return nil
}

func (e *Expr) ExpectTypeOneOf(expectedTypes ...string) error {
	for _, expectedType := range expectedTypes {
		if e.ValueType == expectedType {
			return nil
		}
	}
	return fmt.Errorf("expected one of %v, got %s", expectedTypes, e.ValueType)
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

func getFloatExpr(value float64) Expr {
	return Expr{
		ValueType: "FLOAT",
		value:     fmt.Sprintf("%f", value),
	}
}

func getBoolExpr(value bool) Expr {
	return Expr{
		ValueType: "BOOL",
		value:     fmt.Sprintf("%t", value),
	}
}

func getCallExpr(operator string, operands ...Expr) Expr {
	return Expr{
		ValueType: "CALL",
		operator:  operator,
		operands:  operands,
	}
}

func getBinaryExpr(operator string, operandA Expr, operandB Expr) Expr {
	return Expr{
		ValueType: "BINARY",
		operator:  operator,
		operands:  []Expr{operandA, operandB},
	}
}

func applyComparisonOp[T constraints.Ordered](a T, b T, op string) bool {
	switch op {
	case ">":
		return a > b
	case "<":
		return a < b
	case ">=":
		return a >= b
	case "<=":
		return a <= b
	case "=":
		return a == b
	case "!=":
		return a != b
	default:
		panic("unknown comparison operator")
	}
}

func applyArithmeticOp[T int | float64](a T, b T, op string) T {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unknown arithmetic operator")
	}
}
