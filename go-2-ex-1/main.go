package main

import "fmt"

type FullName struct {
	// TODO: add fields
	firstName string
	lastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	dayOfBirth   int
	monthOfBirth int
	yearOfBirth  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	Name             FullName
	Born             BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		Name: FullName{
			firstName: "Elias",
			lastName:  "Raess",
		},
		Born: BirthDate{
			dayOfBirth:   24,
			monthOfBirth: 10,
			yearOfBirth:  2008,
		},
		NumberOfSiblings: 2,        // TODO: adjust
		ZodiacSign:       '\u264F', // TODO: adjust

	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	fmt.Println("Siblings After:", me.NumberOfSiblings+1)
}
