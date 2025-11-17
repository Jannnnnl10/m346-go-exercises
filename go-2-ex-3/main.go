package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
modules := map[int]string{
	104: "Datenmodell implementieren",
	117: "Informatik und Netzinfrastruktur",
	346: "Cloud Lösungen",
}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 104)
	// TODO: add one
	modules[202] = "Softwarearchitektur"
	// TODO: replace one
	modules[202] = "Testen"
	fmt.Println(modules)
}
// go run go-2-ex-3/main.go