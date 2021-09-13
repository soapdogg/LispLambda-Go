package impl

import "lisp_lambda-go/internal/function/impl/internal"

type gcdCalculator struct {}

func NewGcdCalculator() *gcdCalculator {
	return &gcdCalculator{}
}

func (g gcdCalculator) CalculateGCD(
	i int,
	j int,
) int {
	x := i
	y := j
	for x != 0 && y != 0 {
		if x > y {
			x %= y
		} else {
			y %= x
		}
	}
	if x == 0 {
		x = y
	}
	return x
}

var _ internal.GCDCalculator = &gcdCalculator{}