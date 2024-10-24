package main

import (
	"fmt"
	"testing"
)

func TestSol(t *testing.T) {

	r := solution(90, 500, []int{70, 70, 0}, []int{0, 0, 500}, []int{100, 100, 2}, []int{4, 8, 1})
	fmt.Println(r)
}
