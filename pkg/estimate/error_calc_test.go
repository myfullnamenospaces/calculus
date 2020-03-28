package estimate

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestErrorCalculation(t *testing.T) {
	g := NewGomegaWithT(t)

	f := func(x float64) float64 {
		return x * x
	}

	e := SimpsonsError(f, 0, 2, 10)
	g.Expect(e).To(BeNumerically("~", 7.11111e-5, 0.00001))
}
