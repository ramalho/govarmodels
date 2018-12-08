package main

import (
	"fmt"
)

// Ponto no plano
type Ponto struct {
	x, y float64
}

func main() {
	i := 3
	i2 := i
	i2++
	fmt.Printf("i\t%#v\ni2\t%#v\n\n", i, i2)

	a := [...]int{1, 2, 3}
	a2 := a
	a2[0]++
	fmt.Printf("a\t%#v\na2\t%#v\n\n", a, a2)

	p := Ponto{2, 3}
	p2 := p
	p2.y++
	fmt.Printf("p\t%#v\np2\t%#v\n\n", p, p2)

	pp := &Ponto{2, 3}
	pp2 := pp
	pp2.y++
	fmt.Printf("pp\t%#v\npp2\t%#v\n\n", pp, pp2)

	s := []int{1, 2, 3}
	s2 := s
	s2[0]++
	fmt.Printf("s\t%#v\ns2\t%#v\n\n", s, s2)

	m := map[byte]int{1: 1, 2: 2, 3: 3}
	m2 := m
	m2[3]++
	fmt.Printf("m\t%#v\nm2\t%#v\n\n", m, m2)

}
