package estimate

import "fmt"

type ContinuousFunction func(x float64) float64

type method string

const (
	Midpoint  method = "midpoint"
	Trapezoid method = "trapezoid"
	Simpsons  method = "simpsons"
)

func EstimateString(m method, f ContinuousFunction, intervalStart, intervalEnd float64, slices int64) string {
	value := Estimate(m, f, intervalStart, intervalEnd, slices)
	return fmt.Sprintf("%v estimate of F(x) from %v -> %v, %v slices: %v", m, intervalStart, intervalEnd, slices, value)
}

func Estimate(m method, f ContinuousFunction, intervalStart, intervalEnd float64, slices int64) float64 {
	switch m {
	case Trapezoid:
		return trapezoidEstimate(f, intervalStart, intervalEnd, slices)
	case Midpoint:
		return midpointEstimate(f, intervalStart, intervalEnd, slices)
	case Simpsons:
		return simpsonsEstimate(f, intervalStart, intervalEnd, slices)
	default:
		return 0.0
	}
}

func trapezoidEstimate(f ContinuousFunction, intervalStart, intervalEnd float64, slices int64) float64 {
	s := (intervalEnd - intervalStart) / float64(slices)
	value := 0.0

	value += f(intervalStart)

	for x := intervalStart + s; x < intervalEnd; x += s {
		value += 2 * f(x)
	}

	value += f(intervalEnd)

	return value * s / 2
}

func midpointEstimate(f ContinuousFunction, intervalStart, intervalEnd float64, slices int64) float64 {
	s := (intervalEnd - intervalStart) / float64(slices)
	value := 0.0

	for x := intervalStart + s/2; x < intervalEnd; x += s {
		value += f(x)
	}
	return value * s
}

func simpsonsEstimate(f ContinuousFunction, intervalStart, intervalEnd float64, slices int64) float64 {
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
