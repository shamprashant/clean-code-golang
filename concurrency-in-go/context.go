package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context, id int) {
	select {
	case <-time.After(2 * time.Second): // Simulating a long task
		fmt.Println("Finished:", id)
	case <-ctx.Done():
		fmt.Println("Cancelled:", id)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go slowOperation(ctx, i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Main Done")
}
