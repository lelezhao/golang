package main

import "fmt"

func main() {
	//声明初始化一个变量
	var x int = 100
	var str string = "hello,iflytek!"
	//声明初始化多个变量
	var i, j, k int = 1, 2, 3
	//不要指明类型，通过初始化值来推导
	var b = true //bool型
	//另一种定义方式
	y := 100 //等价于var y int = 100
	//定义常量，使用const关键字
	const s string = "hello,world!"
	const pi float32 = 3.1415926

	fmt.Println(str)
}
