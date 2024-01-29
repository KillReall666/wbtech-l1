package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}

	fmt.Println(sl)
	sl = deleteAtIndex(sl, 2)
	fmt.Println(sl)
}

// slices.Delete()
func deleteAtIndex[T any](in []T, i int) []T {
	return append(in[:i], in[i+1:]...)
}
