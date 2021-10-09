package field

import "egosystem/space"

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

func (g *SquareGrid) Value(point space.Point) float32 {
	return float32(g.nodes[0][0]) // lol
}
