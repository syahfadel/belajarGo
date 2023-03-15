package main

import (
	"fmt"
	"sync"
)

func Rapih() {
	fmt.Println("\n=================Goroutine Menampilkan Rapih=================")
	var lock sync.Mutex
	var wg sync.WaitGroup
	c1 := make(chan interface{})
	c2 := make(chan interface{})

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go coba(i, c2, &lock, &wg)
		wg.Add(1)
		go bisa(i, c1, &lock, &wg)
	}

	for i := 0; i < 4; i++ {
		lock.Lock()
		go resultCoba(c2, &lock, &wg)
		lock.Lock()
		go resultBisa(c1, &lock, &wg)
	}

	wg.Wait()

}
