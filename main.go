package main

import (
	"os"
	"fmt"
	"os/user"
	"github.com/khofesh/ninkasi/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Ninkasi programming language!\n", user.Username)
	fmt.Print(".help won't help you. We don't have enough information\n")
	repl.Start(os.Stdin, os.Stdout)
}