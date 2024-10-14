package main

import "strconv"

/*
양의 정수 x에 대한 함수 f(x)를 다음과 같이 정의합니다.

x보다 크고 x와 비트가 1~2개 다른 수들 중에서 제일 작은 수
*/

func solution(numbers []int64) []int64 {

	rtn := make([]int64, len(numbers))

	for i, n := range numbers {
		rtn[i] = f(n)
	}

	return rtn
}

func f(n int64) int64 {

	if n%2 == 0 {
		return n + 1
	} else {
		s := strconv.FormatInt(n, 2)
		b := []byte(s)
		b = append([]byte{'0'}, b...)

		for i := len(b) - 1; i >= 0; i-- {
			if b[i] == '0' {
				b[i] = '1' // 홀수일때는 끝자리가 0으로 시작하지 않음
				b[i+1] = '0'
				rtn, _ := strconv.ParseInt(string(b), 2, 64)
				return rtn
			}
		}
	}
	return -1
}
