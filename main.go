package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"flag"

	"github.com/yugovtr/markmoji/emoji"
)

func read(in chan []byte) {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	in <- input
}

func main() {
	var l = flag.Bool("l", false, "list all emoji")
	flag.Parse()

	if *l {
		emoji.Update(os.Stdout)
		return
	}

	var (
		input   []byte
		timeout time.Duration = 50
	)

	args := os.Args[1:]
	input = []byte(strings.Join(args, "\n"))

	if len(input) == 0 {
		in := make(chan []byte, 1)
		go read(in)

		select {
		case i := <-in:
			input = i
		case <-time.After(timeout * time.Millisecond):
			return
		}
	}

	output := string(input)
	for name, emoji := range emoji.Emojis {
		rgx := fmt.Sprintf(":%s:", name)
		output = strings.ReplaceAll(output, rgx, string(emoji.Code))
	}

	fmt.Fprintf(os.Stdout, "%s\n", output)
}
