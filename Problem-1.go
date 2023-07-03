package main

import (
	"fmt"
)

type Calculator struct {
	a        float64
	b        float64
	operator string
}

func (c Calculator) calculate() float64 {
	switch c.operator {
	case "+":
		return c.a + c.b
	case "-":
		return c.a - c.b
	case "*":
		return c.a * c.b
	case "/":
		if c.b != 0 {
			return c.a / c.b
		} else {
			fmt.Println("Error: Division by zero!")
			return 0
		}
	default:
		fmt.Println("Error: Invalid operator!")
		return 0
	}
}

func main() {
	calc := Calculator{
		a:        10,
		b:        5,
		operator: "+",
	}
	result := calc.calculate()
	fmt.Printf("Result: %.2f\n", result)
}
