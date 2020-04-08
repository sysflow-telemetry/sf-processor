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
	lists[ctx.ID().GetText()] = extractListFromItems(ctx.Items())
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
	logger.Trace.Printf("Parsing rule %s", ctx.GetText())
	visitExpression(ctx.Expression())

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

func extractListFromItems(ctx parser.IItemsContext) []string {
	s := []string{}
	ls := strings.Split(itemsre.ReplaceAllString(ctx.GetText(), "$2"), LISTSEP)
	for _, v := range ls {
		s = append(s, v)
	}
	return s
}

func extractListFromAtoms(ctxs []parser.IAtomContext) []string {
	s := []string{}
	for _, v := range ctxs {
		s = append(s, reduceList(v.GetText())...)
	}
	return s
}

func reduceList(sl string) []string {
	s := []string{}
	if l, ok := lists[sl]; ok {
		for _, v := range l {
			s = append(s, reduceList(v)...)
		}
	} else {
		s = append(s, sl)
	}
	return s
}

func visitExpression(ctx parser.IExpressionContext) Criterion {
	orCtx := ctx.GetChild(0).(parser.IOr_expressionContext)
	orPreds := make([]Criterion, 0)
	for _, andCtx := range orCtx.GetChildren() {
		if andCtx.GetChildCount() > 0 {
			andPreds := make([]Criterion, 0)
			for _, termCtx := range andCtx.GetChildren() {
				t, isTermCtx := termCtx.(parser.ITermContext)
				if isTermCtx {
					c := visitTerm(t)
					andPreds = append(andPreds, c)
				}
			}
			orPreds = append(orPreds, All(andPreds))
		}
	}
	return Any(orPreds)
}

func visitTerm(ctx parser.ITermContext) Criterion {
	termCtx := ctx.(*parser.TermContext)
	if termCtx.Variable() != nil {
		if m, ok := macros[termCtx.GetText()]; ok {
			return visitExpression(m)
		}
		logger.Error.Println("Unrecognized reference ", termCtx.GetText())
	} else if termCtx.NOT() != nil {
		return visitTerm(termCtx.GetChild(1).(parser.ITermContext)).Not()
	} else if opCtx, ok := termCtx.Unary_operator().(*parser.Unary_operatorContext); ok {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		if opCtx.EXISTS() != nil {
			logger.Trace.Println("Exists: ", lop)
			return Exists(lop)
		}
		logger.Error.Println("Unrecognized unary operator ", opCtx.GetText())
	} else if opCtx, ok := termCtx.Binary_operator().(*parser.Binary_operatorContext); ok {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		rop := termCtx.Atom(1).(*parser.AtomContext).GetText()
		if opCtx.CONTAINS() != nil {
			return Contains(lop, rop)
		} else if opCtx.ICONTAINS() != nil {
			return IContains(lop, rop)
		} else if opCtx.STARTSWITH() != nil {
			return StartsWith(lop, rop)
		} else if opCtx.EQ() != nil {
			return Eq(lop, rop)
		} else if opCtx.NEQ() != nil {
			return NEq(lop, rop)
		} else if opCtx.GT() != nil {
			return Gt(lop, rop)
		} else if opCtx.GE() != nil {
			return Ge(lop, rop)
		} else if opCtx.LT() != nil {
			return Lt(lop, rop)
		}
		logger.Error.Println("Unrecognized binary operator ", opCtx.GetText())
	} else if termCtx.Expression() != nil {
		return visitExpression(termCtx.Expression())
	} else if termCtx.IN() != nil {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		rop := termCtx.AllAtom()[1:]
		return In(lop, extractListFromAtoms(rop))
	} else if termCtx.PMATCH() != nil {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		rop := termCtx.AllAtom()[1:]
		return PMatch(lop, extractListFromAtoms(rop))
	} else {
		logger.Warn.Println("Unrecognized term ", termCtx.GetText())
	}
	return False
}
