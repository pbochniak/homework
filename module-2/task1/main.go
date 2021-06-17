package main

import (
	"fmt"
	"strconv"

	"github.com/pbochniak/computator/fibonacci"
)

func castNumber(input string) (interface{}, error) {
	if u, err := strconv.ParseUint(input, 10, 0); err == nil {
		return uint(u), nil
	}
	if i, err := strconv.ParseInt(input, 10, 0); err == nil {
		return int(i), nil
	}
	if f, err := strconv.ParseFloat(input, 64); err == nil {
		return f, nil
	}
	return nil, fmt.Errorf("invalid type")
}

func main() {
	var inputNumber string
	fmt.Printf("Give me a number: ")
	fmt.Scanf("%s", &inputNumber)
	number, err := castNumber(inputNumber)
	if err != nil {
		fmt.Printf("Error: %q isn't a correct number.\n", inputNumber)
		return
	}
	fmt.Printf("Fibonacci sequence for given number (%v [%T]) equals %v\n", number, number, fibonacci.Fibonacci(number))
}
