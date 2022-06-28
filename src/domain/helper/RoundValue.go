package domain

import "math"

func Round(value float64) float64 {
	ratio := math.Pow(10, float64(2))
	return math.Round(value*ratio) / ratio
}
