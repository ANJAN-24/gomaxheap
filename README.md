# Go MaxHeap Package

A simple and efficient implementation of a Max Heap data structure in Go.

## Installation

### Local Usage
The package is currently set up for local development:

```bash
# Clone the repository
git clone <your-repo-url>
cd heap

# Run the example
go run main.go
```

### Publishing (Future)
When you're ready to publish this package, update the module name in go.mod:

```
module github.com/yourusername/heap
```

Then users can install it with:
```bash
go get github.com/yourusername/heap
```

Remember to replace `yourusername` with your actual GitHub username when you publish this package.

## Features

- Create a max heap from an existing array
- Insert elements into the heap
- Extract the maximum element
- Peek at the maximum element without removing it
- Check heap size and if it's empty
- Get the underlying array representation

## Usage

### Local Usage
```go
package main

import (
    "fmt"
    "heap/pkg/heap"
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

### After Publishing
```go
package main

import (
    "fmt"
    "github.com/yourusername/heap/pkg/heap"
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
    
    // Create an empty heap and insert elements
    heap2 := heap.NewMaxHeap([]int{})
    heap2.Insert(10)
    heap2.Insert(30)
    heap2.Insert(20)
    
    // Peek at the maximum element
    max, _ := heap2.Peek()
    fmt.Println("Maximum element:", max) // Output: 30
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
