package main

import (
	"fmt"
	"learn1/insertion"
)

func main() {
	var arr = []int{9, 6, 5, 0, 8, 2, 7, 1, 3, 3, 4}
	fmt.Println("Before Sorting: ", arr)

	insertion.Sort(arr)

	fmt.Println("After Sorting: ", arr)
}
