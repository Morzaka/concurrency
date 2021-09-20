package main

import (
	"fmt"
	"time"
)

func measurementsOfRuntime() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Програма виконувалась за %v секунди\n", time.Since(start).Seconds())
	}
}

func main() {
	defer measurementsOfRuntime()()

	names := []string{"Богдан", "Юрій", "Андріана", "Олег", "Віталій", "Любомир"}

	//runtime.GOMAXPROCS(8)
	//var wg sync.WaitGroup
	//wg.Add(len(names))

	for _, name := range names {
		//go printNames(name, &wg)
		printNames(name)
	}

	//wg.Wait()
}

func printNames(name string) {
	fmt.Println("Name: ", name)
}

//func printNames(name string) {
//	result := 0.0
//	for i := 0; i < 100_000_000; i++ {
//		result += math.Pi * math.Sin(float64(len(name)))
//	}
//	fmt.Println("Name: ", name)
//}

//func printNames(name string) {
//	result := 0.0
//	for i := 0; i < 100_000_000; i++ {
//		result += math.Pi * math.Sin(float64(len(name)))
//	}
//	fmt.Println("Name: ", name)
//}

//func printNames(name string, wg *sync.WaitGroup) {
//	result := 0.0
//	for i := 0; i < 100_000_000; i++ {
//		result += math.Pi * math.Sin(float64(len(name)))
//	}
//	fmt.Println("Name: ", name)
//	wg.Done()
//}