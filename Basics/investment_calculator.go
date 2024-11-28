package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmetAmount float64
	years := 10.0
	expectedRerturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmetAmount) // & pointer to investmentAmount

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedRerturnRate)

	futureValue := investmetAmount * math.Pow(1+expectedRerturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//fmt.Println("Future Value: ", futureValue)
	//fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)
	fmt.Printf("Future Value: %f\nFuture Value (adjusted for inflation): %f", futureValue, futureRealValue)

}
