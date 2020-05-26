package bug2

import (
	"sync"
)

func bug2(n int, foo func(int) int, out chan int) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < n; i++ {
		wg.Add(1)
		mutex.Lock()
		go func(i int) {
			out <- foo(i)
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(out)
}
