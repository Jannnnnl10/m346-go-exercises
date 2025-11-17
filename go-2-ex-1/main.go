package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName string
	
}

// TODO: declare a structure for birth date

type Profile struct {
	// TODO: embed full name and birth date information
	FirstName string
	LastName string
	DayOfBirth int
	MonthOfBirth string
	YearOfBirth int 
	NumberOfSiblings byte
	HeightInMeters float64
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FirstName: "Jan",
		LastName: "Lehmann",
		DayOfBirth: 29,
		MonthOfBirth: "November",
		YearOfBirth: 2008,
		NumberOfSiblings: 2,
		HeightInMeters: 1.85, 
		ZodiacSign:       '\u2650',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
