package main

import (
	"fmt"
)

// A channel is like a pipe 🛠 that allows safe communication between goroutines. But channels have rules:
// - Unbuffered channels block until both sender and receiver are ready.(like a handshake)
// - Buffered channels allow sending data without waiting (up to a limit).
// - Reading from an empty channel blocks until data is available.

func main() {
	ch1 := make(chan int)    //Unbuffered
	ch2 := make(chan int, 2) //Buffered
	ch3 := make(chan int, 4)

	go input(ch3, 4)

	go func() {
		ch1 <- 5
	}()

	ch2 <- 1
	ch2 <- 2
	close(ch2)

	fmt.Println(<-ch1)
	for i := range ch2 { // This loop keeps reading from ch2 until it is closed. However, you never close ch2, so the loop waits forever for new values after consuming 1 and 2!
		fmt.Println(i)
	}

	// can also use comma, ok syntaxcan also use comma, ok
	for {
		val, ok := <-ch3

		if !ok {
			break
		}

		fmt.Println("Val:", val)
	}

}

func input(ch chan int, count int) {
	defer close(ch)
	for i := 0; i < count; i++ {
		ch <- i
	}

}
