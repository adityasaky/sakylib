package sakylib

import "fmt"

func Function1() {
	fmt.Println("This is a function. 1 indeed!")
}

func Function2(s string) {
	fmt.Println("This is a function. 2 indeed!", s)
}

func Function3(s *string) {
	fmt.Println("This is a function. 3 indeed!", *s)
}

func Function4(t *Type1) {
	fmt.Println("This is a function. 4 indeed!")
	fmt.Println(t.Field)
}

func Function5(t Type2) {
	fmt.Println("This is a function. 5 indeed!")
	fmt.Println(t.Field)
}

func Function6(s string) string {
	fmt.Println("This is a function. 6 indeed!")
	return s
}

func Function7(fn func(int) string) {
	fmt.Println("This is a function. 7 indeed!")
	fmt.Println(fn(10))
}

func Function8(fn func(int) *Type1) {
	fmt.Println("This is a function. 8 indeed!")
	fmt.Println(fn(10))
}

func Function9(fn func(int) Type1) {
	fmt.Println("This is a function. 9 indeed!")
	fmt.Println(fn(10))
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
