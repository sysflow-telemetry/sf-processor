package engine

import (
	"regexp"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"github.ibm.com/sysflow/sf-processor/plugins/sfpe/lang/parser"
)

var lists = make(map[string][]string)
var macros = make(map[string]parser.IExpressionContext)
var filters = make(map[string]parser.IExpressionContext)
var rules = make(map[string]parser.IPruleContext)

var itemsre = regexp.MustCompile(`(^\[)(.*)(\]$?)`)

type sfplListener struct {
	*parser.BaseSfplListener
}

// ExitList is called when production list is exited.
func (s *sfplListener) ExitPlist(ctx *parser.PlistContext) {
	logger.Trace.Printf("Parsing list %s", ctx.GetText())
	lists[ctx.ID().GetText()] = extractList(ctx.Items())
}

// ExitMacro is called when production macro is exited.
func (s *sfplListener) ExitPmacro(ctx *parser.PmacroContext) {
	logger.Trace.Printf("Parsing macro %s", ctx.GetText())
	macros[ctx.ID().GetText()] = ctx.Expression()
}

// ExitFilter is called when production filter is exited.
func (s *sfplListener) ExitPfilter(ctx *parser.PfilterContext) {
	logger.Trace.Printf("Parsing filter %s", ctx.GetText())
	filters[ctx.ID().GetText()] = ctx.Expression()
}

// ExitFilter is called when production filter is exited.
func (s *sfplListener) ExitPrule(ctx *parser.PruleContext) {
	logger.Trace.Printf("Parsing macro %s", ctx.GetText())

}

// Compile parses and interprets an input policy defined in path.
func Compile(path string) {
	// Setup the input
	is, _ := antlr.NewFileStream(path)

	// Create the Lexer
	lexer := parser.NewSfplLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSfplParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&sfplListener{}, p.Policy())
}

func extractList(ctx parser.IItemsContext) []string {
	s := []string{}
	ls := strings.Split(itemsre.ReplaceAllString(ctx.GetText(), "$2"), LISTSEP)
	for _, v := range ls {
		s = append(s, v)
	}
	return s
}

//func visitExpression(ctx parser.ExpressionContext)
