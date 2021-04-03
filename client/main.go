package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/holmes89/burrowdb/client/lib"
	"github.com/holmes89/burrowdb/node"
)

func main() {
	// replaced with connection
	db, err := bolt.Open("example.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	eval := lib.NewEvaluator(node.NewNodeRepo(db))
	defer db.Close()

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

		if err := eval.Eval(exp); err != nil {
			fmt.Fprintln(out, err.Error())
		} else {
			fmt.Fprintln(out, "SUCCESS")
		}

	}
}

const (
	PROMPT = "> "
)
