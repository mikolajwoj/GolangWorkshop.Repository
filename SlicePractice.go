package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	//Creating slice of x variables
	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)
	}
	for y := range p {
		for x := range p[y] {
			p[x][y] = uint8(x * y)
		}
	}
	return p
}

func SlicesPractice() {
	pic.Show(Pic)
}
