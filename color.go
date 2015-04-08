package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/funkygao/golib/color"
	"github.com/funkygao/golib/io"
	"os"
	"strings"
)

var (
	search  string
	colored string
)

func init() {
	flag.StringVar(&search, "s", "", "search text")
	flag.StringVar(&colored, "c", "", "color")
	flag.Parse()
	if search == "" {
		fmt.Println("search and color required")
		os.Exit(0)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := io.ReadLine(reader)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println(err)
			}

			os.Exit(1)
		}

		s := string(line)
		if strings.Contains(s, search) {
			start := strings.Index(s, search)
			text := color.Red(search)
			fmt.Printf("%s%s%s\n", s[:start], text, s[start+len(search):])
		}
	}

}
