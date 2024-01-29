package main

import (
	"fmt"
	"slices"
)

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println("начальный слайс:", input)

	fmt.Println("setViaMap:", setViaMap(input))

	fmt.Println("setViaSortedSlice:", setViaSortedSlice(input))
}

func setViaMap(in []string) map[string]struct{} {
	m := make(map[string]struct{}, len(in))
	for _, i := range in {
		m[i] = struct{}{}
	}
	return m

}

func setViaSortedSlice(in []string) []string {
	slices.Sort(in)

	out := make([]string, 0, len(in))
	var prev string
	for _, cur := range in {
		if cur != prev {
			out = append(out, cur)
			prev = cur
		}
	}
	// обрезаю capacity до получившегося количества элементов.
	return out[:len(out):len(out)]
}
