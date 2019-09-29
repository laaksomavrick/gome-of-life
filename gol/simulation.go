package gol

import (
	"errors"
	"fmt"
	"log"
)

type Simulation struct {
	xAxisSize int
	xAxisSizeZeroIndexed int
	yAxisSize int
	yAxisSizeZeroIndexed int
	cells [][]bool
	queue chan[][]bool
}

// Origin is top left corner
// x extends horizontally
// y extends vertically
// cells[y][x]
func NewSimulation(xAxisSize int, yAxisSize int, queue chan[][]bool) Simulation {
	if xAxisSize < 1 {
		log.Fatalf("xSize must be greater than or equal to 1")
	}

	if yAxisSize < 1 {
		log.Fatalf("ySize must be greater than or equal to 1")
	}

	cells := newCells(xAxisSize, yAxisSize)

	return Simulation{
		xAxisSize: xAxisSize,
		xAxisSizeZeroIndexed: xAxisSize - 1,
		yAxisSize: yAxisSize,
		yAxisSizeZeroIndexed: yAxisSize - 1,
		cells: cells,
		queue: queue,
	}
}

func (s *Simulation) Simulate() {
	go func() {
		for {
			state := s.cells
			s.queue <- state
			s.Tick()
		}
	}()
}

func (s *Simulation) Tick() {
	newCellsState := newCells(s.xAxisSize, s.yAxisSize)
	for y := range s.cells {
		for x := range s.cells[y] {
			state := s.cells[y][x]
			newState := s.tickCell(x, y, state)
			newCellsState[y][x] = newState
		}
	}
	s.cells = newCellsState
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

func newCells(xSize int, ySize int) [][]bool {
	cells := make([][]bool, ySize)
	for row := range cells {
		cells[row] = make([]bool, xSize)
	}
	return cells
}

func (s *Simulation) tickCell(x int, y int, state bool) bool {
	var newState bool
	neighbours := 0
	funcs := []func(int, int) bool{
		s.neighbourExistsEastFrom,
		s.neighbourExistsNorthEastFrom,
		s.neighbourExistsNorthFrom,
		s.neighbourExistsNorthWestFrom,
		s.neighbourExistsWestFrom,
		s.neighbourExistsSouthWestFrom,
		s.neighbourExistsSouthFrom,
		s.neighbourExistsSouthEastFrom,
	}

	for _, fn := range funcs {
		neighbourExists := fn(x, y)
		if neighbourExists {
			neighbours += 1
		}
	}

	// any live cell with fewer than 2 live neighbours dies
	// any live cell with 2 or 3 live neighbours does nothing
	// any live cell with more than 3 live neighbours dies
	// any dead cell with exactly three live neighbours becomes a live cell

	if state == true {
		if neighbours < 2 {
			newState = false
		} else if neighbours == 2 || neighbours == 3 {
			newState = state
		} else if neighbours > 3 {
			newState = false
		}
	} else {
		if neighbours == 3 {
			newState = true
		}
	}

	return newState
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

