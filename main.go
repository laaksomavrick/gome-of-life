package main

import (
	"fmt"
	"time"
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

const (
	tickRate = time.Second * 1
)

func main() {

	// Buffered channel
	comps := make(chan int, 10)

	for {
		// Do a lot of big computations in the background
		go asyncBigComputation(comps)

		select {
		case comp := <-comps:
			fmt.Println(comp)
			time.Sleep(tickRate)

		default:
			fmt.Println("waiting...")
			time.Sleep(tickRate)
		}
	}

}

func asyncBigComputation(ch chan<- int) {
	result := bigComputation()
	ch <- result
}

func bigComputation() int {
	return 42
}
