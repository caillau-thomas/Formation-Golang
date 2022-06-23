package main

import "fmt"

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
