package main

import (
	"fmt"
	"time"
)

func hello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello world goroutine")
}
func singleGoroutine() {
	go hello()
	time.Sleep(2 * time.Second)
	fmt.Println("main function")
}

func main() {
	// singleGoroutine()
	// multipleGoroutine()
	// mainChan()
	// mainHelloWithChanAndSleep()
	// mainSendDataUnidirect()
	// mainproducer()
	mainproducerWithoutRange()
}

func multipleGoroutine() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func helloWithChan(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}
func mainChan() {
	done := make(chan bool)
	go helloWithChan(done)
	<-done
	fmt.Println("main function")
}

func helloWithChanAndSleep(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
func mainHelloWithChanAndSleep() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go helloWithChanAndSleep(done)
	<-done
	fmt.Println("Main received data")
}

func sendDataUnidirect(sendch chan<- int) {
	datas := []int{1, 2, 3}
	for _, d := range datas {
		sendch <- d
	}
	close(sendch)
}

func mainSendDataUnidirect() {
	sendch := make(chan int)
	go sendDataUnidirect(sendch)
	for {
		a, ok := <-sendch
		if !ok {
			break
		}
		fmt.Println(a)
	}
	fmt.Println("berhenti")
}

func closeChan(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func mainproducer() {
	ch := make(chan int)
	go closeChan(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
}

func mainproducerWithoutRange() {
	ch := make(chan int)
	go closeChan(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received ", v)
	}
}
