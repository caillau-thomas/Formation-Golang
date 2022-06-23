// Tuto getCodingKnowledge Golang playlist //
// https://youtube.com/playlist?list=PLuWyq_EO5_AKP_KCaIr53UfOqlPThTXat //

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

/*//////////////////////// FUNCTIONS ///////////////////////*/

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

/////////////////////////////////////////////////////////////

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

/////////////////////////////////////////////////////////////

func myCaculFunc(A, B int) {
	var C int
	fmt.Printf("\n\n°myCalculFunc\n\n")
	fmt.Printf("My first number is: %v, my second is %v, and my total is: %v \n\n", A, B, C)
	/* ou Println pour afficher sans formatage  */
	/* fmt.Println(x, y) */
}

/////////////////////////////////////////////////////////////

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

/////////////////////////////////////////////////////////////

func sayByeFunc(name string) (string, error) {
	if name == "" {
		return "", errors.New("\n\n\bErreur, aucun nom spécifié ! ")
	}

	byeMsg := fmt.Sprintf("\n\n%v s'en va! Bonne soirée", name)
	return byeMsg, nil
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

func myCalendarFunc() {
	myCalendar := map[string]int{
		"JANVIER":   31,
		"FEVRIER":   28,
		"MARS":      31,
		"AVRIL":     30,
		"MAI":       31,
		"JUIN":      30,
		"JUILLET":   31,
		"AOUT":      31,
		"SEPTEMBRE": 30,
		"OCTOBRE":   31,
		"NOVEMBRE":  30,
		"DECEMBRE":  31,
	}
	daysInAYear := 0
	for key, value := range myCalendar {
		fmt.Printf("Mois: %s => %d Jours \n", key, value)
		daysInAYear = daysInAYear + value
	}
	fmt.Printf("\n\nNbr de jours dans l'année: %v jours \n\n", daysInAYear)
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
