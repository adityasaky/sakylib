package sakylib

import "fmt"

func Function1() {
	fmt.Println("This is a function. 1 indeed!")
}

func Function2(s string) {
	fmt.Println("This is a function. 2 indeed!", s)
}

type Type struct {
	Field string
}

func (t *Type) Method() {
	fmt.Println("This is a method.")
}
