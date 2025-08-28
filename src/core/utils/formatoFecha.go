package utils

import "time"

func NormalizarRangoDeFechas(fechaInicio time.Time, fechaFin time.Time) (f1 time.Time, f2 time.Time) {
	f1 = time.Date(fechaInicio.Year(), fechaInicio.Month(), fechaInicio.Day(), 0, 0, 0, 0, fechaInicio.Location())
	f2 = time.Date(fechaFin.Year(), fechaFin.Month(), fechaFin.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), fechaFin.Location())
	return f1, f2
}
