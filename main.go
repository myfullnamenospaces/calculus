package main

import (
	"fmt"
	"math"

	"github.com/myfullnamenospaces/calculus/pkg/estimate"
)

func main() {
	f := func(x float64) float64 {
		return 2 * math.Exp(-x/4)
	}
	fmt.Println("f(x) = 2e^(-4/x)")
	fmt.Println(estimate.EstimateString(estimate.Trapezoid, f, 0, 8, 20))
	fmt.Println(estimate.EstimateString(estimate.Midpoint, f, 0, 8, 20))
	fmt.Println(estimate.EstimateString(estimate.Simpsons, f, 0, 8, 20))
	fmt.Println(estimate.ErrorString(f, 0, 8, 20))
}
