package main

import "fmt"

func main() {
	animals := []string{"cow", "cat", "dog", "mouse"}

	for animals := range animals {
		fmt.Println(animals)
	}
}
