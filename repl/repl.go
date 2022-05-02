package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/leonardchinonso/leon/lexer"
	"github.com/leonardchinonso/leon/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		ln := scanner.Text()
		l := lexer.NewLexer(ln)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
