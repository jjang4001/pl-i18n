package main

import (
	"fmt"
	"os"
	"os/user"
	"pl-i18n/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is an interpreter for pl-i18n!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
