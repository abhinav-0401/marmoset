package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/abhinav-0401/marmoset/lexer"
	"github.com/abhinav-0401/marmoset/parser"
	// "github.com/abhinav-0401/marmoset/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		parser := parser.New(l)
		program := parser.ParseProgram()

		programPretty, _ := json.MarshalIndent(program, "", "    ")
		fmt.Printf("%+v\n", string(programPretty))

		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Printf("%+v\n", tok)
		// }
	}
}
