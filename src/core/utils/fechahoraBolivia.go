package utils

import (
	"time"
)

func FechaHoraBolivia() time.Time {
	fecha := time.Now()
	return fecha.Add(-4 * time.Hour)

}
