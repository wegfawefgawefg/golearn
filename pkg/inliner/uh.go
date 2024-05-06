package main

type A struct {
	A string
	B string
}

type B struct {
	A string
	B string
}

type C struct {
	A string
	B string
}

// function that takes in an a and a b and returns a string
func f(a A, b B) string {
	return a.A + b.A
}

// func that takes in some strings and returns a c
func g(a string, b string) C {
	return C{A: a, B: b}
}

func try() {
	// create an a and a b
	a := A{A: "a", B: "b"}
	b := B{A: "c", B: "d"}

	k := f(a, b)
	a_c := g("e", "f")

	println(k)
	println(a_c.A)
}
