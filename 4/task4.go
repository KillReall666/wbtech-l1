package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type myChan chan struct{}

func main() {

	ch := make(myChan, 1)

	var num int
	// парсинг количества воркеров из аргумента командной строки -w
	flag.IntVar(&num, "w", 10, "Enter number of workers.")
	flag.Parse()

	// переменная context получит уведомление от операционной системы
	// о нажатии ctrl+c пользователем. Этот сигнал далее передаётся всем,
	// кто подписан на данный контекст.
	ctx, _ := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	fmt.Println("The program will end by pressing Ctrl+c")

	//Создание указанного пользователем количества воркеров
	//Для каждого воркера запускается функция reader(c) в отдельной горутине
	for i := 0; i < num; i++ {
		go reader(ch, i)
	}

	//запускает в новой горутине функцию writer, которая
	//будет записывать данные в канал
	writer(ch, ctx)
}

func writer(ch myChan, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\n writer: Graceful Shutdown by Ctrl+c")

			// закрытие канала остановит циклы for-range во всех reader'ах,
			// благодаря чему все читающие горутины сэйфово завершатся без утечек памяти.
			close(ch)
			return
		default:
			time.Sleep(100 * time.Millisecond) // пауза
			ch <- struct{}{}                   // отправка пустоты в канал
		}
	}
}

func reader(ch myChan, id int) {
	// for i := range ch {
	for range ch {
		fmt.Print(id, " ") // принт в консоль id'шника горутины
	}
}
