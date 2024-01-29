package main

import (
	"fmt"
	"reflect"
)

type printer interface {
	Print()
}

type example struct{}

func (e example) Print() {
	fmt.Println("Check")
}

func main() {
	ch := make(chan int)
	check := example{}
	var st = printer(check)
	arr := []any{1, "string", true, 0.5, ch, check, st}

	for _, i := range arr {
		switch i.(type) {
		case int:
			fmt.Println("Int")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Bool")
		case float64:
			fmt.Println("Float64")
		default:
			fmt.Println("unexpected type:", reflect.TypeOf(i))
		}
	}
}
