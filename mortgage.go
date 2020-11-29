package main

import (
	"encoding/json"
	"log"
	"math"
)

// Fee what
type Fee struct {
	Year                  int     `json:"year"`
	Month                 int     `json:"month"`
	Price                 float64 `json:"price"`
	PriceForTime          float64 `json:"priceForTime"`
	Interest              float64 `json:"interest"`
	InterestForTime       float64 `json:"interestForTime"`
	Amortization          float64 `json:"amortization"`
	AmortizationFortime   float64 `json:"amortizationFortime"`
	PendingCapital        float64 `json:"pendingCapital"`
	PendingCapitalForTime float64 `json:"pendingCapitalForTime"`
}

// CalcMortgageAmortization wht
func CalcMortgageAmortization(capitalInput float64, termsInput int, interestTypeInput float64, amortizationAmountInput float64, amortizationYearInput int, amortizationMonthInput int) (float64, float64, int, int, int, float64, []Fee) {

	numberOfpayments := termsInput * 12
	fees := make([]Fee, int(numberOfpayments+1))
	monthlyPayment := Pmt(float64(interestTypeInput), float64(numberOfpayments), capitalInput*-1, 0, 0)
	monthlyInterest := interestTypeInput / 12
	totalToPay := monthlyPayment * float64(numberOfpayments)
	totalInterest := totalToPay - capitalInput
	totalTimeInterest := totalToPay - capitalInput
	monthlyPrice := float64(0)
	pendingPayments := 0

	var accInterest float64
	jsonFees := make([]Fee, int(numberOfpayments+1))
	for i := 0; i <= numberOfpayments; i++ {
		if i == 0 {
			fees[0] = Fee{0, 0, 0, 0, 0, 0, 0, 0, capitalInput, capitalInput}
			continue
		}

		var amortization float64
		var amortizationFortime float64
		year := float64((i / 12) + 1)
		previousFee := fees[i-1]
		interest := (previousFee.PendingCapital * monthlyInterest) / 100
		interestForTime := (previousFee.PendingCapitalForTime * monthlyInterest) / 100
		var priceForTime float64
		if (previousFee.PendingCapitalForTime < previousFee.PriceForTime) || previousFee.PendingCapitalForTime == 0 {
			priceForTime = previousFee.PendingCapitalForTime + interestForTime
		} else {
			priceForTime = Pmt(interestTypeInput, float64(numberOfpayments), capitalInput, 0, 0) * -1
		}
		price := Pmt(interestTypeInput, float64(int(numberOfpayments)-previousFee.Month), previousFee.PendingCapital, 0, 0) * -1
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

		pendingCapital := previousFee.PendingCapital - amortization
		pendingCapitalForTime := previousFee.PendingCapitalForTime - amortizationFortime

		if finalPrice > 0 {
			pendingPayments++
			totalTimeInterest -= interest
		}

		itemFee := Fee{
			Year:                  int(math.Floor(year)),
			Month:                 i,
			Price:                 finalPrice,
			PriceForTime:          math.Round(priceForTime*100) / 100,
			Interest:              math.Round(interest*100) / 100,
			InterestForTime:       math.Round(interestForTime*100) / 100,
			Amortization:          math.Round(amortization*100) / 100,
			AmortizationFortime:   math.Round(amortizationFortime*100) / 100,
			PendingCapital:        math.Round(pendingCapital*100) / 100,
			PendingCapitalForTime: math.Round(pendingCapitalForTime*100) / 100,
		}

		bytes, err := json.Marshal(&itemFee)

		if err != nil {
			log.Fatal(err)
		} else {
			var p Fee
			err = json.Unmarshal(bytes, &p)
			if err != nil {
				panic(err)
			}
			jsonFees[i] = p
		}

		fees[i] = itemFee
		accInterest += interest
	}

	timeSavingsYear := (numberOfpayments - pendingPayments) / 12
	timeSavingsMonth := int(math.Mod(float64(numberOfpayments)-float64(pendingPayments), 12))

	interestSavingsForPrice := math.Round((totalInterest-accInterest)*100) / 100

	return interestSavingsForPrice, monthlyPrice, pendingPayments, timeSavingsYear, timeSavingsMonth, math.Round((totalTimeInterest)*100) / 100, jsonFees
}
