// package main

// import "fmt"

// func main() {

// 	messages := make(chan string)

// 	messages <- "buffered"
// 	messages <- "channel"

// 	fmt.Println(<-messages)
// 	fmt.Println(<-messages)
// 	fmt.Println(<-messages)

// }

package main

import (
	"fmt"
	"time"
)

func main() {
	// Creating an unbuffered channel
	messages := make(chan string)

	// Attempting to send two messages in quick succession will result in a deadlock
	go func() {
		messages <- "Message 1"
		fmt.Println("Sent: Message 1")

		time.Sleep(2 * time.Second) // Simulate some work

		messages <- "Message 2" // This line will block until a receiver is ready
		fmt.Println("Sent: Message 2")
	}()

	// Main goroutine receiving messages from the channel
	message1 := <-messages
	fmt.Println("Received:", message1)

	// Sleep added to demonstrate that the sender is blocked until a receiver is ready
	time.Sleep(1 * time.Second)

	message2 := <-messages
	fmt.Println("Received:", message2)
}
