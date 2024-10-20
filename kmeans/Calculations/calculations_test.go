package Calculations

import (
	"math/rand"
	"testing"
)

// BenchmarkEucliDistance benchmarks the EucliDistance function
func BenchmarkEucliDistance(b *testing.B) {
	// Generate random test points for benchmarking
	var point1, point2 [2]float64
	for i := 0; i < 2; i++ {
		point1[i] = rand.Float64() * 100.0
		point2[i] = rand.Float64() * 100.0
	}

	// Run the benchmark
	b.ResetTimer() // Reset timer to ignore setup time
	for i := 0; i < b.N; i++ {
		EucliDistance(point1, point2)
	}
}