package main

import (
	"egosystem/field"
	"egosystem/space"
)

type ground struct {
	moisture field.SquareGrid
}

func newGround(size space.Distance, gridResolution space.Distance) *ground {
	gridSize := size.Div(gridResolution)
	return &ground{moisture: field.NewSquareGrid(gridSize, gridResolution)}
}

func main() {
}
