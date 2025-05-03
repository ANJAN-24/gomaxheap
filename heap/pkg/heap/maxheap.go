// Package heap provides heap data structure implementations
package heap

// MaxHeap represents a max heap data structure
type MaxHeap struct {
	array []int
}

// NewMaxHeap creates a new max heap from an existing array
func NewMaxHeap(arr []int) *MaxHeap {
	// Create a copy of the input array
	heapArray := make([]int, len(arr))
	copy(heapArray, arr)

	// Create the heap
	heap := &MaxHeap{
		array: heapArray,
	}

	// Heapify the array (bottom-up approach)
	for i := len(heap.array)/2 - 1; i >= 0; i-- {
		heap.heapifyDown(i)
	}

	return heap
}

// Insert adds a new element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

// Extract removes and returns the maximum element
func (h *MaxHeap) Extract() (int, bool) {
	if len(h.array) == 0 {
		return 0, false
	}

	max := h.array[0]
	lastIdx := len(h.array) - 1
	h.array[0] = h.array[lastIdx]
	h.array = h.array[:lastIdx]

	if len(h.array) > 0 {
		h.heapifyDown(0)
	}

	return max, true
}

// Peek returns the maximum element without removing it
func (h *MaxHeap) Peek() (int, bool) {
	if len(h.array) == 0 {
		return 0, false
	}
	return h.array[0], true
}

// Size returns the number of elements in the heap
func (h *MaxHeap) Size() int {
	return len(h.array)
}

// IsEmpty returns true if the heap is empty
func (h *MaxHeap) IsEmpty() bool {
	return len(h.array) == 0
}

// GetArray returns a copy of the underlying array
func (h *MaxHeap) GetArray() []int {
	result := make([]int, len(h.array))
	copy(result, h.array)
	return result
}

// heapifyUp maintains the heap property after insertion
func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIdx := (index - 1) / 2
		if h.array[parentIdx] >= h.array[index] {
			break
		}
		h.array[index], h.array[parentIdx] = h.array[parentIdx], h.array[index]
		index = parentIdx
	}
}

// heapifyDown maintains the heap property after extraction
func (h *MaxHeap) heapifyDown(index int) {
	lastIdx := len(h.array) - 1
	l, r := 2*index+1, 2*index+2
	childToCompare := 0

	// Loop while index has at least one child
	for l <= lastIdx {
		if l == lastIdx { // When left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // When left child is larger
			childToCompare = l
		} else { // When right child is larger
			childToCompare = r
		}

		// Compare with current index
		if h.array[index] > h.array[childToCompare] {
			break
		}

		// Swap
		h.array[index], h.array[childToCompare] = h.array[childToCompare], h.array[index]
		index = childToCompare
		l, r = 2*index+1, 2*index+2
	}
}
