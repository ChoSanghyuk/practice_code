package main

func solution(n int, left int64, right int64) []int {

	rtn := make([]int, right-left+1)
	n2 := int64(n)

	for i := range rtn {
		rtn[i] = getNum(n2, left)
	}
	return rtn
}

func getNum(n int64, t int64) int {

	r := t / n
	c := t % n

	return max(int(r), int(c)) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
