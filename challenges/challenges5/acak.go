package main

import (
	"fmt"
	"sync"
)

func Acak() {
	fmt.Println("\n=================Goroutine Menampilkan Acak=================")
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

	for i := 0; i < 8; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("[bisa1 bisa2 bisa3] ", msg1)
			wg.Done()
		case msg2 := <-c2:
			fmt.Println("[coba1 coba2 coba3] ", msg2)
			wg.Done()
		}
	}

	wg.Wait()

}
