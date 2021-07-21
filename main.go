package main

import "fmt"

type rectangle struct {
	a int
	b int
}

func main() {
	answer := rectangle{
		a: 2,
		b: 2,
	}
	fmt.Println(answer)
}

func S(rec rectangle) int {
	summaS := rec.a * rec.b
	return summaS
}

func P(rec rectangle) int {
	summaP := 2 * (rec.a * rec.b)
	return summaP
}