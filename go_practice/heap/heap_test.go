package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &stockPrice{
		priority: 10,
	})
	heap.Push(&pq, &stockPrice{
		priority: 20,
	})
	heap.Push(&pq, &stockPrice{
		priority: 5,
	})

	for _, s := range pq {
		fmt.Printf("%+v\n", *s)
	}

}
