package main

import (
	"fmt"
	"sort"
)

func main() {
	var mySlice []int

	mySlice = append(mySlice, 6, 4, 92, 21, 10)

	sort.Ints(mySlice)

	fmt.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}

	fmt.Println(numbers[4:5])
}
