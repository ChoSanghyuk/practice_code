package main

import (
	"fmt"
	"testing"
)

func TestSol(t *testing.T) {

	r := solution(
		[]string{"john", "mary", "edward", "sam", "emily", "jaimie", "tod", "young"},
		[]string{"-", "-", "mary", "edward", "mary", "mary", "jaimie", "edward"},
		[]string{"young", "john", "tod", "emily", "mary"},
		[]int{12, 4, 2, 5, 10})
	fmt.Println(r)
}
