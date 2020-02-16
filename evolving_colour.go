package main

import (
	"math/rand"
)

type EvolvingColour struct {
	value     int
	ascOrDesc string
}

func NewEvolvingColour() *EvolvingColour {
	eColour := new(EvolvingColour)
	eColour.value = rand.Intn(256)
	if rand.Float32() < 0.5 {
		eColour.ascOrDesc = "asc"
	} else {
		eColour.ascOrDesc = "desc"
	}
	return eColour
}

func adjustColourValue(eColour EvolvingColour) EvolvingColour {
	switch eColour.ascOrDesc {
	case "asc":
		if eColour.value > 254 {
			eColour.ascOrDesc = "desc"
		} else {
			eColour.value += 1
		}
	case "desc":
		if eColour.value < 1 {
			eColour.ascOrDesc = "asc"
		} else {
			eColour.value -= 1
		}
	}

	return eColour
}
