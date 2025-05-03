package main

import (
	"fmt"

	"github.com/ANJAN-24/gomaxheap/pkg/heap"
)

func main() {
	fmt.Println("MaxHeap Package Example")
	fmt.Println("=======================")

	// Sample array
	arr := []int{5, 2, 7, 3, 6, 1, 4}
	fmt.Println("Original array:", arr)

	// Create a max heap
	maxHeap := heap.NewMaxHeap(arr)
	fmt.Printf("Heap created with %d elements\n", maxHeap.Size())

	// Get max element
	max, success := maxHeap.Peek()
	if success {
		fmt.Println("Maximum element:", max)
	}

	// Print heap elements in descending order
	fmt.Println("Elements in descending order:")
	for !maxHeap.IsEmpty() {
		element, _ := maxHeap.Extract()
		fmt.Print(element, " ")
	}
	fmt.Println()
}
