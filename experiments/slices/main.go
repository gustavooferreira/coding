package main

import (
	"fmt"
)

func main() {
	// create a slice (nil)
	var mySlice []int
	printSlice(mySlice)

	for i := 1; i <= 10; i++ {
		mySlice = append(mySlice, i)
		printSliceInternals(mySlice)
	}

	fmt.Println("-------------------")

	// modify a slice inside array
	mySlice2 := make([]int, 1, 10)
	printSliceInternals(mySlice2)
	changeSlice(mySlice2)
	printSliceInternals(mySlice2)

}

func changeSlice(slice []int) {
	slice = append(slice, 1)
}

// printSliceInternals prints the a slice's length, capacity and the underlying array pointer.
func printSliceInternals(mySlice []int) {
	fmt.Printf("Len: %3d --- Cap: %3d --- Address: %p --- First array elem: %p\n", len(mySlice), cap(mySlice), (*[0]int)(mySlice), &mySlice[0])
}

func printSlice(mySlice []int) {
	fmt.Printf("Len: %3d --- Cap: %3d --- Slice: %+v\n", len(mySlice), cap(mySlice), mySlice)
}
