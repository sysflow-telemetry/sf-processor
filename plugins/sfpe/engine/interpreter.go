package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.ibm.com/sysflow/sf-processor/plugins/sfpe/lang/parser"
)

type sfplListener struct {
	*parser.BaseSfplListener
}

// ExitF_list is called when production f_list is exited.
func (s *sfplListener) ExitF_list(ctx *parser.F_listContext) {
	fmt.Println("Exiting list parsing")
	fmt.Println(ctx.GetText())
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("- list: test\nitems: [1, 2, 3]")

	// Create the Lexer
	lexer := parser.NewSfplLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSfplParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&sfplListener{}, p.Policy())
	//fmt.Println(p.Policy().GetText())
}
