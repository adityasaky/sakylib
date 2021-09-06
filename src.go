package sakylib

import "fmt"

func Function1() {
	fmt.Println("This is a function. 1 indeed!")
}

func Function2(s string) {
	fmt.Println("This is a function. 2 indeed!", s)
}

type Type1 struct {
	Field string
}

func (t *Type1) Method(s string) {
	fmt.Println("This is a method.")
	t.Field = s
}

type Type2 struct {
	Field string
}

func (t Type2) Method(s string) {
	fmt.Println("This is a method.")
	t.Field = s
}
