package estimate

import (
	"fmt"
	"math"
)

func ErrorString(f ContinuousFunction, intervalStart, intervalEnd float64, slices int64) string {
	value := SimpsonsError(f, intervalStart, intervalEnd, slices)
	return fmt.Sprintf("error estimate of F(x) %v -> %v, %v slices: %v", intervalStart, intervalEnd, slices, value)
}

func SimpsonsError(f ContinuousFunction, intervalStart, intervalEnd float64, slices int64) float64 {
	s := (intervalEnd - intervalStart) / float64(slices)
	max := float64(0)

	for x := intervalStart; x < intervalEnd; x += s {
		value := f(x)
		if math.Abs(value) > max {
			max = math.Abs(value)
		}
	}

	return max * math.Pow(intervalEnd-intervalStart, 5) / (180 * math.Pow(float64(slices), 4))
}
