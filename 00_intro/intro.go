package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	names := []string{"Богдан", "Юрій", "Андріана", "Олег", "Віталій", "Любомир"}

	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(len(names))

	for _, name := range names {
		go printNames(name, &wg)
		//printNames(name, &wg)
	}

	wg.Wait()
	fmt.Printf("Програма виконувалась за %s\n", time.Since(start))
}

//func printNames(name string) {
//	fmt.Println("Name: ", name)
//}

//func printNames(name string) {
//	result := 0.0
//	for i := 0; i < 100_000_000; i++ {
//		result += math.Pi * math.Sin(float64(len(name)))
//	}
//	fmt.Println("Name: ", name)
//}

func printNames(name string, wg *sync.WaitGroup) {
	result := 0.0
	for i := 0; i < 100_000_000; i++ {
		result += math.Pi * math.Sin(float64(len(name)))
	}
	fmt.Println("Name: ", name)
	wg.Done()
}
