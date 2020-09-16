package main

import "fmt"

// fac returns n

func fac(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fac(n-1)

	//reutrn 5 * fac함수 호출 fac(5-1 = 4) ..... if 조건판단 retrun 1
}

func main() {
	fmt.Println(fac(5))
}

//factorial 재귀함수 호출
