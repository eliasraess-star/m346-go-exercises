package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		firstName string
		lastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		name     string
		students []Student
	}
	classA := Class{
		name: "Class A",
		students: []Student{
			{firstName: "Elias", lastName: "Raess"},
			{firstName: "Max", lastName: "Muster"},
			{firstName: "John", lastName: "Doe"},
		},
	}
	classB := Class{
		name: "Class B",
		students: []Student{
			{firstName: "Jean", lastName: "Dupont"},
			{firstName: "Mario", lastName: "Rossi"},
			{firstName: "Jane", lastName: "Public"},
		},
	}
	// TODO: declare a map of modules being attended by multiple classes
	modules := map[int][]Class{
		104: {classA, classB},
		117: {classB},
		346: {classA},
	}
	// TODO: output everything using fmt.Println()
	for moduleID, classes := range modules {
		fmt.Println("Modul:", moduleID)
		for _, class := range classes {
			fmt.Println(" Klasse:", class.name)
			for _, student := range class.students {
				fmt.Println("  -", student.firstName, student.lastName)
			}
		}
		fmt.Println()
	}
}
