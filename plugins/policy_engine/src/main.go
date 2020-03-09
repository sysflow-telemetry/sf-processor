package main

import (
	"fmt"

	"../lang/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type sfplListener struct {
	*parser.BaseSfplListener
}

// ExitF_list is called when production f_list is exited.
func (s *BaseSfplListener) ExitF_list(ctx *F_listContext) {
	fmt.Println(ctx)
	ctx.Println()
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("- list: testlist\nitems: [1, 2, 3]")

	// Create the Lexer
	lexer := parser.NewSfplLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSfplParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&sfplListener{}, p.Start())
}
