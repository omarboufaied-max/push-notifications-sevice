package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("worker started")
	for {
		time.Sleep(10 * time.Second)
	}
}
