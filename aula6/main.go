package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	for index := 0; index <= 10; index++ {
		waitGroup.Add(1)
		go hardTask(strconv.Itoa(index))
	}
	waitGroup.Wait()
	fmt.Println("DONE")
}

func hardTask(name string) {
	defer waitGroup.Done()
	for i := 0; i <= 10; i++ {
		time.Sleep(1 + time.Second)
		fmt.Printf("Hard Taks %s...,\n", name)
	}
	fmt.Printf("hard task %s DONE\n", name)
}
