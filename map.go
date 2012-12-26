package main

import "fmt"

func main() {
	m := make(map[string]int) //使用make创建一个新的map

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m)      //输出map,每次运行输出顺序可能不同
	fmt.Println(len(m)) //输出map的大小

	v := m["two"]  //从map里取值
	fmt.Println(v) //输出2

	delete(m, "two") //删除一个键值对
	fmt.Println(m)

	//另一种形式的初始化
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(m1) //输出map

	//for循环遍历map并输出
	for key, val := range m1 {
		fmt.Printf("%s => %d\n", key, val) //输出
	}
}
