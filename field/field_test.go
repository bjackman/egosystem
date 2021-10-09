package field

import (
	"egosystem/space"
	"testing"
)

func TestSquareGridUniform(t *testing.T) {
	grid := NewSquareGrid(10, space.Meter)
	for _, value := range []int8{-100, 0, 1} {
		grid.SetUniform(value)
		for _, point := range []space.Point{
			// Corner
			{X: 0, Y: 0},
			// Edge node
			{X: space.Meter, Y: 0},
			// Edge
			{X: space.Centimeter, Y: 0},
			// Node
			{X: space.Meter, Y: space.Meter},
			// Not node
			{X: space.Centimeter, Y: space.Centimeter},
		} {
			got := grid.Value(point)
			if got != float32(value) {
				t.Errorf("After SetUniform(%v), Value(%v) returned %v, want %v",
					value, point, got, value)
			}
		}
	}
}
