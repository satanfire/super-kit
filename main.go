package main

import (
	"fmt"

	"github.com/satanfire/super-kit/cmd"
)

func main() {
	if err := cmd.Exec(); err != nil {
		fmt.Printf("main, cmd.Exec failed, %s\n", err.Error())
	}
}
