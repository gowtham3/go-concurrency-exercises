package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.

		mu.Lock()
		for len(sharedRsc) == 0 {
			//time.Sleep(1 * time.Millisecond)
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.

		mu.Lock()
		for len(sharedRsc) == 0 {
			//time.Sleep(1 * time.Millisecond)
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc2"])
		 mu.Unlock()
	}()

	// writes changes to sharedRsc
	mu.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	c.Broadcast()
	mu.Unlock()

	wg.Wait()
}
