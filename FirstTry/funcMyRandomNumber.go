package main

import (
	"fmt"
	"math/rand"
	"time"
)

func myRandNumberFunc() {
	rand.Seed(time.Now().UnixNano())
	z := rand.Int()

	fmt.Printf("\n\n°myRandNumberFunc\n\n")
	fmt.Printf("My random number is: %v \n\n", z)
	/* Ou on peut dirrectement appeler la rand.Int() */
	/* fmt.Printf("My random number is: %v", rand.Int()) */

	if z%2 == 0 { /* On verifie que z est paire ou impaire */
		fmt.Printf("%v est paire \n\n", z)
	} else {
		fmt.Printf("%v est impaire\n\n", z)
	}
}
