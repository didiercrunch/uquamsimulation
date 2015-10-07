package main

import (
	"math"
	"math/rand"
)

var normalMarsagliaAlreadyComputer bool = false
var normalMarsagliaAlreadyComputedValue float64 = 0

func NormalMarsaglia(r *rand.Rand) float64 {
	if normalMarsagliaAlreadyComputer {
		normalMarsagliaAlreadyComputer = false
		return normalMarsagliaAlreadyComputedValue
	}
	x, y := getPointInUnitDisc(r)
	radius := math.Sqrt(x*x + y*y)

	normalMarsagliaAlreadyComputer = true
	coeff := math.Sqrt(-4 * math.Log(radius) / radius)
	normalMarsagliaAlreadyComputedValue = x * coeff
	return y * coeff
}

func getPointInUnitDisc(r *rand.Rand) (float64, float64) {
	x := 2 * (r.Float64() - 0.5)
	y := 2 * (r.Float64() - 0.5)
	if x*x+y*y < 1 {
		return x, y
	}
	return getPointInUnitDisc(r)
}
