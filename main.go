package main

import (
	"gol/gol"
)

func main() {
	queue := make(chan [][]bool, 10)

	renderer := gol.NewRenderer(queue)
	renderer.Init()

	x, y := renderer.Size()

	simulation := gol.NewSimulation(x - 1, y - 1, queue)

	_ = simulation.ToggleCell(20, 21)
	_ = simulation.ToggleCell(20, 22)
	_ = simulation.ToggleCell(20, 23)

	_ = simulation.ToggleCell(31, 30)
	_ = simulation.ToggleCell(32, 31)
	_ = simulation.ToggleCell(30, 32)
	_ = simulation.ToggleCell(31, 32)
	_ = simulation.ToggleCell(32, 32)

	simulation.Simulate()

	renderer.Read()
}
