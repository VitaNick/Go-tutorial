package main

import (
	"fmt"
	"go-functions/simplemath"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	fmt.Printf("%f\n", double(3, 2, mathExpression(SubtractExpr)))
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {

	case AddExpr:
		return simplemath.Add

	case SubtractExpr:
		return simplemath.Subtract

	case MultiplyExpr:
		return simplemath.Multiply

	default:
		return func(f, f2 float64) float64 {
			return 0
		}
	}
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}
