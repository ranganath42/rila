package repl

import (
	"bufio"
	"fmt"
	"github.com/ranganath42/rila/lexer"
	"github.com/ranganath42/rila/token"
	"io"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, ">> ")
		scanned := scanner.Scan()
		if !scanned || isExitCommand(scanner.Text()) {
			return
		}
		line := scanner.Text()
		lex := lexer.New(line)
		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Fprintf(out, "%-10s %s\n", tok.Type, tok.Literal)
		}
	}
}

func isExitCommand(cmd string) bool {
	for _, exitCmd := range []string{"exit", "quit", "q", "bye"} {
		if cmd == exitCmd {
			return true
		}
	}
	return false
}
