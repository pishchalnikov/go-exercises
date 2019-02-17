package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)

	for x := 0; x < dy; x++ {
		res[x] = make([]uint8, dx)

		for y := 0; y < dx; y++ {
			res[x][y] = uint8(x ^ y)
		}
	}

	return res
}

func main() {
	pic.Show(Pic)
}
