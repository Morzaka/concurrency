package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "yes"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
