package main

import (
	"fmt"
	"sync"
)

func main() {
	doOnce()
}

func doOnce() {
	var once sync.Once
	var wg sync.WaitGroup
	wg.Add(10)
	onceBody := func() {
		fmt.Println("only once")
	}

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)

			defer wg.Done()
		}()
	}
	wg.Wait()
}
