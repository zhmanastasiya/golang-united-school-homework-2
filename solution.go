package solution

import (
	"math"
)

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

type myInt int

func CalcSquare(sideLen float64, sidesNum myInt) float64 {
	if sidesNum == SidesTriangle {
		x := (sideLen * 3) / 2
		s := math.Sqrt(x * (x - sideLen) * (x - sideLen) * (x - sideLen))
		return s
	} else if sidesNum == SidesSquare {
		s := sideLen * sideLen
		return s
	} else if sidesNum == SidesCircle {
		s := math.Pi * (sideLen * sideLen)
		return s
	} else {
		return 0
	}
}
