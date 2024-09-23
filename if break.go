package main

import (
	"fmt"
	"time"
)

func main() {
	count := 10

	for {
		if count < 0 {
			break
		}
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
}
