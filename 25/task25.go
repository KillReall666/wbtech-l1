package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Before sleep")
	sleep(5) // Вызов нашей пользовательской функции sleep на 5 секунд
	fmt.Println("After sleep")
}

