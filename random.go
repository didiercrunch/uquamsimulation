package main

import (
	"math"
	"math/rand"
)

func ExpFloat64(r *rand.Rand) float64 {
	return -math.Log(r.Float64())
}

func main() {

}
