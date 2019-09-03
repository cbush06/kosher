package main

import (
	"github.com/cbush06/kosher/cmd"
)

//go:generate go run gen/axecore_gen.go

func main() {
	cmd.Init()
}
