package main

import (
	"context"
	"fmt"
)

// ToDo ТАК РОБИТИ НЕ МОЖНА
func doSomething() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	someArg := "loremipsum"
	go doSomethingElse(someArg)
}

func doSomethingElse(someArg string) {
	fmt.Println(someArg)
}

func main() {
	doSomething()
}
