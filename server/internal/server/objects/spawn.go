package objects

import "math/rand/v2"

func SpawnCoords() (float64, float64) {
	var bound float64 = 3000
	return rand.Float64() * bound, rand.Float64() * bound
}
