package main

import (
	"fmt"

	"github.com/ANJAN-24/gomaxheap/pkg/heap"
)

func main() {
	// Example array
	arr := []int{5, 2, 7, 3, 6, 1, 4}
	fmt.Println("Original array:", arr)

	// Build a max heap from the array
	maxHeap := heap.NewMaxHeap(arr)
	fmt.Println("Heap created. Size:", maxHeap.Size())
	fmt.Println("Heap array representation:", maxHeap.GetArray())

	// Extract elements in descending order
	fmt.Println("\nExtracting elements in descending order:")
	for !maxHeap.IsEmpty() {
		max, _ := maxHeap.Extract()
		fmt.Print(max, " ")
	}
	fmt.Println()

	// Create a heap and insert elements
	maxHeap2 := heap.NewMaxHeap([]int{})
	fmt.Println("\nCreating a new empty heap and inserting elements:")
	maxHeap2.Insert(10)
	maxHeap2.Insert(30)
	maxHeap2.Insert(20)
	maxHeap2.Insert(5)
	maxHeap2.Insert(15)

	// Peek at the maximum element
	max, _ := maxHeap2.Peek()
	fmt.Println("Maximum element:", max)

	// Extract elements in descending order
	fmt.Println("Elements in descending order:")
	for !maxHeap2.IsEmpty() {
		max, _ := maxHeap2.Extract()
		fmt.Print(max, " ")
	}
	fmt.Println()
}
