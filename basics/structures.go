package basics

import "fmt"

func LearnStructures() {

	holger := user{
		name: "Holger",
		adresse: adresse{
			strasse: "MyWay 1",
			stadt:   "MyCity 12345",
		},
	}

	fmt.Printf("User: %v\n", holger)
	fmt.Printf("User: %#v\n", holger)
	fmt.Printf("User: %+v\n", holger)

	// Anonymous structs

	data := struct {
		input    string
		expected string
	}{
		input:    "Test",
		expected: "MyData",
	}

	fmt.Printf("Anonymous struct: %+v\n", data)
}

type adresse struct {
	strasse string
	stadt   string
}

type user struct {
	name string
	adresse
}

type AdressJson struct {
	Street string `json:"street"`
	City   string `json:"city"`
}
