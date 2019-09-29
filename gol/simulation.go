package gol

import (
	"errors"
	"fmt"
	"log"
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
	xAxisSizeZeroIndexed int
	yAxisSize int
	yAxisSizeZeroIndexed int
	cells [][]bool
}

// Origin is top left corner
// x extends horizontally
// y extends vertically
// cells[y][x]
func NewSimulation(xSize int, ySize int) *Simulation {
	if xSize < 1 {
		log.Fatalf("xSize must be greater than or equal to 1")
	}

	if ySize < 1 {
		log.Fatalf("ySize must be greater than or equal to 1")
	}

	cells := make([][]bool, ySize)
	for row := range cells {
		cells[row] = make([]bool, xSize)
	}
	return &Simulation{
		xAxisSize: xSize,
		xAxisSizeZeroIndexed: xSize - 1,
		yAxisSize: ySize,
		yAxisSizeZeroIndexed: ySize - 1,
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

	if s.neighbourExistsEastFrom(x, y) {
		neighbours += 1
	}

	// ... TODO

	s.cells[y][x] = true
}

func (s *Simulation) neighbourExistsEastFrom(x int, y int) bool {
	if x == s.xAxisSizeZeroIndexed {
		neighbour := s.cells[y][0]
		return neighbour
	} else {
		neighbour := s.cells[y][x + 1]
		return neighbour
	}
}

func (s *Simulation) neighbourExistsNorthEastFrom(x int, y int) bool {
	if y == 0 || x == s.xAxisSizeZeroIndexed {
		var wrappedXPos int
		if x == s.xAxisSizeZeroIndexed {
			wrappedXPos = 0
		} else {
			wrappedXPos = x + 1
		}
		neighbour := s.cells[s.yAxisSizeZeroIndexed][wrappedXPos]
		return neighbour
	} else {
		neighbour := s.cells[y - 1][x + 1]
		return neighbour
	}
}

func (s *Simulation) neighbourExistsNorthFrom(x int, y int) bool {
	if y == 0 {
		neighbour := s.cells[s.yAxisSizeZeroIndexed][x]
		return neighbour
	} else {
		neighbour := s.cells[y - 1][x]
		return neighbour
	}
}

func (s *Simulation) neighbourExistsNorthWestFrom(x int, y int) bool {
	if y == 0 || x == 0 {
		var wrappedXPos int
		if x == 0 {
			wrappedXPos = s.xAxisSizeZeroIndexed
		} else {
			wrappedXPos = x - 1
		}
		neighbour := s.cells[s.yAxisSizeZeroIndexed][wrappedXPos]
		return neighbour
	} else {
		neighbour := s.cells[y - 1][x - 1]
		return neighbour
	}
}

func (s *Simulation) neighbourExistsWestFrom(x int, y int) bool {
	if x == 0 {
		neighbour := s.cells[y][s.xAxisSizeZeroIndexed]
		return neighbour
	} else {
		neighbour := s.cells[y][x - 1]
		return neighbour
	}
}

func (s *Simulation) neighbourExistsSouthWestFrom(x int, y int) bool {
	if y == s.yAxisSizeZeroIndexed || x == 0 {
		var wrappedXPos int
		if x == 0 {
			wrappedXPos = s.xAxisSizeZeroIndexed
		} else {
			wrappedXPos = x - 1
		}
		neighbour := s.cells[0][wrappedXPos]
		return neighbour
	} else {
		neighbour := s.cells[y + 1][x - 1]
		return neighbour
	}
}

func (s *Simulation) neighbourExistsSouthFrom(x int, y int) bool {
	if y == s.yAxisSizeZeroIndexed {
		neighbour := s.cells[0][x]
		return neighbour
	} else {
		neighbour := s.cells[y + 1][x]
		return neighbour
	}
}

func (s *Simulation) neighbourExistsSouthEastFrom(x int, y int) bool {
	if y == s.yAxisSizeZeroIndexed || x == s.xAxisSizeZeroIndexed {
		var wrappedXPos int
		if x == s.xAxisSizeZeroIndexed {
			wrappedXPos = 0
		} else {
			wrappedXPos = x + 1
		}
		neighbour := s.cells[0][wrappedXPos]
		return neighbour
	} else {
		neighbour := s.cells[y + 1][x + 1]
		return neighbour
	}
}

