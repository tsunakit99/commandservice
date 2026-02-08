package main

import (
	"github.com/tsunakit99/commandservice/command/presen"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		presen.CommandDepend,
	).Run()
}
