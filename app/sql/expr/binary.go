package expr

import "fmt"

func (e *Expr) evalBinaryOp() error {
	// Evaluate the operands first
	err := e.evalOperands()
	if err != nil {
		return fmt.Errorf("failed to evaluate operands: %v", err)
	}

	// Check that the number of operands is exactly 2
	if len(e.operands) != 2 {
		return fmt.Errorf("expected 2 operands, got %d", len(e.operands))
	}
	opA := e.operands[0]
	opB := e.operands[1]

	var resExpr Expr
	switch e.operator {
	case "+", "-", "*", "/":
		resExpr, err = arithmeticOp(opA, opB, e.operator)
	case "AND", "OR":
		resExpr, err = logicalOp(opA, opB, e.operator)
	case ">=", "<=", ">", "<", "=", "!=":
		resExpr, err = comparisionOp(opA, opB, e.operator)
	default:
		return fmt.Errorf("unknown operator: %s", e.operator)
	}

	if err != nil {
		return fmt.Errorf("failed to evaluate binary operation: %v", err)
	}

	// Update the expression with the result
	*e = resExpr
	return nil
}

func arithmeticOp(opA Expr, opB Expr, op string) (Expr, error) {
	resExpr := Expr{}

	if err := opA.ExpectTypeOneOf("INTEGER", "FLOAT"); err != nil {
		return resExpr, fmt.Errorf("operand A is not supported for Arithmetic operations: %v", err)
	}
	if err := opB.ExpectTypeOneOf("INTEGER", "FLOAT"); err != nil {
		return resExpr, fmt.Errorf("operand B is not supported for Arithmetic operations: %v", err)
	}

	if opA.ValueType != opB.ValueType {
		return resExpr, fmt.Errorf("operands should have same type for arithmetic operation, got %s and %s", opA.ValueType, opB.ValueType)
	}

	resExpr.ValueType = opA.ValueType
	switch opA.ValueType {
	case "INTEGER":
		resExpr.value = fmt.Sprintf("%d", applyArithmeticOp(opA.IntVal(), opB.IntVal(), op))
	case "FLOAT":
		resExpr.value = fmt.Sprintf("%f", applyArithmeticOp(opA.FloatVal(), opB.FloatVal(), op))
	default:
		panic("unreachable code in arithmeticOp")
	}

	return resExpr, nil
}

func logicalOp(opA Expr, opB Expr, op string) (Expr, error) {
	resExpr := Expr{}

	if err := opA.ExpectType("BOOL"); err != nil {
		return resExpr, fmt.Errorf("operand A is not supported for Logical operations: %v", err)
	}
	if err := opB.ExpectType("BOOL"); err != nil {
		return resExpr, fmt.Errorf("operand B is not supported for Logical operations: %v", err)
	}

	resExpr.ValueType = "BOOL"
	switch op {
	case "AND":
		resExpr.value = fmt.Sprintf("%t", opA.BoolVal() && opB.BoolVal())
	case "OR":
		resExpr.value = fmt.Sprintf("%t", opA.BoolVal() || opB.BoolVal())
	default:
		panic("unreachable code in logicalOp")
	}

	return resExpr, nil
}

func comparisionOp(opA Expr, opB Expr, op string) (Expr, error) {
	resExpr := Expr{}
	resExpr.ValueType = "BOOL"

	switch op {
	case ">=", "<=", ">", "<":
		// These are only supported for INTEGER, STRING, FLOAT type
		if err := opA.ExpectTypeOneOf("INTEGER", "FLOAT", "STRING"); err != nil {
			return resExpr, fmt.Errorf("operand A is not supported for Comparision operation with %s: %v", op, err)
		}
		if err := opB.ExpectTypeOneOf("INTEGER", "FLOAT", "STRING"); err != nil {
			return resExpr, fmt.Errorf("operand B is not supported for Comparision operation with %s: %v", op, err)
		}

	case "=", "!=":
		// These are supported for INTEGER, FLOAT, STRING, BOOL
		if err := opA.ExpectTypeOneOf("INTEGER", "FLOAT", "STRING", "BOOL"); err != nil {
			return resExpr, fmt.Errorf("operand A is not supported for Comparision operation with %s: %v", op, err)
		}
		if err := opB.ExpectTypeOneOf("INTEGER", "FLOAT", "STRING", "BOOL"); err != nil {
			return resExpr, fmt.Errorf("operand B is not supported for Comparision operation with %s: %v", op, err)
		}

	default:
		panic("unreachable code in comparisionOp")
	}

	// Check that both operands have same type
	if opA.ValueType != opB.ValueType {
		return resExpr, fmt.Errorf("operands should have same type for comparision operation, got %s and %s", opA.ValueType, opB.ValueType)
	}

	// Apply the comparison operation
	switch opA.ValueType {
	case "INTEGER":
		resExpr.value = fmt.Sprintf("%t", applyComparisonOp(opA.IntVal(), opB.IntVal(), op))
	case "FLOAT":
		resExpr.value = fmt.Sprintf("%t", applyComparisonOp(opA.FloatVal(), opB.FloatVal(), op))
	case "STRING":
		resExpr.value = fmt.Sprintf("%t", applyComparisonOp(opA.StringVal(), opB.StringVal(), op))
	case "BOOL":
		switch op {
		case "=":
			resExpr.value = fmt.Sprintf("%t", opA.BoolVal() == opB.BoolVal())
		case "!=":
			resExpr.value = fmt.Sprintf("%t", opA.BoolVal() != opB.BoolVal())
		default:
			panic("unreachable code in comparisionOp for BOOL type")
		}
	default:
		panic("unreachable code in comparisionOp for ValueType")
	}

	return resExpr, nil
}
