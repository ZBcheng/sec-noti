package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func writeChan(channel chan int) {
	for i := 0; i < 20; i++ {
		channel <- i
	}
	wg.Done()
}

func main() {
	channel := make(chan int, 10)
	// closeChan := make(chan bool, 1)
	go func() {
		for i := 0; i < 20; i++ {
			wg.Add(1)
			channel <- i
			wg.Done()
		}
	}()
	for v := range channel {
		fmt.Println(v)
	}
	wg.Wait()
	// redishandler.Subscribe("bot", channel)
}
