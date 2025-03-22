package main

import (
	"fmt"
	"time"
)

func sendData1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Message from Channel 1"
}

func sendData2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Message from Channel 2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendData1(ch1)
	go sendData2(ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}

	fmt.Println("All messages received.")
}
