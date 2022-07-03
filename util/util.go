package util

import "fmt"

var i = 0
var F = func(s string) int {
	i++
	fmt.Printf("本次被%s調用\n", s)
	fmt.Printf("匿名工具函式被調用%v次\n", i)
	return i
}

func ForTest20(text ...string) {
	for i, v := range text {
		fmt.Printf("i=%d, v=%v\n", i, v)
	}
}
