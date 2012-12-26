package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	b := a[2:4] //a[2]和a[3],但不包括a[4]
	fmt.Println(b)

	b = a[:4] //从a[0]到a[4],但不包括a[4]
	fmt.Println(b)

	b = a[2:] //从a[2]到a[4],切包括a[2]
	fmt.Println(b)
}
