package main

import (
	"fmt"
	"github.com/leonardchinonso/leon/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Leon programming language!\n", usr.Username)
	fmt.Printf("Type in a command to get started\n")
	repl.Start(os.Stdin, os.Stdout)
}
