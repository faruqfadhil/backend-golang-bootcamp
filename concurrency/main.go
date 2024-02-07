package main

import (
	"fmt"
	"sync"
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
	// mainproducerWithoutRange()
	// mainWriteBufferChan()
	// mainSelect()
	// mainProcessWithWg()
	mainIncrementRC()
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
	done := make(chan bool, 0)
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

func writeBufferChan(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func mainWriteBufferChan() {
	ch := make(chan int, 2)
	go writeBufferChan(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func mainSelect() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

func processWithWg(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func mainProcessWithWg() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go processWithWg(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

var x = 0

func incrementRC(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func mainIncrementRC() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementRC(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
