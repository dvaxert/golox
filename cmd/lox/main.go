package main

import (
	"fmt"
	"os"

	"github.com/dvaxert/golox/pkg/lox"
)

func main() {
	switch len(os.Args) {
	case 1:
		lox.RunPromt()
	case 2:
		lox.RunFile(os.Args[1])
	default:
		fmt.Println("Usage golox [script]")
	}
}
