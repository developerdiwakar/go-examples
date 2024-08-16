package main

import (
	"fmt"
	"math"
)

// CompoundInterest calculates the compound interest using the formula A = P(1 + r/n)^(nt)
func CompoundInterest(principal float64, rate float64, timesCompounded int, years float64) float64 {
	// Convert interest rate to decimal
	rateDecimal := rate / 100

	// Calculate the compound interest
	amount := principal * math.Pow(1+(rateDecimal/float64(timesCompounded)), float64(timesCompounded)*years)
	return amount
}

func main() {
	// Calculate Compound Interest
	// CI = P(1+r/n)^t - P
	// 	This formula is also called periodic compounding formula.

	// Here,

	// A represents the new principal sum or the total amount of money after compounding period
	// P represents the original amount or initial amount
	// r is the annual interest rate
	// n represents the compounding frequency or the number of times interest is compounded in a  year
	// t represents the number of years
	// Input values
	principal := 250000.0 // Principal amount (P)
	rate := 6.67          // Annual interest rate (r) in percentage
	timesCompounded := 1  // Number of times interest is compounded per year (n)
	years := 9.0          // Number of years the money is invested for (t)

	// Calculate compound interest
	finalAmount := CompoundInterest(principal, rate, timesCompounded, years)
	compoundInterest := finalAmount - principal

	// Output results
	fmt.Printf("Initial Investment: Rs.%.2f\n", principal)
	fmt.Printf("Annual Interest Rate: %.2f%%\n", rate)
	fmt.Printf("Times Compounded per Year: %d\n", timesCompounded)
	fmt.Printf("Investment Period: %.2f years\n", years)
	fmt.Printf("Final Amount: Rs.%.2f\n", finalAmount)
	fmt.Printf("Compound Interest Earned: Rs.%.2f\n", compoundInterest)

}
