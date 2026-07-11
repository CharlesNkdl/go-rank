package app

import (
	"fmt"

	"github.com/cnkdl/go-rank/internal/array"
	"github.com/cnkdl/go-rank/internal/loop"
)

type App struct {
	name string
}

func CreateApp() App {
	return App{name: "go-rank"}
}

func Run() error {
	//app := CreateApp()
	w := 10
	h := 20
	loop.MakeSquare(w)
	fmt.Println()
	loop.MakeRectangle(h, w)
	fmt.Println()
	loop.MakeTriangle(w)

	loop.Sheep(2, 100, 1, 300)

	array.TicTacToe(h)

	return nil
}
