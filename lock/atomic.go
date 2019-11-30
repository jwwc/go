package main

import (
	"sync"
	"sync/atomic"
	"fmt"
)

var total uint64

func worker(wg *sync.WaitGroup,user string) {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total, i)
		fmt.Println(total,user)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg,"user1")
	go worker(&wg,"user2")
	wg.Wait()
}
