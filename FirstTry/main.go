// Tuto getCodingKnowledge Golang playlist //
// https://youtube.com/playlist?list=PLuWyq_EO5_AKP_KCaIr53UfOqlPThTXat //

package main

import (
	"fmt"
)

var globalVar string

func main() {
	myCaculFunc(2, 2)
	myRandNumberFunc()
	myAgeFunc(18)
	myArrayFunc()
	fmt.Println(sayByeFunc("Alex"))

	globalVar = "Thomas"
	sayHello()
	mysuperMarketMapPriceFunc()
	myCalendarFunc()
	annonyme()

	nbr := 10
	fmt.Printf("nbr: %v", nbr)
	nbr = updateValue(nbr)
	fmt.Printf("nbr: %v", nbr)
}

/////////////////////////////////////////////////////////////
func mysuperMarketMapPriceFunc() {
	superMarketPrice := map[string]int{
		"item1": 1,
		"item2": 2,
		"item3": 3,
	}
	fmt.Printf("\n\n°mysuperMarketMapPriceFunc\n\n")
	fmt.Println(superMarketPrice["item1"])
	fmt.Printf("\n\n")
	updateList(superMarketPrice)
	for key, value := range superMarketPrice {
		fmt.Printf("Item %v Price %v € \n\n", key, value)
	}
}

/////////////////////////////////////////////////////////////

func annonyme() {
	//Fonctions Annonymes
	w := func() {
		fmt.Printf("\n\n°Fonction annonyme\n\n")
		println("Ma fonction anon")
	}
	w()
	z := func() string {
		fmt.Printf("\n\n°Fonction annonyme avec return\n\n")
		return "valeur de la fonction => Thomas"
	}() // ()permet d'appeler dirrectemement la func
	println(z)

	x, y := func() (int, int) {
		fmt.Printf("\n\n°Fonction annonyme avec 2 return\n\n")
		return 5, 7
	}()
	fmt.Printf("valeur x => %v , valeur y => %v", x, y)

	func(a, b int) {
		fmt.Printf("\n\n°Fonction annonyme avec calcule\n\n")
		println("a * a + b * b = ", a*a+b*b)
	}(x, y) // on instencie a et b avec les valeurs de x et y
}

/////////////////////////////////////////////////////////////

func updateValue(nbr int) int {
	fmt.Printf("\n\n°Fonction passByValue on update une valeur en memoire\n\n")
	nbr = 5
	return nbr
}

/////////////////////////////////////////////////////////////

func updateList(item map[string]int) {
	fmt.Printf("\n\n°Fonction passByValue on update une list en ajoutant un item\n\n")
	item["item4"] = 4
}
