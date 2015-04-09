package main

import (
	"bufio"
	"fmt"
	c "github.com/funkygao/golib/color"
	"github.com/funkygao/golib/io"
	_io "io"
	"os"
	"strings"
)

var (
	search string
	color  string = "red"

	colorTable = map[string]func(string) string{
		"red":    c.Red,
		"blue":   c.Blue,
		"green":  c.Green,
		"yellow": c.Yellow,
	}
)

func init() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s text [color]\ncolor: red | blue | green | yellow\n",
			os.Args[0])
		os.Exit(0)
	}

	search = os.Args[1]
	if len(os.Args) > 2 {
		color = os.Args[2]
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := io.ReadLine(reader)
		if err != nil {
			if err != _io.EOF {
				fmt.Println(err)
			}

			break
		}

		s := string(line)
		if strings.Contains(s, search) {
			start := strings.Index(s, search)
			text := colored(color, search)
			fmt.Printf("%s%s%s\n", s[:start], text, s[start+len(search):])
		} else {
			fmt.Println(s)
		}
	}

}

func colored(c, s string) string {
	if f, present := colorTable[c]; present {
		return f(s)
	}

	// invalid color
	return s
}
