package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) Say(word string) {
	fmt.Println(word)
}

type Action struct {
	h Human
}

func (a *Action) Jump() {
	fmt.Println("jump jump jump")
	a.h.Say("Hello")
}

func main() {
	var a Action
	a.h.Say("Hello!")
}
