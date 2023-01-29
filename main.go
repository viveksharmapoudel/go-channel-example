package main

import (
	"fmt"
	queue_datastructure "go-channel-example/queue-datastructure"
	"sync"
	"time"
)

func main() {
	bufferSize := 5
	counter := 20
	q := queue_datastructure.NewQueue(bufferSize)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i := 0; i < counter; i++ {
			q.Push(i)
			time.Sleep(time.Second)

		}
	}()
	go func() {
		for j := 0; j < counter-10; j++ {
			fmt.Println("Popped data :", q.Pop())
		}
		wg.Done()
	}()
	wg.Wait()
}
