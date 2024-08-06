package repl

import (
	"bufio"
	"fmt"
	"io"

	"dalton.dog/termina/lexer"
	"dalton.dog/termina/token"
)

const PROMPT = ">> "

func StartREPL(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.MakeNewLexer(line)
		for t := l.GetNextToken(); t.Type != token.EOF; t = l.GetNextToken() {
			fmt.Fprintf(out, "%+v\n", t)
		}
	}
}
