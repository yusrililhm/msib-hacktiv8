package entity

import "gorm.io/gorm"

type Weather struct {
	gorm.Model
	Water uint
	Wind uint
}

func (t *Weather) StatusWater() (status string) {
	if t.Water <= 5 {
		status = "aman"
		return
	} else if t.Water >= 6 && t.Water <= 8 {
		status = "siaga"
		return
	} else {
		status = "bahaya"
		return
	}
}

func (t *Weather) StatusWind() (status string) {
	if t.Wind <= 6 {
		status = "aman"
		return
	} else if t.Wind >= 7 && t.Wind <= 15 {
		status = "siaga"
		return
	} else {
		status = "bahaya"
		return
	}
}
