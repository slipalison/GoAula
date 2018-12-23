package main

import (
	"fmt"
	"time"
)

//var waitGroup sync.WaitGroup

func main() {
	channelTask()
	mensageria()
}

var msg chan string
var done chan bool
var contador int

func mensageria() {
	msg = make(chan string)
	done = make(chan bool)

	// inverti a ordem para mostra que isso não importa no Channel
	go sendMessage("olar")
	go receiveMessage()

	go sendMessage("olar1")
	go receiveMessage()

	go sendMessage("olar2")
	go receiveMessage()

	go sendMessage("olar3")
	go receiveMessage()

	<-done
}

func sendMessage(text string) {
	msg <- text
	contador++
}

func receiveMessage() {
	fmt.Println(<-msg)
	contador--
	fmt.Printf("contador esta em :%v\n", contador)

	if contador == 0 {
		done <- true
	}
}

func channelTask() {

	done := make(chan bool)
	go task(done)
	// le o canal
	<-done
}

func task(done chan bool) {
	fmt.Println("Start task")
	time.Sleep(time.Second * 3)
	fmt.Println("End task")

	// envia para o canal
	done <- true
}
