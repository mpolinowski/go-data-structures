package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices as Data Structures")

	// Slices don't have a pre-defined length
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Slice: ", mySlice)

	fmt.Println("First Element: ", mySlice[0])

	fmt.Println("Last Element: ", mySlice[len(mySlice)-1])
}
