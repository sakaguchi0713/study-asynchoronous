package main

import (
	"fmt"
	"time"
	"runtime"
)

//
//func main() {
//	ch := fanIn(generator("Hello"), generator("Bye"))
//	for i := 0; i < 10; i++ {
//		msg1 := <-ch
//		fmt.Println(msg1.str)
//
//		msg2 := <-ch
//		fmt.Println(msg2.str)
//
//		<- msg1.block // reset channel, stop blocking
//		<- msg2.block
//	}
//}
//
//// fanIn is itself a generator
//func fanIn(ch1, ch2 <-chan Message) <-chan Message { // receives two read-only channels
//	new_ch := make(chan Message)
//	go func() { for { new_ch <- <-ch1 } }() // launch two goroutine while loops to continuously pipe to new channel
//	go func() { for { new_ch <- <-ch2 } }()
//	return new_ch
//}
//
//func generator(msg string) <-chan Message { // returns receive-only channel
//	ch := make(chan Message)
//	blockingStep := make(chan int) // channel within channel to control exec, set false default
//	go func() { // anonymous goroutine
//		for i := 0; ; i++ {
//			ch <- Message{fmt.Sprintf("%s %d", msg, i), blockingStep}
//			time.Sleep(time.Second)
//			blockingStep <- 1 // block by waiting for input
//		}
//	}()
//	return ch
//}

//func main() {
//	ch1 := generator("Hello")
//	ch2 := generator("Bye")
//	for i := 0; i < 5; i++ {
//		fmt.Println(<- ch1)
//		fmt.Println(<- ch2)
//	}
//}
//
//func generator(msg string) <-chan string { // returns receive-only channel
//	ch := make(chan string)
//	go func() { // anonymous goroutine
//		for i := 0; ; i++ {
//			ch <- fmt.Sprintf("%s %d", msg, i)
//			time.Sleep(time.Second)
//		}
//	}()
//	return ch
//}

//// takes two int channels, stores right val (+1) into left
//func f(left, right chan int) {
//	left <- 1 + <-right // bafter 1st right read, locks until left read
//}
//
//func main() {
//	const n = 10000
//
//	// construct an array of n+1 int channels
//	var channels [n + 1]chan int
//	for i := range channels {
//		channels[i] = make(chan int)
//	}
//
//	// wire n goroutines in a chain
//	for i := 0; i < n; i++ {
//		go f(channels[i], channels[i+1])
//	}
//
//	// insert a value into right-hand end
//	go func(c chan<- int) { c <- 1 }(channels[n])
//
//	// get value from the left-hand end
//	fmt.Println(<-channels[0])
//}


func main() {
	ch := make(chan string)
	go channel_print("Hello", ch)
	fmt.Println(runtime.NumGoroutine())
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch) // ends of channel block until both are ready
		// NOTE: golang supports buffered channels, like mailboxes (no sync)
	}
	fmt.Println("Done!")
}

func channel_print(msg string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Second)
	}
}