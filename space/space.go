package space

import "fmt"

type Distance int

var (
	Centimeter = Distance(1)
	Meter      = 100 * Centimeter
	Kilometer  = 1000 * Meter
)

func (d Distance) Centimeters() int {
	return int(d)
}

func (d Distance) Meters() int {
	return int(d / Meter)
}

func (d Distance) Kilometers() int {
	return int(d / Kilometer)
}

func (d Distance) Div(denominator Distance) int {
	return int(d / denominator)
}

func (d Distance) MulInt(i int) Distance {
	return d * Distance(i)
}

func (d Distance) String() string {
	switch {
	case d < Meter:
		return fmt.Sprintf("%dcm", d)
	case d < Kilometer:
		if d%Meter == 0 {
			return fmt.Sprintf("%dm", d/Meter)
		} else {
			return fmt.Sprintf("%d.%02dm", d/Meter, d%Meter)
		}
	case d%Kilometer == 0:
		return fmt.Sprintf("%dkm", d/Kilometer)
	case d%Meter == 0:
		return fmt.Sprintf("%d.%03dkm", d/Kilometer, (d%Kilometer)/Meter)
	default:
		return fmt.Sprintf("%d.%dm", d/Meter, d%Meter)
	}
}

type Point struct {
	X, Y Distance
}
