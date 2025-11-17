package main

import "fmt"

func main() {
	var firstName string = "Jan"
	var lastName string = "Lehmann"
	var dayOfBirth int = 29
	var monthOfBirth string = "November"
	var yearOfBirth int = 2008
	var numberOfSiblings int = 2
	var heightInMeters float64 = 1.85
	var zodiacSign rune = '\u2650'

	fmt.Printf("Name: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d. %s %d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse: %.2f m\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
// go run go-1-ex-1/main.go 