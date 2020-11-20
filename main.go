package main

import (
	"fmt"
)

func main() {
	interestSavingsForPrice, monthlyPrice, pendingPayments, timeSavingsYear, timeSavingsMonth, totalTimeInterest := CalcMortgageAmortization(4, 2)
	fmt.Println("===================================")
	fmt.Println("💰 Para reducir cuota")
	fmt.Println("===================================")
	fmt.Println("- Cuota mensual: ", monthlyPrice, "€")
	fmt.Println("- Ahorro en intereses: ", interestSavingsForPrice, "€")
	fmt.Println("")
	fmt.Println("===================================")
	fmt.Println("⏰ For time amortization")
	fmt.Println("===================================")
	fmt.Println("- Cuotas pendientes: ", pendingPayments)
	fmt.Println("- Ahorro en tiempo: ", timeSavingsYear, "años y", timeSavingsMonth, "meses")
	fmt.Println("- Ahorro en intereses: ", totalTimeInterest, "€")
	fmt.Println("")
}
