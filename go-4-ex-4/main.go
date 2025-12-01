package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	fmt.Println(convertCelsiusToFahrenheit(0))   // 32
	fmt.Println(convertCelsiusToFahrenheit(100)) // 212

	// TODO: call the function convertFahrenheitToCelsius
	fmt.Println(convertFahrenheitToCelsius(32))  // 0
	fmt.Println(convertFahrenheitToCelsius(212)) // 100
}