package repl

import (
    "bufio"
    "fmt"
    "io"
    "monkey/lexer"
    "monkey/parser"
    "monkey/evaluator"
    "monkey/object"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    env := object.NewEnvironment()
    macroEnv := object.NewEnvironment()

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

        evaluator.DefineMacros(program, macroEnv)
        expanded := evaluator.ExpandMacros(program, macroEnv)

        evaluated := evaluator.Eval(expanded, env)
        if evaluated != nil {
            io.WriteString(out, evaluated.Inspect())
            io.WriteString(out, "\n")
        }
    }
}

func printParserErrors(out io.Writer, errors []string) {
    io.WriteString(out, "oops! we ran into an error here!\n")
    io.WriteString(out, " parser errrs: \n")
    for _, msg := range errors {
        io.WriteString(out, "\t" + msg + "\n")
    }
}

