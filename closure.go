package main

import "fmt"

func nextNum() func() int {
	i, j := 1, 1
	return func() int {
		var temp = i + j
		i, j = j, temp
		return temp
	}
}

//main函数中是对nextNum的调用,其主要是打出下一个斐波那契数
func main() {
	nextNumFunc := nextNum()
	for i := 0; i < 10; i++ {
		fmt.Println(nextNumFunc())
	}
}
