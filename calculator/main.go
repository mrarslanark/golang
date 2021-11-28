/*
Written in Go v1.17
Code Author: Arslan Mushtaq
Contact Email: arslan@devden.org
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Create a console input reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("A Simple Calculator")

	// Get numbers from the user
	float1 := getValues(reader, "Value 1: ")
	float2 := getValues(reader, "Value 2: ")

	// Get operation from the user
	operationInput := getOperation(reader)

	// Create the result and set the operation
	// based on the selected operation
	var result float64
	var operation string
	switch strings.TrimSpace(operationInput) {
		case "+":
			result, operation = sum(float1, float2)
		case "-":
			result, operation = subtract(float1, float2)
		case "*":
			result, operation = multiply(float1, float2)
		case "/":
			result, operation = division(float1, float2)
		default:
			panic("Unknown identifier")
	}
	// Round the result
	result = math.Round(result * 100) / 100

	// Output the result
	fmt.Printf("Values: %.1f & %.1f\n", float1, float2)
	fmt.Printf("Selected Operation: %v\n", operation)
	fmt.Printf("Result: %.1f\n", result)
	t := time.Now()
	fmt.Println("Timestamp: ", t.Format(time.ANSIC))
}

// Read a value from the input console 
func getValues(reader *bufio.Reader, prompt string) (float64) {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value should be a number")
	}
	return number;
}

// Read an operation from the input console
func getOperation(reader *bufio.Reader) string {
	fmt.Print("Select an operation (+ - * /): ")
	operation, err := reader.ReadString('\n')
	if (err != nil) {
		fmt.Println(err)
		panic("Operation is required")
	}
	return strings.TrimSpace(operation);
}

// Add two values and return the result
// along with operation type
func sum(value1 float64, value2 float64) (float64, string) {
	return value1 + value2, "Addition"
}

// Subtracts two values and return the result
// along with operation type
func subtract(value1 float64, value2 float64) (float64, string) {
	return value1 - value2, "Subtraction"
}

// Multiplies two values and return the result
// along with operation type
func multiply(value1 float64, value2 float64) (float64, string) {
	return value1 * value2, "Multiplication"
}

// Divides two values and return the result
// along with operation type
func division(value1 float64, value2 float64) (float64, string) {
	return value1 / value2, "Division"
}