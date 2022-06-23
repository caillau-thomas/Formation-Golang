package main

import "fmt"

func myArrayFunc() {
	fmt.Printf("\n\n°myArrayFunc\n\n")
	var list [3]int
	list[0] = 10
	list[1] = 20
	list[2] = 30

	fmt.Println(list)
	fmt.Printf("\n\n")

	list2 := [...]int{10, 20, 30, 40, 50, 60, 70, 80}

	fmt.Println(list2)
	fmt.Printf("\n\n")

	for pos, value := range list2 {
		fmt.Printf("Position %d est egale à %d \n\n", pos, value)
	}
}
