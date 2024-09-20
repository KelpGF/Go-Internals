package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered()
	// buffered()
	closing()
}

func unbuffered() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		fmt.Println("Value sent 1") // occurs in the same time that "Value received 1" sync
		ch <- 2
		fmt.Println("Value sent 2")
	}()

	fmt.Println("Waiting")
	<-ch
	fmt.Println("Value received 1")

	time.Sleep(2 * time.Second)
	fmt.Println("Waiting")
	<-ch
	fmt.Println("Value received 2")

	time.Sleep(3 * time.Second)
}

func buffered() {
	ch := make(chan int, 2)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		fmt.Println("Value sent 1")
		ch <- 2
		fmt.Println("Value sent 2")
		ch <- 3
		fmt.Println("Value sent 3")
		ch <- 4
		fmt.Println("Value sent 4")
	}()

	fmt.Println("Waiting")
	<-ch
	fmt.Println("Value received 1")

	time.Sleep(2 * time.Second)
	fmt.Println("Waiting")
	<-ch
	fmt.Println("Value received 2")

	time.Sleep(3 * time.Second)
}

func closing() {
	ch := make(chan int, 2)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		close(ch)
	}()

	<-ch
	<-ch

	time.Sleep(3 * time.Second)
}
