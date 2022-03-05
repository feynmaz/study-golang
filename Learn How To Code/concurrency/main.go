package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Routines:", runtime.NumGoroutine())
	runtime.GOMAXPROCS(4)

	var counter int64 

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter:", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Routines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Routines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}