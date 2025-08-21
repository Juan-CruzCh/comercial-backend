package utils

import (
	"time"
)

func FechaHoraBolivia() (time.Time, error) {
	loc, err := time.LoadLocation("America/La_Paz")
	if err != nil {

		return time.Time{}, err
	}
	now := time.Now().In(loc)
	return now, nil

}
