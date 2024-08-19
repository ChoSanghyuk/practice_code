package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line1 := strings.Split(scanner.Text(), " ")

	N, _ := strconv.Atoi(line1[0])
	K, _ := strconv.Atoi(line1[1])

	pq_r := make(PriorityQueue, 0)
	heap.Init(&pq_r)

	pq_l := make(PriorityQueue, 0)
	heap.Init(&pq_l)

	for _ = range N {
		scanner.Scan()
		i, _ := strconv.Atoi(scanner.Text())
		if i > 0 {
			heap.Push(&pq_r, &Item{
				value:    i,
				priority: i,
			})
		} else {
			heap.Push(&pq_l, &Item{
				value:    -1 * i,
				priority: -1 * i,
			})
		}
	}

	var rtn *big.Int = big.NewInt(0)

	idx := 0
	for pq_r.Len() > 0 {
		if idx%K == 0 {
			rtn = rtn.Add(rtn, big.NewInt(2*int64(heap.Pop(&pq_r).(*Item).value)))
		} else {
			_ = heap.Pop(&pq_r)
		}
		idx++
	}

	idx = 0
	for pq_l.Len() > 0 {
		if idx%K == 0 {
			rtn = rtn.Add(rtn, big.NewInt(2*int64(heap.Pop(&pq_l).(*Item).value)))
		} else {
			_ = heap.Pop(&pq_l)
		}
		idx++
	}

	fmt.Println(rtn)

}
