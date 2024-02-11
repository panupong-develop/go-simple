package main

import "fmt"

func sum(num *[2]int32) {
	fmt.Println("Inside func")
	fmt.Printf("0: %p\n", &num[0])
	fmt.Printf("1: %p\n", &num[1])
}

// Default main.go
func main() {
	var numArr = [2]int32{1, 2}
	fmt.Printf("0: %p\n", &numArr[0])
	fmt.Printf("1: %p\n", &numArr[1])

	sum(&numArr)

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
