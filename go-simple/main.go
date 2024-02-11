package main

import "fmt"

// Default main.go
func main() {
	fmt.Println("Hello, World!")
	// fatal error: all goroutines are asleep - deadlock!
	// Sending or receiving from a nil channel will block forever
	// var ch chan int
	// ch <- 1

	// Unbuffered channels are channels without any capacity for storing values.
	// It will block until another goroutine is ready to receive/send that value.
	// fatal error: all goroutines are asleep - deadlock!
	//
	// c := make(chan int)
	// c <- 1
	// <-c
	//
	// ch := make(chan int) // Unbuffered channel
	// go func() {
	// 	ch <- 42 // Send value to channel
	// }()
	// value := <-ch      // Receive value from channel
	// fmt.Println(value) // Prints 42
}
