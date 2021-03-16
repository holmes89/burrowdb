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

		p := lib.NewParser(lib.NewLexer(strings.NewReader(line)))

		exp, err := p.Parse()
		if err != nil {
			fmt.Fprintln(out, err.Error())
		}

		fmt.Fprintf(out, "%+v\n", exp)
	}
}

const (
	PROMPT = "> "
)
