package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "Datenmodell implementieren",
		117: "Verschlüsselungsverfahren einsetzen",
		346: "Cloud Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 104)
	// TODO: add one
	modules[320] = "Objekt Orientiert programmieren"
	// TODO: replace one
	modules[104] = "SQL Datenbanken aufsetzen"
	fmt.Println(modules)
}
