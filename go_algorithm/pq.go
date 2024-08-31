package main

type Item struct {
	value    int // The value of the item; arbitrary data
	priority int // The priority of the item in the queue
	index    int // The index of the item in the heap (required by `container/heap`)
}

type PriorityQueue []*Item

// Len returns the number of items in the queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less compares the priority of two items (higher priority comes first).
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority // For max-heap (higher priority first)
	// return pq[i].priority < pq[j].priority // For min-heap (lower priority first)
}

// Swap swaps the position of two items in the queue.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds a new item to the queue.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the item with the highest priority.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1 // For safety
	*pq = old[0 : n-1]
	return item
}
