package expr

import (
	"fmt"
	"strconv"
)

// Accessor methods for values of Expr
// Should be called only when ValueType is checked previously in the code
func (e *Expr) StringVal() string {
	if e.ValueType != "STRING" {
		panic(fmt.Sprintf("Trying to get string value of a non-string type: %s for Expr %v", e.ValueType, e))
	}
	return e.value
}

func (e *Expr) IntVal() int {
	if e.ValueType != "INT" {
		panic(fmt.Sprintf("Trying to get int value of a non-int type: %s for Expr %v", e.ValueType, e))
	}
	intVal, err := strconv.Atoi(e.value)
	if err != nil {
		panic(fmt.Sprintf("Error converting string to int: %v", err))
	}
	return intVal
}

func (e *Expr) BoolVal() bool {
	if e.ValueType != "BOOL" {
		panic(fmt.Sprintf("Trying to get bool value of a non-bool type: %s for Expr %v", e.ValueType, e))
	}
	boolVal, err := strconv.ParseBool(e.value)
	if err != nil {
		panic(fmt.Sprintf("Error converting string to bool: %v", err))
	}
	return boolVal
}

func (e *Expr) FloatVal() float64 {
	if e.ValueType != "FLOAT" {
		panic(fmt.Sprintf("Trying to get float value of a non-float type: %s for Expr %v", e.ValueType, e))
	}
	floatVal, err := strconv.ParseFloat(e.value, 64)
	if err != nil {
		panic(fmt.Sprintf("Error converting string to float: %v", err))
	}
	return floatVal
}
