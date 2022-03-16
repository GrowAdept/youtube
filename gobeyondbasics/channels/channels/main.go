package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func printA(c chan string) {
	defer wg.Done()
	c <- "A"
	fmt.Println("printA() done")
}

func printB(c chan string) {
	defer wg.Done()
	c <- "B"
	fmt.Println("printB() done")
}

func printC(c chan string) {
	defer wg.Done()
	c <- "C"
	fmt.Println("printC() done")
}

func printD(c chan string) {
	defer wg.Done()
	c <- "D"
	fmt.Println("printD() done")
}

func main() {
	//     |goroutines1|      <-->      |channel|      <-->      |goroutine2|
	// the values the channel accepts are type safe
	var myChan chan string // declared but not initialized channels are nil
	fmt.Println("myChan:", myChan)
	// nil channel is never ready for communication
	/*
		myChan <- "E" // this will deadloack, not initialized
		fmt.Println(<-myChan)
	*/
	// make takes our channel type and capacity that will set the size in the buffer
	myChan = make(chan string, 1)
	fmt.Println("myChan:", myChan)
	myChan <- "E"
	fmt.Println(<-myChan)
	// length of channel is increased / decreased with sends and recieves

	fmt.Println("myChan capacity:", cap(myChan))
	fmt.Println("myChan length:", len(myChan))
	myChan <- "E"
	fmt.Println("myChan length:", len(myChan))
	fmt.Println(<-myChan)
	fmt.Println("myChan length:", len(myChan))

	wg.Add(4)
	// try change channel buffer to 4 instead of 1 and all the function not get blocked
	go printA(myChan)
	go printB(myChan)
	go printC(myChan)
	go printD(myChan)
	time.Sleep(time.Second * 10)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	wg.Wait()

}
