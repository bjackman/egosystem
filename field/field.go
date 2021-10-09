package field

import (
	"egosystem/space"
	"log"
)

type SquareGrid struct {
	nodes      [][]int8
	size       int
	resolution space.Distance
}

func NewSquareGrid(size int, resolution space.Distance) SquareGrid {
	g := SquareGrid{size: size, resolution: resolution}
	g.nodes = make([][]int8, size)
	for i := range g.nodes {
		g.nodes[i] = make([]int8, size)
	}
	return g
}

func (g *SquareGrid) SetUniform(value int8) {
	for x := 0; x < g.size; x++ {
		for y := 0; y < g.size; y++ {
			g.nodes[x][y] = value
		}
	}
}

type index2D struct{ x, y int }

func (i index2D) InSquare(size int) bool {
	return i.x >= 0 && i.x < size && i.y >= 0 && i.y < size
}

// Given a unit square field with nodes zxy at each of the four corners, do
// bilinear interpolation to get the value at x,y inside that square.
// https://blogs.sas.com/content/iml/2020/05/18/what-is-bilinear-interpolation.html
func bilinearInterpUnit(z00, z01, z10, z11 int8, x, y float32) float32 {
	return float32(z00)*(1-x)*(1-y) +
		float32(z10)*x*(1-y) +
		float32(z01)*(1-x)*y +
		float32(z11)*x*y
}

func (g *SquareGrid) Value(point space.Point) float32 {
	if g.resolution == space.Distance(0) {
		log.Panic("what")
	}

	// Index of originwards node of the square that point is in
	origin := index2D{
		x: point.X.Div(g.resolution),
		y: point.Y.Div(g.resolution),
	}
	if !origin.InSquare(g.size) {
		// TODO: Handle this, log non-fatally but only if it's outside by a non-negligible margin
		log.Fatalf("Point %v not within square grid of size %v@%v", point, g.size, g.resolution)
	}

	// Get position of point within the square, treating the square's size as unit.
	res := float32(g.resolution.Centimeters())
	squareX := float32(point.X%g.resolution) / res
	squareY := float32(point.Y%g.resolution) / res

	return bilinearInterpUnit(
		g.nodes[origin.x+0][origin.y+0],
		g.nodes[origin.x+0][origin.y+1],
		g.nodes[origin.x+1][origin.y+0],
		g.nodes[origin.x+1][origin.y+1],
		squareX, squareY,
	)
}
