package main

import "fmt"

func main() {
	//经典的for循环 init;condition;post
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("************************************")

	//经典的for语句 condition
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("*************************************")

	//死循环的for语句 相当于for(;;)
	j := 1
	for {
		if j > 10 {
			break
		} else {
			fmt.Println("haha")
			j++
		}
	}
}
