package field

import (
	"egosystem/space"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

func TestSquareGridInterpolation(t *testing.T) {
	grid := SquareGrid{
		// Warning: X axis is veritical in this literal!
		nodes: [][]int8{
			{0, 100, 50},
			{-70, 100, -100},
			{90, -50, 10},
		},
		size:       3,
		resolution: space.Meter,
	}
	for point, want := range map[space.Point]float32{
		// Corner node
		{}: 0.0,
		// Edge node
		{X: 1 * space.Meter, Y: 0}: -70.0,
		// Originwards edge of grid
		{X: 50 * space.Centimeter, Y: 0}:  -35.0,
		{X: 0, Y: 125 * space.Centimeter}: 87.5,
		// Non-originwards edge of grid
		{X: 2 * space.Meter, Y: 10 * space.Centimeter}: 76.0,
		{X: 50 * space.Centimeter, Y: 2 * space.Meter}: -25.0,
		// OOB in one dimension
		{X: 3 * space.Meter, Y: 10 * space.Centimeter}: 76.0,
		{X: 50 * space.Centimeter, Y: 3 * space.Meter}: -25.0,
		// OOB in both dimensions
		{X: -1 * space.Kilometer, Y: -1 * space.Kilometer}: 0.0,
		{X: -1 * space.Kilometer, Y: +1 * space.Kilometer}: 50.0,
		{X: +1 * space.Kilometer, Y: -1 * space.Kilometer}: 90.0,
		{X: +1 * space.Kilometer, Y: +1 * space.Kilometer}: 10.0,
		// Node
		{X: 1 * space.Meter, Y: 1 * space.Meter}: 100.0,
		// Edge
		{X: 1 * space.Meter, Y: 170 * space.Centimeter}: -40.0,
		// In the middle
		// Calculated by
		// https://www.ajdesigner.com/phpinterpolation/bilinear_interpolation_equation.php
		{X: 145 * space.Centimeter, Y: 163 * space.Centimeter}: -19.79,
	} {
		got := grid.Value(point)
		if diff := cmp.Diff(got, want, cmpopts.EquateApprox(0, 0.000001)); diff != "" {
			t.Errorf("Value(%v) returned %v, want %v", point, got, want)
		}
	}
}
