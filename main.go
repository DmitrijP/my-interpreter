// main.go

package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/DmitrijP/my-interpreter/ball/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Ball programmin language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
