package expr

import "fmt"

func (e *Expr) evalUnaryOp() error {
	// Evaluate the operand first
	err := e.evalOperands()
	if err != nil {
		return fmt.Errorf("failed to evaluate operands: %v", err)
	}

	// Check that the number of operands is exactly 1
	if len(e.operands) != 1 {
		return fmt.Errorf("expected 1 operand, got %d", len(e.operands))
	}
	op := e.operands[0]

	var resExpr Expr
	switch e.operator {
	case "NOT":
		resExpr, err = logicalNot(op)
	case "-":
		resExpr, err = arithmeticNeg(op)
	default:
		return fmt.Errorf("unknown operator: %s", e.operator)
	}

	if err != nil {
		return fmt.Errorf("failed to evaluate unary operation: %v", err)
	}

	// Update the expression with the result
	*e = resExpr
	return nil
}

func logicalNot(op Expr) (Expr, error) {
	resExpr := Expr{}

	if err := op.ExpectType("BOOL"); err != nil {
		return resExpr, fmt.Errorf("operand is not supported for NOT operation: %v", err)
	}

	resExpr.ValueType = "BOOL"
	resExpr.value = fmt.Sprintf("%t", !op.BoolVal())
	return resExpr, nil
}

func arithmeticNeg(op Expr) (Expr, error) {
	resExpr := Expr{}

	if err := op.ExpectTypeOneOf("INTEGER", "FLOAT"); err != nil {
		return resExpr, fmt.Errorf("operand is not supported for NEG operation: %v", err)
	}

	resExpr.ValueType = op.ValueType
	switch op.ValueType {
	case "INTEGER":
		resExpr.value = fmt.Sprintf("%d", -op.IntVal())
	case "FLOAT":
		resExpr.value = fmt.Sprintf("%f", -op.FloatVal())
	default:
		return resExpr, fmt.Errorf("unexpected operand type: %s", op.ValueType)
	}
	return resExpr, nil
}
