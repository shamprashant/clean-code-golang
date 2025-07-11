// Write a program where 10 goroutines print “Hello” with their number.

package main

import (
	"fmt"
	"sync"
)

func sayHello(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello: ", i)
}

// func main() {
// 	var wg sync.WaitGroup
// 	for i := range 10 {
// 		wg.Add(1)
// 		go sayHello(i, &wg)
// 	}
// 	wg.Wait()
// }
