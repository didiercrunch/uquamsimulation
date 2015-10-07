package main

import (
	"math"
	"math/rand"
)

var normalBoxMullerAlreadyComputer bool = false
var normalBoxMullerAlreadyComputedValue float64 = 0

func NormalBoxMuller(r *rand.Rand) float64 {
	if normalBoxMullerAlreadyComputer {
		normalBoxMullerAlreadyComputer = false
		return normalBoxMullerAlreadyComputedValue
	}
	radius := math.Sqrt(-2 * math.Log(r.Float64()))
	theta := 2 * math.Pi * r.Float64()

	normalBoxMullerAlreadyComputer = true
	normalBoxMullerAlreadyComputedValue = radius * math.Cos(theta)
	return radius * math.Cos(theta)
}
