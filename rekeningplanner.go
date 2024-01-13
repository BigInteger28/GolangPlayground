package main

import (
    "fmt"
)

func main() {
	var monthlyIncome, otherAnnualIncome, fixedMonthlyCosts, plannedAnnualExpenses, annualSavings float64


        //Maandelijks inkomen: ")
        monthlyIncome=2140

        //Extra jaarlijks inkomen: ")
        otherAnnualIncome=4000

        //Vaste verplichte maandelijkse kosten: ")
        fixedMonthlyCosts=400

        //Geplande jaarlijkse uitgaven (b.v., reizen, geplande aankopen): ")
        plannedAnnualExpenses=11000

        //Jaarlijks doel om te sparen: ")
        annualSavings=10000

        // Berekeningen
        annualSavingsRequired := annualSavings - otherAnnualIncome
        if annualSavingsRequired < 0 {
            plannedAnnualExpenses += -annualSavingsRequired
            annualSavingsRequired = 0
        }
        totalAnnualIncome := (monthlyIncome * 12)
        totalPlannedAnnualExpenses := plannedAnnualExpenses
        totalAnnualSavings := annualSavingsRequired
        totalExpenses := (fixedMonthlyCosts * 12) + totalPlannedAnnualExpenses
        monthlyFunBudget := (totalAnnualIncome - totalExpenses - totalAnnualSavings) / 12

        // Output
        fmt.Printf("\nMaandelijks naar Fun Rekening: €%.2f\n", monthlyFunBudget)
        fmt.Printf("Maandelijks laten staan op Zichtrekening (Vaste verplichte kosten): €%.2f\n", fixedMonthlyCosts)
        fmt.Printf("Maandelijks naar Geplande Spaarrekening: €%.2f\n", plannedAnnualExpenses/12)
        fmt.Printf("Maandelijks naar Lange Termijn Spaarrekening: €%.2f\n", totalAnnualSavings/12)
        fmt.Printf("Jaarlijkse bedrag extra sparen op Lange Termijn Spaarrekening: €%.2f\n", totalAnnualSavings)
}
