package main

import (
	"dalton.dog/termina/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Termina REPL!\n", user.Username)
	repl.StartREPL(os.Stdin, os.Stdout)
}
