package main

import (
	sorting "dsa/Sort"
	"dsa/search"
	"fmt"
)

func main() {

	a := []int{10, 8, 07, 60, 5, 4, 3, 2, 1, -6}

	b := []int{10, 8, 07, 60, 5, 4, 3, 2, 1, -6}

	c := []int{10, 8, 07, 60, 5, 4, 3, 2, 1, -6}

	d := []int{10, 8, 07, 60, 5, 4, 3, 2, 1, -6}

	fmt.Println("before sort ", a)

	sorting.Bubble(a)
	fmt.Println("bubble sort ", a)

	sorting.QuickSort(b, 0, len(b)-1)
	fmt.Println("Quick sort", b)

	sorting.MergeSort(c)
	fmt.Println("Merge sort", c)

	fmt.Println(search.LinearSearch(d, 07))

	fmt.Println(search.BinarySearch(a, 60))

}
