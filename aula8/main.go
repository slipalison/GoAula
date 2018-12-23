package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	done := make(chan bool)
	go hardTask(done)
	waitGroup.Add(1)
	go loader(done)
	waitGroup.Wait()
}

// limita a função para apenas escrever no channel
func hardTask(done chan<- bool) {
	fmt.Printf("\r")
	fmt.Println("Start task")
	time.Sleep(time.Second * 3)
	done <- true
}

//apenas esculta o channel
func loader(done <-chan bool) {
	defer waitGroup.Done()
	// loop eterno até que aconteça algo
	i := 0
	load := []rune(`|\-/`)
	for {
		// parecido com switch case
		select {
		case <-done:
			fmt.Printf("\r")
			fmt.Println("Done")
			return
		default:
			fmt.Printf("\r")
			fmt.Printf(string(load[i]))
			time.Sleep(time.Microsecond * 100000)
			i++
			if i == len(load) {
				i = 0
			}
		}
	}
}
