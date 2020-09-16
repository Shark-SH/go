package main

import (
	"fmt"
)

func main() {
	s1 := "타잔이"
	s2 := "원 짜리 팬티를 입고,"
	s3 := "원 짜리 칼을 차고 노래를 한다."
	x := 10
	y := x + 10

	for i := 1; i <= 10; i++ {
		x = i * 10
		fmt.Print(s1)
		fmt.Print(x)
		fmt.Print(s2)
		fmt.Print(y)
		fmt.Print(s3)
		fmt.Println()
	}

}
