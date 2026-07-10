package loop

import "fmt"

func MakeSquare(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print("X")
		}
		fmt.Print("\n")
	}
}

func MakeRectangle(h int, w int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Print("X")
		}
		fmt.Print("\n")
	}
}

func MakeTriangle(h int) {
	w := h / 2
	for i := 0; i < h; i++ {
		for j := -h; j < h+i; j++ {
			if j == w || j == w+i || j == w-i {
				fmt.Print("X")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}
