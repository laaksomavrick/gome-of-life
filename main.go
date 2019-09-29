package main

import (
	"gol/gol"
)

// every n ms a tick happens, and then a render() function is called
// we can have a thead calculating an array of states (e.g buffer up to 10) in the background
//     - every add/remove to the array must be locked
// render pops the first item in the list and uses that as it's state

//main.go
//config.go
//renderer.go
//simulation.go
//simulation_queue.go

func main() {

	queue := make(chan [][]bool, 10)

	renderer := gol.NewRenderer(queue)
	renderer.Init()

	x, y := renderer.Size()

	//fmt.Println(x)
	//fmt.Println(y)

	simulation := gol.NewSimulation(x, y, queue)

	_ = simulation.ToggleCell(20, 21)
	_ = simulation.ToggleCell(20, 22)
	_ = simulation.ToggleCell(20, 23)

	simulation.Simulate()

	renderer.Read()

}

//func asyncBigComputation(ch chan<- int) {
//	result := bigComputation()
//	ch <- result
//}
//
//func bigComputation() int {
//	return 42
//}
