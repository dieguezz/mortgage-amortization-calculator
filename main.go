package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func startServer() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5000")
		c.Header("Access-Control-Allow-Credentials", "true")

		capital, err := strconv.ParseFloat(c.PostForm("capital"), 64)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}
		terms, err := strconv.Atoi(c.PostForm("terms"))
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}
		interestType, err := strconv.ParseFloat(c.PostForm("interestType"), 64)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}
		amortizationAmount, err := strconv.ParseFloat(c.PostForm("amortizationAmount"), 64)
		if err != nil {
			amortizationAmount = 0
		}
		year, err := strconv.Atoi(c.PostForm("year"))
		if err != nil {
			year = 0
		}
		month, err := strconv.Atoi(c.PostForm("month"))
		if err != nil {
			month = 0
		}

		interestSavingsForPrice, monthlyPrice, pendingPayments, timeSavingsYear, timeSavingsMonth, totalTimeInterest, fees := CalcMortgageAmortization(capital, terms, interestType, amortizationAmount, year, month)
		// fmt.Println(fees)

		// parsedFees, _ := json.Marshal(&fees)
		// fmt.Println(parsedFees)
		result := map[string]interface{}{
			"interestSavingsForPrice": interestSavingsForPrice,
			"monthlyPrice":            monthlyPrice,
			"pendingPayments":         pendingPayments,
			"timeSavingsYear":         timeSavingsYear,
			"timeSavingsMonth":        timeSavingsMonth,
			"totalTimeInterest":       totalTimeInterest,
			"fees":                    &fees,
		}
		c.JSON(http.StatusOK, gin.H(result))
	})

	r.Run()
}

func main() {
	// capitalInput := float64(50000)
	// termsInput := 40
	// interestTypeInput := float64(5.10)
	// amortizationAmountInput := float64(10000)
	// interestSavingsForPrice, monthlyPrice, pendingPayments, timeSavingsYear, timeSavingsMonth, totalTimeInterest := CalcMortgageAmortization(capitalInput, termsInput, interestTypeInput, amortizationAmountInput, 4, 2)
	// fmt.Println("===================================")
	// fmt.Println("💰 Para reducir cuota")
	// fmt.Println("===================================")
	// fmt.Println("- Cuota mensual: ", monthlyPrice, "€")
	// fmt.Println("- Ahorro en intereses: ", interestSavingsForPrice, "€")
	// fmt.Println("")
	// fmt.Println("===================================")
	// fmt.Println("⏰ For time amortization")
	// fmt.Println("===================================")
	// fmt.Println("- Cuotas pendientes: ", pendingPayments)
	// fmt.Println("- Ahorro en tiempo: ", timeSavingsYear, "años y", timeSavingsMonth, "meses")
	// fmt.Println("- Ahorro en intereses: ", totalTimeInterest, "€")
	// fmt.Println("")

	startServer()
}
