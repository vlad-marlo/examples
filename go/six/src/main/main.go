package main

import (
	"github.com/vlad-marlo/example/go/modules/src/internal/controller"
)

func main() {
	ctrl := controller.New()
	if err := ctrl.Start(); err != nil {
		panic(err)
	}
}
