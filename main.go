package main

import (
	"fmt"
	"math"
)

func main() {
	f := func(x float64) float64 {
		return math.Exp(-4 * x * x)
	}
	fmt.Println("f(x) = e^(-4x^2)")
	fmt.Printf("trapezoid estimate F(x) 0 -> 1, 4 slices = %v\n", estimate(trapezoid, f, 0, 1, 4))
	fmt.Printf("midpoint estimate F(x) 0 -> 1, 4 slices = %v\n", estimate(midpoint, f, 0, 1, 4))
	fmt.Printf("simpsons estimate F(x) 0 -> 1, 4 slices = %v\n", estimate(simpsons, f, 0, 1, 4))
}

type continuousFunction func(x float64) float64

type method string

const (
	midpoint  method = "midpoint"
	trapezoid method = "trapezoid"
	simpsons  method = "simpsons"
)

func estimate(m method, f continuousFunction, intervalStart, intervalEnd float64, slices int64) float64 {
	switch m {
	case trapezoid:
		return trapezoidEstimate(f, intervalStart, intervalEnd, slices)
	case midpoint:
		return midpointEstimate(f, intervalStart, intervalEnd, slices)
	case simpsons:
		return simpsonsEstimate(f, intervalStart, intervalEnd, slices)
	default:
		return 0.0
	}
}

func trapezoidEstimate(f continuousFunction, intervalStart, intervalEnd float64, slices int64) float64 {
	s := (intervalEnd - intervalStart) / float64(slices)
	value := 0.0

	value += f(intervalStart)

	for x := intervalStart + s; x < intervalEnd; x += s {
		value += 2 * f(x)
	}

	value += f(intervalEnd)

	return value * s / 2
}

func midpointEstimate(f continuousFunction, intervalStart, intervalEnd float64, slices int64) float64 {
	s := (intervalEnd - intervalStart) / float64(slices)
	value := 0.0

	for x := intervalStart + s/2; x < intervalEnd; x += s {
		value += f(x)
	}
	return value * s
}

func simpsonsEstimate(f continuousFunction, intervalStart, intervalEnd float64, slices int64) float64 {
	s := (intervalEnd - intervalStart) / float64(slices)
	value := 0.0

	value += f(intervalStart)

	magicNum := 4.0

	for x := intervalStart + s; x < intervalEnd; x += s {
		value += f(x) * magicNum
		if magicNum == 2 {
			magicNum = 4
		} else {
			magicNum = 2
		}
	}

	value += f(intervalEnd)

	return value * s / 3
}
