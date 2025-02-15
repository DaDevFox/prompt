package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/DaDevFox/prompt"
	"github.com/DaDevFox/prompt/input"
)

func main() {
	val, err := prompt.New().Ask("Input an integer:").
		Input("", input.WithInputMode(input.InputInteger))
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	fmt.Println(val)
}
