package main

import "fmt"

func main() {
	a := 4
	var b = 2
	//선언 대입문

	fmt.Printf("%v\n", a&b)
	fmt.Printf("%v\n", a|b)
}
// fmt 라이브러리를 import 하는 경우 println 사용시
// 대소문자 구분 안하면 저장하는 경우import "fmt"가 사라진다. 
// 앞으로 참 무섭다......
