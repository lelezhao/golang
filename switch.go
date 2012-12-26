package main

import "fmt"

func main() {
	x := 3
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5, 6:
		fmt.Println("four,five,six")
	default:
		fmt.Println("invalid value!")
	}
}
