package Calculations

import (
	"math"
	"math/rand"
	"time"
)

// EucliDistance calculates the Euclidean distance between two points
func EucliDistance(x, y [2]float64) float64 {
	distance := math.Sqrt(math.Pow(x[0]-y[0], 2) + math.Pow(x[1]-y[1], 2))
	return distance
}

// getRandomCentroid generates a random centroid (point) within a 2x2 range
func GetRandomCentroid() [2]float64 {

	rand.Seed(time.Now().UnixNano())
	var centroid [2]float64
	centroid[0] = 0 + rand.Float64()*(2-0)
	centroid[1] = 0 + rand.Float64()*(2-0)
	return centroid
}
