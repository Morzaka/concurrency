package main

import (
	"fmt"
	"time"
)

//

func main() {
	count("sheep")
}

func count(thing string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(thing)
		time.Sleep(time.Second * 1)
	}
}
