package main

import (
	"fmt"
	"sync"
)

func main() {
	var ans = []int{}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		ans = append(ans, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		ans = append(ans, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		ans = append(ans, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println(ans)
}
