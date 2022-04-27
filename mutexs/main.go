package mutexs

import (
	"fmt"
	"sync"
)

type WaitGroup = sync.WaitGroup
type Mutex = sync.Mutex

var x = 0

func increment(wg *WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func Main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
