package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/masonictemple4/boss/evaluator"
	"github.com/masonictemple4/boss/lexer"
	"github.com/masonictemple4/boss/object"
	"github.com/masonictemple4/boss/parser"
)

const PROMPT = "$ "

// TODO: Find better art for this.
const LOGO = `
  . .. .  . .  . .. .  . .. . ..&@  . .. .  . .  . .. .  . .  . .. .  . .  . .. 
.. .@@@@@. . #@@@@ . .. .,@@@   . .. @@@  .  .  @@@  . .. @@%. . @@@@. .  . @@ .
.. .@@@@@. .  @@@@@. .. @@@@ .  . .. .@@@@.  .@@@@@  . .. . %. .@@@@ . .  . .@ .
.. .@@@@@. .  @@@@@. ..@@@@@ .  . .. .@@@@@  .@@@@@@@. .. . .. .@@@@@@@.  . .  .
.. .@@@@@.  @@@%.  . .@@@@@@ .  . .. .@@@@@/ .  @@@@@@@@@@. .. . &@@@@@@@@@@.  .
.. .@@@@@. .  @@@@@. . @@@@@ .  . .. .@@@@@  . .. . @@@@@@@@@. .  . . @@@@@@@@@.
.. .@@@@@. .  .@@@@@ ..@@@@@ .  . .. .@@@@@  .@.. .  . .@@@@@. .@ . .. .  @@@@@.
.. .@@@@@. .  .@@@@# .. @@@@@.  . .. @@@@@.  .@@. .  . ..@@@@. .@ . .. .  .@@@@.
.. .@@@@@. . @@@@  . .. .  @@@  . .&@@@ . .  .@@@@.  . .@@@ .. .@@@*.. . @@@   .
.. .  . .. .  . .  . .. .  . .  . .. .  . .  . .. .  . .. . .. .  . .. .  . .  .
`

// CHALLENGE: Add multi line support to the REPL.
// To enable function definitions, etc..
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evalutated := evaluator.Eval(program, env)
		if evalutated != nil {
			io.WriteString(out, evalutated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, LOGO)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
