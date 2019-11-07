package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	backOff := make(chan bool, 1)

	for i := 1; i <= 100; i++ {
		waitGroup.Add(1)
		go doSomething(i, &waitGroup, backOff)
		time.Sleep(time.Millisecond)

		select {
		case isBackOffRequired := <-backOff:
			if isBackOffRequired {
				backOffABit()
			}
		default:
		}
	}

	waitGroup.Wait()
}

func backOffABit() {
	fmt.Println("🚨 An error was reported, sleeping for 10 seconds before continuing")
	time.Sleep(time.Second * 10)
}

func doSomething(someID int, waitGroup *sync.WaitGroup, backOff chan bool) {

	fmt.Println("🔈 I'm number", someID, "and I'm starting")
	if someID%7 == 0 {
		backOff <- true
	}
	time.Sleep(time.Second * 2)
	fmt.Println("🔈 I'm number", someID, "and I'm stopping")
	waitGroup.Done()
}
