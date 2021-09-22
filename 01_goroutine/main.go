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
		time.Sleep(time.Second / 2)
		fmt.Println(i, thing)
	}
}
