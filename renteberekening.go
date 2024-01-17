package main

import (
	"fmt"
)

func main() {
	startkapitaal := 1000.0
	maandelijksBedrag := 500.0
	basisrente := 0.9 / 100
	getrouwheidspremie := 1.65 / 100
	aantalJaar := 10
	aantalExtraMaanden := 0 // De gebruiker kan dit wijzigen naar wens

	totaalSaldo := startkapitaal
	totaleStortingen := startkapitaal // Begin met het startkapitaal

	// Bereken het totaal aantal maanden
	totaalAantalMaanden := aantalJaar*12 + aantalExtraMaanden

	for maand := 1; maand <= totaalAantalMaanden; maand++ {
		// Voeg maandelijks bedrag toe aan totaalSaldo en totaleStortingen
		totaalSaldo += maandelijksBedrag
		totaleStortingen += maandelijksBedrag

		// Bereken getrouwheidspremie elke 3 maanden vanaf het tweede jaar
		if maand > 12 && maand%3 == 0 {
			totaalSaldo += totaalSaldo * getrouwheidspremie
		}

		// Voeg basisrente toe aan het einde van elk jaar
		if maand%12 == 0 {
			totaleBasisrente := totaalSaldo * basisrente
			totaalSaldo += totaleBasisrente
		}
	}

	// Bereken en print de resultaten
	winst := totaalSaldo - totaleStortingen
	fmt.Printf("Totaal gespaarde bedrag na %d jaar en %d maanden is €%.2f\n", aantalJaar, aantalExtraMaanden, totaalSaldo)
	fmt.Printf("Totaal gestort bedrag na %d jaar en %d maanden is €%.2f\n", aantalJaar, aantalExtraMaanden, totaleStortingen)
	fmt.Printf("Totale winst na %d jaar en %d maanden is €%.2f\n", aantalJaar, aantalExtraMaanden, winst)
}
