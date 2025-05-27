package main

import (
	"fmt"
	"github.com/abhinav-0401/marmoset/repl"
	"os"
)

func main() {
	fmt.Printf("Hello! This is the Monkey programming language!\n")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
