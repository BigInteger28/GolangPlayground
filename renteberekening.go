package main

import (
	"fmt"
)

func main() {
	startkapitaal := 2000.0
	maandelijksBedrag := 500.0
	basisrente := 0.9 / 100
	getrouwheidspremie := 1.65 / 100
	aantalJaar := 5 // Aantal jaren voor renteberekening

	totaalSaldo := float64(startkapitaal) // Begin met het startkapitaal

	for jaar := 1; jaar <= aantalJaar; jaar++ {
		for maand := 1; maand <= 12; maand++ {
			// Voeg maandelijks bedrag toe aan totaalSaldo
			totaalSaldo += float64(maandelijksBedrag)

			if jaar > 1 && maand%3 == 0 {
				// Voeg getrouwheidspremie toe elke 3 maanden vanaf het tweede jaar
				totaalSaldo += totaalSaldo * getrouwheidspremie
			}
		}

		// Bereken en voeg basisrente toe aan het einde van elk jaar
		totaleBasisrente := totaalSaldo * basisrente
		totaalSaldo += totaleBasisrente

		fmt.Printf("Jaar %d: Totale saldo is €%.2f\n", jaar, totaalSaldo)
	}
}
