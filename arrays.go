package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays as Data Structures")

	myArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Array: ", myArray)

	fmt.Println("First Element: ", myArray[0])

	fmt.Println("Last Element: ", myArray[len(myArray)-1])
}
