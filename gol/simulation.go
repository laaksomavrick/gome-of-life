package gol

// 2 dimensional grid of cells, infinite (ie wrapping)
// each cell is in one of two states; alive or dead
// each cell interacts with it's neighbors (8), n, ne, w, sw, s, se, e, ne

// each tick:
// any live cell with fewer than 2 live neighbours dies
// any live cell with 2 or 3 live neighbours does nothing
// any live cell with more than 3 live neighbours dies

// any dead cell with exactly three live neighbours becomes a live cell

// Alive = true, Dead = false

type Cell bool

type Simulation struct {
	RowSize int
	ColSize int
	cells [][]Cell
}

func (s *Simulation) NewSimulation(rowSize int, colSize int) *Simulation {
	cells := make([][]Cell, rowSize)
	for row := range cells {
		cells[row] = make([]Cell, colSize)
	}
	return &Simulation{
		RowSize: rowSize,
		ColSize: colSize,
		cells: cells,
	}
}

func (s *Simulation) Tick() [][]Cell {
	for row := range s.cells {
		for col := range s.cells[row] {
			state := s.cells[row][col]
			s.tickCell(row, col, state)
		}
	}

	return s.cells
}

func (s *Simulation) tickCell(row int, col int, state Cell) {
	neighbours := 0

	if s.neighbourExistsEast(row, col) {
		neighbours += 1
	}

	// ... TODO

	s.cells[row][col] = true
}

func (s *Simulation) neighbourExistsEast(row int, col int) Cell {
	if row == s.RowSize - 1 {
		// wrap
		neighbour := s.cells[0][col]
		return neighbour
	} else {
		neighbour := s.cells[row + 1][col]
		return neighbour
	}
}

// TEST

