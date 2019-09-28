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

		It("should be find a neighbor directly east", func() {
			_ = simulation.ToggleCell(1, 0)
			neighbour := simulation.neighbourExistsEast(0, 0)
			Expect(neighbour).To(Equal(true))
		})

		It("should be find a neighbor wrapped east", func() {
			_ = simulation.ToggleCell(0, 0)
			neighbour := simulation.neighbourExistsEast(4, 0)
			Expect(neighbour).To(Equal(true))
		})

	})


})
