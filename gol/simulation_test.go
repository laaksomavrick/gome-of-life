package gol

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGol(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gol Suite")
}

var _ = Describe("Simulation", func() {
	var simulation *Simulation

	BeforeEach(func () {
		simulation = NewSimulation(5, 5)
	})

	Describe("Cells", func() {
		It("Doesn't crash when cell is in a corner", func () {
			type coord struct {
				x, y int
			}

			corners := []coord{
				{0, 0},
				{4, 0},
				{0, 4},
				{4, 4},
			}

			funcs := []func(int, int) bool{
				simulation.neighbourExistsEastFrom,
				simulation.neighbourExistsNorthEastFrom,
				simulation.neighbourExistsNorthFrom,
				simulation.neighbourExistsNorthWestFrom,
				simulation.neighbourExistsWestFrom,
				simulation.neighbourExistsSouthWestFrom,
				simulation.neighbourExistsSouthFrom,
				simulation.neighbourExistsSouthEastFrom,
			}

			for _, fn := range funcs {
				for _, coord := range corners {
					neighbour := fn(coord.x, coord.y)
					Expect(neighbour).To(Equal(false))
				}
			}
		})

		Describe("Find neighbour east", func () {
			It("should find a neighbour directly east", func() {
				_ = simulation.ToggleCell(1, 0)
				neighbour := simulation.neighbourExistsEastFrom(0, 0)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped east", func() {
				_ = simulation.ToggleCell(0, 0)
				neighbour := simulation.neighbourExistsEastFrom(4, 0)
				Expect(neighbour).To(Equal(true))
			})
		})

		Describe("Find neighbour north east", func() {
			It("should find a neighbour directly north east", func() {
				_ = simulation.ToggleCell(3, 1)
				neighbour := simulation.neighbourExistsNorthEastFrom(2, 2)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped north east", func() {
				_ = simulation.ToggleCell(3, 4)
				neighbour := simulation.neighbourExistsNorthEastFrom(2, 0)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped north east on a corner", func() {
				_ = simulation.ToggleCell(0, 4)
				neighbour := simulation.neighbourExistsNorthEastFrom(4, 0)
				Expect(neighbour).To(Equal(true))
			})
		})

		Describe("Find neighbour north", func () {
			It("should find a neighbour directly north", func() {
				_ = simulation.ToggleCell(0, 3)
				neighbour := simulation.neighbourExistsNorthFrom(0, 4)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped north", func() {
				_ = simulation.ToggleCell(0, 4)
				neighbour := simulation.neighbourExistsNorthFrom(0, 0)
				Expect(neighbour).To(Equal(true))
			})
		})

		Describe("Find neighbour north west", func() {
			It("should find a neighbour directly north west", func() {
				_ = simulation.ToggleCell(1, 1)
				neighbour := simulation.neighbourExistsNorthWestFrom(2, 2)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped north west", func() {
				_ = simulation.ToggleCell(1, 4)
				neighbour := simulation.neighbourExistsNorthWestFrom(2, 0)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped north west on a corner", func() {
				_ = simulation.ToggleCell(4, 4)
				neighbour := simulation.neighbourExistsNorthWestFrom(0, 0)
				Expect(neighbour).To(Equal(true))
			})
		})

		Describe("Find neighbour west", func () {
			It("should find a neighbour directly west", func() {
				_ = simulation.ToggleCell(0, 0)
				neighbour := simulation.neighbourExistsWestFrom(1, 0)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped west", func() {
				_ = simulation.ToggleCell(4, 0)
				neighbour := simulation.neighbourExistsWestFrom(0, 0)
				Expect(neighbour).To(Equal(true))
			})
		})

		Describe("Find neighbour south west", func() {
			It("should find a neighbour directly south west", func() {
				_ = simulation.ToggleCell(1, 3)
				neighbour := simulation.neighbourExistsSouthWestFrom(2, 2)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped south west", func() {
				_ = simulation.ToggleCell(1, 0)
				neighbour := simulation.neighbourExistsSouthWestFrom(2, 4)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped south west on a corner", func() {
				_ = simulation.ToggleCell(4, 0)
				neighbour := simulation.neighbourExistsSouthWestFrom(0, 4)
				Expect(neighbour).To(Equal(true))
			})
		})

		Describe("Find neighbour south", func () {
			It("should find a neighbour directly south", func() {
				_ = simulation.ToggleCell(0, 4)
				neighbour := simulation.neighbourExistsSouthFrom(0, 3)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped south", func() {
				_ = simulation.ToggleCell(0, 0)
				neighbour := simulation.neighbourExistsSouthFrom(0, 4)
				Expect(neighbour).To(Equal(true))
			})
		})

		Describe("Find neighbour south east", func() {
			It("should find a neighbour directly south east", func() {
				_ = simulation.ToggleCell(3, 3)
				neighbour := simulation.neighbourExistsSouthEastFrom(2, 2)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped south east", func() {
				_ = simulation.ToggleCell(3, 0)
				neighbour := simulation.neighbourExistsSouthEastFrom(2, 4)
				Expect(neighbour).To(Equal(true))
			})

			It("should find a neighbour wrapped south east on a corner", func() {
				_ = simulation.ToggleCell(0, 0)
				neighbour := simulation.neighbourExistsSouthEastFrom(4, 4)
				Expect(neighbour).To(Equal(true))
			})
		})

	})


})
