package main

import (
	"gordle/gordle"
	"os"
)

func main() {
	g := gordle.New(os.Stdin, "hello", 5)
	g.Play()
}
