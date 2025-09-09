package utils

import "math"

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Porcentaje(valor float64) float64 {
	if valor <= 0 {
		return 0
	}
	return (valor / 100) * 100
}
