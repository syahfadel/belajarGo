package main

import (
	"fmt"
	"sync"
)

func bisa(data interface{}, c1 chan interface{}, lock *sync.Mutex, wg *sync.WaitGroup) {
	c1 <- data
}
func coba(data interface{}, c2 chan interface{}, lock *sync.Mutex, wg *sync.WaitGroup) {
	c2 <- data
}

func resultBisa(c1 chan interface{}, lock *sync.Mutex, wg *sync.WaitGroup) {
	fmt.Println("[bisa1 bisa2 bisa3] ", <-c1)
	lock.Unlock()
	wg.Done()
}
func resultCoba(c2 chan interface{}, lock *sync.Mutex, wg *sync.WaitGroup) {
	fmt.Println("[coba1 coba2 coba3] ", <-c2)
	lock.Unlock()
	wg.Done()
}
