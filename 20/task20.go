package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Type some words for revert: ")

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// убираем символы каретки
	str := strings.ReplaceAll(input, "\r\n", "")
	fmt.Printf("%q \n", str)

	sliceOfStrings := strings.Split(str, " ")

	slices.Reverse(sliceOfStrings)

	fmt.Println(strings.Join(sliceOfStrings, " "))
}
