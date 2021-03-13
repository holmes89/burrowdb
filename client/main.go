package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/holmes89/burrowdb/client/lib"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	out := os.Stdout
	var line string
	for {

		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line = scanner.Text()

		p := lib.NewLexer(strings.NewReader(line))

		for {
			n := p.NextToken()
			if n == nil || n.IsEOF() {
				break
			}
			fmt.Printf("%+v\n", n)
		}
	}
}

const (
	PROMPT = "> "
)
