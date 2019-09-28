package gol

import (
	"errors"
	"fmt"
)

// 2 dimensional grid of cells, infinite (ie wrapping)
// each cell is in one of two states; alive or dead
// each cell interacts with it's neighbors (8), n, ne, w, sw, s, se, e, ne

// each tick:
// any live cell with fewer than 2 live neighbours dies
// any live cell with 2 or 3 live neighbours does nothing
// any live cell with more than 3 live neighbours dies

// any dead cell with exactly three live neighbours becomes a live cell

// Alive = true, Dead = false

type Simulation struct {
	xAxisSize int
	yAxisSize int
	cells [][]bool
}

// Origin is top left corner
// x extends horizontally
// y extends vertically
// cells[y][x]
func NewSimulation(xSize int, ySize int) *Simulation {
	cells := make([][]bool, ySize)
	for row := range cells {
		cells[row] = make([]bool, xSize)
	}
	return &Simulation{
		xAxisSize: xSize,
		yAxisSize: ySize,
		cells: cells,
	}
}

func (s *Simulation) Tick() [][]bool {
	for row := range s.cells {
		for col := range s.cells[row] {
			state := s.cells[row][col]
			s.tickCell(row, col, state)
		}
	}

	return s.cells
}

func (s *Simulation) ToggleCell(x int, y int) error {
	if x >= s.xAxisSize {
		return errors.New("")
	}

	if y >= s.yAxisSize {
		return errors.New("")
	}

	cellState := s.cells[y][x]
	s.cells[y][x] = !cellState

	return nil
}

func (s *Simulation) String() string {
	stringRepresentation := ""
	cells := s.cells
	for _, y := range cells {
		for _, x := range y {
			var val int
			if x {
				val = 1
			}
			stringRepresentation += fmt.Sprintf("%d", val)
		}
		stringRepresentation += fmt.Sprintf("\n")
	}
	return stringRepresentation
}

func (s *Simulation) tickCell(x int, y int, state bool) {
	neighbours := 0

	if s.neighbourExistsEast(x, y) {
		neighbours += 1
	}

	// ... TODO

	s.cells[y][x] = true
}

func (s *Simulation) neighbourExistsEast(x int, y int) bool {
	if x == s.xAxisSize - 1 {
		// wrap
		neighbour := s.cells[y][0]
		return neighbour
	} else {
		neighbour := s.cells[y][x + 1]
		return neighbour
	}
}

