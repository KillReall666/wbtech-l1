package main

import "fmt"

func main() {
	arr := []int{3, 2, 1, 4, 5, 6, 4, 8, 9, 10, 11, 21, 2, 13, 13, 3, 24, 3, 1, 22, 3, 2, 43, 5, 64, 6, 7, 5}
	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) < 2 { //так как функция рекурсивная обязательна проверка на то что длина массива больше 2
		return arr
	}

	left := 0             // определяем начало массива
	right := len(arr) - 1 // определяем конец массива
	target := 0           // определяем фиксированную точку

	// меняем фиксированную точку и конец массива местами
	arr[target], arr[right] = arr[right], arr[target]

	// проходим по массиву arr
	//мы меняем его местами с элементом, на который указывает left, и затем увеличиваем left
	//В результате получается, что слева от target находятся все элементы меньше его,
	//а справа - больше или равны.
	for i, _ := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[right], arr[left] = arr[left], arr[right]

	//Функция рекурсивно вызывается для левого и правого подмассивов.
	//Они не включают опорный элемент, поскольку он уже находится на правильном месте.
	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}
