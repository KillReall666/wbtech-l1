package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type MyChannel struct {
	ch       chan string
	isClosed bool
	once     sync.Once
}

func (mc *MyChannel) SafeClose() {
	fmt.Println("SafeClose: safety close channel")
	mc.once.Do(func() {
		close(mc.ch)
		mc.isClosed = true
	})
}

func main() {
	var seconds int
	flag.IntVar(&seconds, "s", 5, "number of seconds the program runs. default: 10")
	flag.Parse()
	fmt.Println("Program finished in", seconds, "seconds.")

	mc := &MyChannel{ch: make(chan string)}

	wg := &sync.WaitGroup{}
	wg.Add(4) // 2 readera + 2 writera

	go writer(mc, wg) // запуск отправителей
	go writer(mc, wg)

	go reader(mc, wg) // запуск получателей
	go reader(mc, wg)

	// даём программе отработать указанное количество секунд
	time.Sleep(time.Second * time.Duration(seconds))
	fmt.Println("\n", "main: time is out")

	mc.SafeClose() // закрываем канал

	wg.Wait() // ждём завершения работы всех горутин
}

func writer(mc *MyChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("writer: start sending data in ch")
	for {
		if mc.isClosed {
			fmt.Println("writer: ch closed, exit")
			return
		}
		mc.ch <- "."
		time.Sleep(100 * time.Millisecond) // пауза 100 мс чтобы не заспамить консоль.
	}
}

func reader(mc *MyChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("reader: start reading ch")

	// range среагирует на закрытие канала и горутина завершится
	for data := range mc.ch {
		fmt.Print(data)
	}
	fmt.Println("reader: nothing to read, exit")
}
