package main

/*
## Custome Sync.WaitGroup

Write our own version of `Sync.WaitGroup` which basically mimics the overall
behaviour of the native liberary.

**Time Required: 15 mins max**

### Thinks look for

- Not doing long polling, use channels and mutex
- What if not `add` is called
- What if not `wait` is called
- And thinking for other corner cases

*/

import (
	"fmt"
	"sync"
	"time"
)

type WaitGroup struct {
	counter int
	done    chan bool
	wait    bool
	mu      sync.Mutex
}

func main() {
	wg := newWaitGroup()
	count := 10
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%d\t done\n", i+1)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func newWaitGroup() *WaitGroup {
	return &WaitGroup{
		counter: 0,
		done:    make(chan bool),
	}
}

func (cwg *WaitGroup) Add(x int) {
	cwg.mu.Lock()
	defer cwg.mu.Unlock()
	cwg.counter += x

	// send signal on all goroutines complete and if wait is started
	if cwg.counter == 0 && cwg.wait {
		cwg.done <- true
	}
}

func (cwg *WaitGroup) Done() {
	cwg.Add(-1)
}

func (cwg *WaitGroup) Wait() {
	cwg.wait = true
	if cwg.counter == 0 {
		// all goroutines might completed before wait
		return
	}
	<-cwg.done
}
