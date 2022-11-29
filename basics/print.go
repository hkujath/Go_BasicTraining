package basics

import (
	"fmt"
	"os"
)

func LearnPrint() {
	var nr = 1
	var name = "Holger"

	fmt.Print("Simple ", "Print", ".", "\n")
	fmt.Printf("%03d: Hello, %s. Hier is a Printf.\n", nr, name)
	var mySprinf = fmt.Sprintf("%03d: Hello, %s. Hier is a SPrintf.\n", nr, name)
	fmt.Print(mySprinf)

	// Schreiben in eine Datei
	file, _ := os.Create("PrintLog.txt")
	fmt.Fprintf(file, "%03d: Hello, %s. Using a Fprintf.\n", nr, name)
	file.Close()

}
