package main

import (
	"fmt"
)

// Temperature struct with a value field
type Temperature struct {
	Value float64
}

// ToFahrenheit converts from Celsius to Fahrenheit
func (t *Temperature) ToFahrenheit() {
	t.Value = (t.Value * 9 / 5) + 32
}

// ToCelsius converts from Fahrenheit to Celsius
func (t *Temperature) ToCelsius() {
	t.Value = (t.Value - 32) * 5 / 9
}

func main() {
	// User Input
	var tempValue float64
	var choice int

	fmt.Print("Enter temperature value: ")
	fmt.Scan(&tempValue)

	// Create a Temperature instance
	temp := &Temperature{Value: tempValue}

	// Conversion Menu
	fmt.Println("Choose conversion:")
	fmt.Println("1: Convert to Fahrenheit")
	fmt.Println("2: Convert to Celsius")
	fmt.Scan(&choice)

	// Perform conversion
	switch choice {
	case 1:
		temp.ToFahrenheit()
		fmt.Printf("Temperature in Fahrenheit: %.2f°F\n", temp.Value)
	case 2:
		temp.ToCelsius()
		fmt.Printf("Temperature in Celsius: %.2f°C\n", temp.Value)
	default:
		fmt.Println("Invalid choice.")
	}
}
