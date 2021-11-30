package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/yugovtr/markmoji/emoji"
)

func main() {
	var input []byte
	var err error
	args := os.Args[1:]

	input = []byte(strings.Join(args, "\n"))

	if len(input) == 0 {
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	}

	var output string = string(input)

	for name, emoji := range emoji.Emojis {
		rgx := fmt.Sprintf(":%s:", name)
		output = strings.ReplaceAll(output, rgx, string(emoji.Code))
	}

	fmt.Fprintf(os.Stdout, "%s\n", output)
}
