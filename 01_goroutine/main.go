package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		count("cat")
		wg.Done()
	}()

	go func() {
		count("dog")
		wg.Done()
	}()

	wg.Wait()
}

func count(thing string) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		fmt.Println("Lolo")
		wg.Done()
	}()

	wg.Wait()
	for i := 0; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second / 2)
	}
}
