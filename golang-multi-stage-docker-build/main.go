package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var op string
	var a, b float64

	flag.StringVar(&op, "op", "", "operation: add, sub, mul, div")
	flag.Float64Var(&a, "a", 0, "first number")
	flag.Float64Var(&b, "b", 0, "second number")
	flag.Parse()

	if op == "" || (op != "add" && op != "sub" && op != "mul" && op != "div") {
		fmt.Println("Operation must be one of: add, sub, mul, div")
		os.Exit(1)
	}

	var result float64
	switch op {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		if b == 0 {
			fmt.Println("Division by zero is not allowed")
			os.Exit(1)
		}
		result = a / b
	}

	fmt.Printf("Result: %f\n", result)
}
