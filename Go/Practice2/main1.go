package main

import "fmt"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	m.FirstName = "Eyyub Ensar"
	return m.FirstName
}

func main() {

	var myVar myStruct
	myVar.FirstName = "Eyyub Ensar"
	fmt.Println("var1 ", myVar.printFirstName())

	myVar2 := myStruct{
		FirstName: "Eyyub Ensar MERMER",
	}

	fmt.Println("var2", myVar2.printFirstName())
}
