package main

import "fmt"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	m.FirstName = "Berkan"
	return m.FirstName
}

func main() {

	var myVar myStruct
	myVar.FirstName = "Berkan"
	fmt.Println("var1 ", myVar.printFirstName())

	myVar2 := myStruct{
		FirstName: "Berkan mutlu",
	}

	fmt.Println("var2", myVar2.printFirstName())
}
