package main

import (
	"fmt"
	"os"

	"github.com/ranganath42/rila/repl"
)

func main() {
	fmt.Printf("Rila REPL\n")
	repl.Start(os.Stdin, os.Stdout)

}
