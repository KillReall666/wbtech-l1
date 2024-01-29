package main

import (
	"fmt"
	"sync"
)

type Storage struct {
	mu sync.RWMutex
	m  map[int]int
}

func (s *Storage) Add(key int, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func main() {
	m := Storage{
		mu: sync.RWMutex{},
		m:  make(map[int]int),
	}

	var wg sync.WaitGroup
	numOfGoroutines := 10

	for i := 0; i < numOfGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			m.Add(i, i)
			wg.Done()
		}(i)
	}

	wg.Wait() //ждем выполнения всех воркеров

	fmt.Println(m.m)
}
