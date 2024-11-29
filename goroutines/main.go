package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // send data through that channel to where we started the go routine, the arrow describes the direction of data flow
	close(doneChan)
}

// without channels, dispatching the 4 goroutines is too fast for printing the messages
func main() {
	// Dispatch 4 go routines
	dones := make([]chan bool, 4)
	//done := make(chan bool)                            // create a single channel that sends a bool
	dones[0] = make(chan bool)
	go greet("Nice to meet you!", dones[0]) // run a function as goroutine which runs the function in parallel
	dones[1] = make(chan bool)
	go greet("How are you?", dones[1]) // parallel
	dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", dones[2]) // parallel
	dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", dones[3]) // parallel
	//fmt.Println(<-dones)                                   // emit data, data that comes out of that channel, go will only continue when data came out of that channel
	for _, done := range dones { // wait for all channel operations to finish
		<-done
	}

	// or

	/*for doneChan := range done { // this is supported by go as well, get directly the booleans from the single channel
		fmt.Println(doneChan)
	}*/
}
