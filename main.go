package main

func main() {

}

type continuousFunction func(x float64) float64

type method string

const (
	midpoint  method = "midpoint"
	trapezoid method = "trapezoid"
	other     method = "other"
)

func estimate(m method, f continuousFunction, intervalStart, intervalEnd float64, slices int64) float64 {
	switch m {
	case midpoint:
		return midpointEstimate(f, intervalStart, intervalEnd, slices)
	default:
		return 0.0
	}
}

func midpointEstimate(f continuousFunction, intervalStart, intervalEnd float64, slices int64) float64 {
	s := (intervalEnd - intervalStart) / float64(slices)
	value := 0.0

	for x := s/2 + intervalStart; x < intervalEnd; x += s {
		value += f(x) * s
	}
	return value
}
