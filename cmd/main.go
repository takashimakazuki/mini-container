package main

import (
	"os"
	minicontainer "takashimakazuki/mini-container/pkg"
)

func main() {
	switch os.Args[1] {
	case "run":
		minicontainer.Run()
	case "child":
		minicontainer.Child()
	default:
		panic("help")
	}
}
