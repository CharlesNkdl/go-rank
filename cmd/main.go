package main

import (
	"log"

	"github.com/cnkdl/go-rank/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalf("%s", err)
	}
	return
}
