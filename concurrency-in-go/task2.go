// Launch 5 goroutines that send a message (like "Message from 1") to a shared channel, and collect and print the messages in main.

package main

// func main() {
// 	ch := make(chan string, 5)

// 	for i := 1; i <= 5; i++ {
// 		ch <- fmt.Sprintf("Hello from %d", i)
// 	}

// 	close(ch)
// 	// ✅ range ch automatically reads values from the channel until it is closed.
// 	for msg := range ch {
// 		fmt.Println(msg)
// 	}
// }
