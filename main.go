package main

import (
	"fmt"

	"github.com/melkeydev/go-blueprint/cmd"
)

func main() {
	cmd.Execute()
	fmt.Println("hello CI")
}
