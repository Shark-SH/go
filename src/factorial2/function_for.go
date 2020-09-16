package main

import "fmt"

func function(n int) int {
	result := 1
	for n > 0 {

		result *= n
		n--
	}
	return result
}

func main() {

	fmt.Println(function(5))
	fmt.Println("for문을 이용한 factorial")
}
