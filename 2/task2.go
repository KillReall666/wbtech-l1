package main

import (
	"fmt"
	"sync"
)

func doubleNum(num int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := num * num
	resultChan <- result
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	//Канал для передачи результатов
	resultChan := make(chan int)

	//Используем ВейтГруп для ожидания завершения всех горутин
	var wg sync.WaitGroup

	//Запускаем горутины для расчета квадратов чисел
	for _, num := range nums {
		wg.Add(1)
		go doubleNum(num, resultChan, &wg)
	}

	//Запаускаем горутину для закрытия канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	//Читаем результаты из канала и выводим в стдаут
	for result := range resultChan {
		fmt.Println(result)
	}
}
