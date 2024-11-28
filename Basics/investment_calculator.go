package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmetAmount float64
	var years float64
	expectedRerturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmetAmount) // & pointer to investmentAmount

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedRerturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmetAmount, expectedRerturnRate, years)
	// futureValue := investmetAmount * math.Pow(1+expectedRerturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	//formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)
	//fmt.Println("Future Value: ", futureValue)
	//fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)
	fmt.Printf("Future Value: %f\nFuture Value (adjusted for inflation): %f", futureValue, futureRealValue)

	// fmt.Printf(`Future Value: %f\n
	// Future Value (adjusted for inflation): %f`, futureValue, futureRealValue)

	//fmt.Print(formattedFV, formattedRFV)

}

func calculateFutureValue(investmetAmount, expectedRerturnRate, years float64) (fv float64, rfv float64) {
	fv = investmetAmount * math.Pow(1+expectedRerturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
