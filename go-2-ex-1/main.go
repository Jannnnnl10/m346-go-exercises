package main

import "fmt"

type FullName struct {
	// TODO: add fields
}

// TODO: declare a structure for birth date

type Profile struct {
	// TODO: embed full name and birth date information
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		NumberOfSiblings: 2,   // TODO: adjust
		ZodiacSign:       'Schuetze', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
