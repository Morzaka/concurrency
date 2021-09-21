package main

import (
	"fmt"
	"time"
)

func main() {
	count("sheep")
}

func count(thing string) {
	for i := 0; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}
