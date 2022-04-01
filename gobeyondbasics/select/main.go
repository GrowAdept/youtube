package main

import (
	"fmt"
	"time"
)

type Result struct {
	message string
	err     error
}

func sender1(ch chan string, s string) {
	time.Sleep(time.Second * 1)
	ch <- s
}

func sender2(ch chan int, i int) {
	time.Sleep(time.Second * 4)
	ch <- i
}

func sender3(ch chan Result, r Result) {
	time.Sleep(time.Second * 8)
	ch <- r
}

func exit(ch chan bool) {
	time.Sleep(time.Second * 8)
	ch <- true
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan int)
	chan3 := make(chan Result)
	done := make(chan bool)

	go sender1(chan1, "Hello Word")
	go sender2(chan2, 71)
	go sender3(chan3, Result{"Hello", nil})
	go exit(done)

	for {
		select {
		case s := <-chan1:
			fmt.Println("received", s, "from chan1")
		case i := <-chan2:
			fmt.Println("received", i, "from chan2")
		case result := <-chan3:
			if result.err != nil {
				fmt.Println("err:", result.err)
				return
			}
			fmt.Println("received", result.message, "from chan3")
			// timeout in case resources take too long to recieve
		/*
			case <-time.After(time.Second * 10):
				fmt.Println("took too long, good bye")
				return
		*/
		// we can send a message to end the loop
		case b := <-done:
			fmt.Println("received,", b, "from done chan, time to exit")
			return

		default:
			fmt.Println("no communication")
			time.Sleep(time.Second * 1)
		}

	}
}
