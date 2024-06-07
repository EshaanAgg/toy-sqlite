package expr

import (
	"fmt"
	"strings"
)

// Function to be called when the expression is of type CALL
func (e *Expr) evalCall() error {
	// Evaluate the operands first
	err := e.evalOperands()
	if err != nil {
		return fmt.Errorf("failed to evaluate operands: %v", err)
	}

	// Call the appropriate function based on the function name
	var resExpr Expr
	switch e.functionName {
	case "UPPER":
		resExpr, err = upperExprFunc(e.operands)
	case "LOWER":
		resExpr, err = lowerExprFunc(e.operands)
	case "CONCAT":
		resExpr, err = concatExprFunc(e.operands)
	default:
		return fmt.Errorf("unknown function name: %s", e.functionName)
	}

	if err != nil {
		return fmt.Errorf("failed to evaluate function %s: %v", e.functionName, err)
	}

	// Update the expression with the result
	*e = resExpr
	return nil
}

// Performs the UPPER function on the operands
func upperExprFunc(operands []Expr) (Expr, error) {
	resExpr := Expr{}

	// Check that the number of operands is exactly 1
	// and that the operand is of type STRING
	if len(operands) != 1 {
		return resExpr, fmt.Errorf("expected 1 operand, got %d", len(operands))
	}
	if operands[0].ValueType != "STRING" {
		return resExpr, fmt.Errorf("expected operand to be of type STRING, got %s", operands[0].ValueType)
	}

	resExpr.ValueType = "STRING"
	resExpr.value = strings.ToUpper(operands[0].value)
	return resExpr, nil
}

// Performs the LOWER function on the operands
func lowerExprFunc(operands []Expr) (Expr, error) {
	resExpr := Expr{}

	// Check that the number of operands is exactly 1
	// and that the operand is of type STRING
	if len(operands) != 1 {
		return resExpr, fmt.Errorf("expected 1 operand, got %d", len(operands))
	}
	if operands[0].ValueType != "STRING" {
		return resExpr, fmt.Errorf("expected operand to be of type STRING, got %s", operands[0].ValueType)
	}

	resExpr.ValueType = "STRING"
	resExpr.value = strings.ToLower(operands[0].value)
	return resExpr, nil
}

// Performs the CONCAT function on the operands
func concatExprFunc(operands []Expr) (Expr, error) {
	resExpr := Expr{}

	// Check that the number of operands is at least 2
	// and that all operands are of type STRING
	if len(operands) < 2 {
		return resExpr, fmt.Errorf("expected at least 2 operands, got %d", len(operands))
	}
	for _, operand := range operands {
		if operand.ValueType != "STRING" {
			return resExpr, fmt.Errorf("expected all operands to be of type STRING, got atleast one operand of type %s", operand.ValueType)
		}
	}

	resExpr.ValueType = "STRING"
	for _, operand := range operands {
		resExpr.value += operand.value
	}
	return resExpr, nil
}
