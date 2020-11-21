package main

import (
	"math"
)

type fee struct {
	year                  int
	month                 int
	price                 float64
	priceForTime          float64
	interest              float64
	interestForTime       float64
	amortization          float64
	amortizationFortime   float64
	pendingCapital        float64
	pendingCapitalForTime float64
}

// CalcMortgageAmortization wht
func CalcMortgageAmortization(capitalInput float64, termsInput int, interestTypeInput float64, amortizationAmountInput float64, amortizationYearInput int, amortizationMonthInput int) (float64, float64, int, int, int, float64) {

	numberOfpayments := termsInput * 12
	fees := make([]fee, int(numberOfpayments+1))
	monthlyPayment := Pmt(float64(interestTypeInput), float64(numberOfpayments), capitalInput*-1, 0, 0)
	monthlyInterest := interestTypeInput / 12
	totalToPay := monthlyPayment * float64(numberOfpayments)
	totalInterest := totalToPay - capitalInput
	totalTimeInterest := totalToPay - capitalInput
	monthlyPrice := float64(0)
	pendingPayments := 0

	var accInterest float64

	for i := 0; i <= numberOfpayments; i++ {
		if i == 0 {
			fees[0] = fee{0, 0, 0, 0, 0, 0, 0, 0, capitalInput, capitalInput}
			continue
		}

		var amortization float64
		var amortizationFortime float64
		year := float64((i / 12) + 1)
		previousFee := fees[i-1]
		interest := (previousFee.pendingCapital * monthlyInterest) / 100
		interestForTime := (previousFee.pendingCapitalForTime * monthlyInterest) / 100
		var priceForTime float64
		if (previousFee.pendingCapitalForTime < previousFee.priceForTime) || previousFee.pendingCapitalForTime == 0 {
			priceForTime = previousFee.pendingCapitalForTime + interestForTime
		} else {
			priceForTime = Pmt(interestTypeInput, float64(numberOfpayments), capitalInput, 0, 0) * -1
		}
		price := Pmt(interestTypeInput, float64(int(numberOfpayments)-previousFee.month), previousFee.pendingCapital, 0, 0) * -1
		finalPrice := math.Round(priceForTime*100) / 100
		amortizationRealMonth := int(((amortizationYearInput - 1) * 12) + amortizationMonthInput)

		if i == (amortizationYearInput)*12+amortizationMonthInput+1 {
			monthlyPrice = math.Round(price*100) / 100
		}
		if i == amortizationRealMonth {
			amortization = price - interest + amortizationAmountInput
			amortizationFortime = priceForTime - interestForTime + amortizationAmountInput
		} else {
			amortization = price - interest
			amortizationFortime = priceForTime - interestForTime
		}

		pendingCapital := previousFee.pendingCapital - amortization
		pendingCapitalForTime := previousFee.pendingCapitalForTime - amortizationFortime

		if finalPrice > 0 {
			pendingPayments++
			totalTimeInterest -= interest
		}

		itemFee := fee{
			int(math.Floor(year)),
			i,
			finalPrice,
			math.Round(priceForTime*100) / 100,
			math.Round(interest*100) / 100,
			math.Round(interestForTime*100) / 100,
			math.Round(amortization*100) / 100,
			math.Round(amortizationFortime*100) / 100,
			math.Round(pendingCapital*100) / 100,
			math.Round(pendingCapitalForTime*100) / 100,
		}

		fees[i] = itemFee

		accInterest += interest
	}

	timeSavingsYear := (numberOfpayments - pendingPayments) / 12
	timeSavingsMonth := int(math.Mod(float64(numberOfpayments)-float64(pendingPayments), 12))

	interestSavingsForPrice := math.Round((totalInterest-accInterest)*100) / 100

	return interestSavingsForPrice, monthlyPrice, pendingPayments, timeSavingsYear, timeSavingsMonth, math.Round((totalTimeInterest)*100) / 100
}
