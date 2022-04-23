package main

import "fmt"

type User struct {
	FirstName string
	Age       int
}

func main() {
	myMap := make(map[string]User)

	kedi := User{
		FirstName: "Fahri",
		Age:       9,
	}
	myMap["fahri"] = kedi

	fmt.Println(myMap["fahri"].FirstName)

}
