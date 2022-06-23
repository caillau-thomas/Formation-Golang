package main

import (
	"fmt"
)

func myAgeFunc(age int) {
	fmt.Printf("\n\n°myAgeFunc\n\n")
	if age > 18 {
		fmt.Printf("Je suis majeur\n\n")
	} else if age == 18 {
		fmt.Printf("Je viens d'etre majeur\n\n")
	} else {
		fmt.Printf("Je sui mineur\n\n")
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("La valeur de i est: %v \n\n", i)
	}
}
