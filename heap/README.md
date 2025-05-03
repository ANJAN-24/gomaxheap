# Go MaxHeap Package

A simple and efficient implementation of a Max Heap data structure in Go.

## Installation

You can install this package with the go get command:

```bash
go get github.com/ANJAN-24/gomaxheap
```

## Features

- Create a max heap from an existing array
- Insert elements into the heap
- Extract the maximum element
- Peek at the maximum element without removing it
- Check heap size and if it's empty
- Get the underlying array representation

## Usage

```go
package main

import (
    "fmt"
    "github.com/ANJAN-24/gomaxheap/pkg/heap"
)

func main() {
    // Create a max heap from an array
    arr := []int{5, 2, 7, 3, 6, 1, 4}
    maxHeap := heap.NewMaxHeap(arr)
    
    // Get the size of the heap
    fmt.Println("Heap size:", maxHeap.Size())
    
    // Get the underlying array
    fmt.Println("Heap array:", maxHeap.GetArray())
    
    // Extract elements in descending order
    for !maxHeap.IsEmpty() {
        max, _ := maxHeap.Extract()
        fmt.Print(max, " ") // Output: 7 6 5 4 3 2 1
    }
}
```

## API

### NewMaxHeap(arr []int) *MaxHeap

Creates a new max heap from an existing array. Time complexity: O(n).

### Insert(key int)

Adds a new element to the heap. Time complexity: O(log n).

### Extract() (int, bool)

Removes and returns the maximum element. The boolean value indicates whether the operation was successful. Time complexity: O(log n).

### Peek() (int, bool)

Returns the maximum element without removing it. The boolean value indicates whether the operation was successful.

### Size() int

Returns the number of elements in the heap.

### IsEmpty() bool

Returns true if the heap is empty.

### GetArray() []int

Returns a copy of the underlying array representation of the heap.

## License

[MIT License](LICENSE) 