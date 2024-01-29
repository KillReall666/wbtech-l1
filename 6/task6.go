package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// использование тикера
func tikerGoroutine() {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			fmt.Println("Stopping timerGoroutine...")
			return
		default:
			fmt.Println("timerGoroutine working...")
			time.Sleep(time.Second)
		}
	}
}

// с использованием time
func timeGoroutine() {
	fmt.Println("timeGoroutine sleeping...")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("timeGoroutine woke up.")
}

// с использованием канала завершения
func structGoroutine(ch chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("Stopping boolGoroutine...")
			return
		default:
			fmt.Println("boolGoroutine working...")
			time.Sleep(time.Second)
		}
	}

}

// с использованием flag
func flagGoroutine() {
	defer wg.Done()
	for {
		if stopFlag {
			fmt.Println("Stopping flagGoroutine...")
			return
		}
		fmt.Println("flagGoroutine working...")
		time.Sleep(time.Second)
	}
}

// с использованием context
func ctxGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping ctxGoroutine...")
			return
		default:
			fmt.Println("ctxGoroutine working...")
			time.Sleep(1 * time.Second)
		}
	}
}

var stopFlag bool
var wg sync.WaitGroup

func main() {
	go tikerGoroutine()
	time.Sleep(6 * time.Second)

	ch := make(chan struct{})
	go structGoroutine(ch)
	time.Sleep(5 * time.Second)
	ch <- struct{}{}

	wg.Add(1)
	go flagGoroutine()
	time.Sleep(5 * time.Second)
	stopFlag = true
	wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	go ctxGoroutine(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	go timeGoroutine()
	time.Sleep(5 * time.Second)
}
