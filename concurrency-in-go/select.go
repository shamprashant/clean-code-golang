package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "data from channel 1"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "data from channel 2"
	}()

	select {
	case data := <-ch1:
		fmt.Println(data)
	case data := <-ch2:
		fmt.Println(data)
	case <-time.After(time.Second):
		fmt.Println("Nothing ready yet")
	}

}
