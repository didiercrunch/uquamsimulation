package main

import (
	"math/rand"
	"testing"

	"github.com/bszcz/mt19937_64"
)

func BenchmarkMersenneTwister(b *testing.B) {
	r := mt19937_64.New()
	for i := 0; i < b.N; i++ {
		r.Real1()
	}
}

func BenchmarkStdUniforme01(b *testing.B) {
	r := rand.New(rand.NewSource(99))
	for i := 0; i < b.N; i++ {
		r.Float64()
	}
}

func BenchmarkAESRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aesRandom()
	}
}

func BenchmarkStdExpoVariate(b *testing.B) {
	r := rand.New(rand.NewSource(99))
	for i := 0; i < b.N; i++ {
		r.ExpFloat64()
	}
}

func BenchmarkExpoVariate(b *testing.B) {
	r := rand.New(rand.NewSource(99))
	for i := 0; i < b.N; i++ {
		ExpFloat64(r)
	}
}

func BenchmarkNormalBoxMuller(b *testing.B) {
	r := rand.New(rand.NewSource(99))
	for i := 0; i < b.N; i++ {
		NormalBoxMuller(r)
	}
}

func BenchmarkStdNormal(b *testing.B) {
	r := rand.New(rand.NewSource(99))
	for i := 0; i < b.N; i++ {
		r.NormFloat64()
	}
}

func BenchmarkNormalMarsaglia(b *testing.B) {
	r := rand.New(rand.NewSource(99))
	for i := 0; i < b.N; i++ {
		NormalMarsaglia(r)
	}
}
