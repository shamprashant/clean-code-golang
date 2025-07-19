/*
5 goroutines fetch data

Timeout is set to 1.5 seconds

Each fetch takes 1–3 seconds randomly

Cancel the rest if timeout occurs?
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fetchData(wg *sync.WaitGroup, ctx context.Context, id int) {
	defer wg.Done()
	select {
	case <-time.After(time.Second * time.Duration(rand.Intn(3)+1)):
		fmt.Println("Data fetched from Goroutine: ", id)
	case <-ctx.Done():
		fmt.Printf("Goroutine %d cancelled!\n", id)
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second+time.Second/2)
	rand.Seed(time.Now().UnixNano()) // Add this to vary delay each time

	defer cancel()

	for id := range 5 {
		wg.Add(1)
		go fetchData(&wg, ctx, id)
	}
	wg.Wait()
}
