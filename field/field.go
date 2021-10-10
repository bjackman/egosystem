package field

import (
	"egosystem/space"
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

// Given two nodes unit distance apart with values z0 and z1, get the value at d
// along the line between those nodes.
func linearInterpUnit(z0, z1, d float32) float32 {
	return float32(z0) + d*float32(z1-z0)
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
	xOob := point.X < 0 || point.X >= g.dimension()
	yOob := point.Y < 0 || point.Y >= g.dimension()

	res := float32(g.resolution.Centimeters())

	// Handle if point is out of bounds on one or both axes.
	if xOob && yOob {
		// Use value from closest corner node.
		var nodeX, nodeY int
		if point.X > 0 {
			nodeX = g.size - 1
		}
		if point.Y > 0 {
			nodeY = g.size - 1
		}
		return float32(g.nodes[nodeX][nodeY])
	} else if xOob {
		// Do linear interpolation between two nodes on the Y edge.
		origin := index2D{
			y: point.Y.Div(g.resolution),
		}
		if point.X < 0 {
			origin.x = 0 // Originwards y-edge
		} else {
			origin.x = g.size - 1 // Other y-edge
		}
		return linearInterpUnit(
			float32(g.nodes[origin.x][origin.y]),
			float32(g.nodes[origin.x][origin.y+1]),
			float32(point.Y%g.resolution)/res,
		)
	} else if yOob {
		// Same as above but with x/y swapped.
		origin := index2D{
			x: point.X.Div(g.resolution),
		}
		if point.Y < 0 {
			origin.y = 0
		} else {
			origin.y = g.size - 1
		}
		return linearInterpUnit(
			float32(g.nodes[origin.x][origin.y]),
			float32(g.nodes[origin.x+1][origin.y]),
			float32(point.X%g.resolution)/res,
		)
	}

	// Index of originwards node of the square that point is in
	origin := index2D{
		x: point.X.Div(g.resolution),
		y: point.Y.Div(g.resolution),
	}

	// Get position of point within the square, treating the square's size as unit.
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

func (g *SquareGrid) dimension() space.Distance {
	return g.resolution.MulInt(g.size - 1)
}
