package main

import (
	"fmt"
	"sync"
)

func main() {
	// sequencing therough channels
	var ans = []int{}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		<-ch1
		mut.Lock()
		ans = append(ans, 1)
		mut.Unlock()
		wg.Done()
		close(ch2)
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		<-ch2
		mut.Lock()
		ans = append(ans, 2)
		mut.Unlock()
		wg.Done()
		close(ch3)
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		<-ch3
		mut.Lock()
		ans = append(ans, 3)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	close(ch1)
	wg.Wait()
	fmt.Println(ans)

	// send and recive through channel
	ch := make(chan int, 2)
	wg.Add(2)
	go func(ch <-chan int) {
		val, err := <-ch
		fmt.Println("recieved the value", err)
		fmt.Println(val)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 8
		wg.Done()
	}(ch)
	wg.Wait()

}
