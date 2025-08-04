package bench

import (
	"math"
	"math/rand"
)

func RangeFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RoundToDigits(f float32, howMany float64) float32 {
	return float32(math.Round(float64(f)*howMany) / howMany)
}

func Chance(percent float64) bool {
	return rand.Float64() > percent
}
