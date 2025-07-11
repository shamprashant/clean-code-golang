// previous task(2) with 5 goroutines sending to the channel concurrently

package main

import (
	"fmt"
	"sync"
)

func sendDataToChannel(i int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- fmt.Sprintf("Hello from Goroutine %d", i)
}

func main() {
	ch := make(chan string, 5)
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go sendDataToChannel(i+1, ch, &wg)
	}
	wg.Wait()
	close(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}
