package engine

import (
	"regexp"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.ibm.com/sysflow/sf-processor/common/logger"
	"github.ibm.com/sysflow/sf-processor/core/sfpe/lang/parser"
)

// Parsed rule and filter object maps.
var rules = make(map[string]Rule)
var filters = make(map[string]Filter)

// Accessory parsing maps.
var lists = make(map[string][]string)
var macroCtxs = make(map[string]parser.IExpressionContext)

// Regular expression for pasting lists.
var itemsre = regexp.MustCompile(`(^\[)(.*)(\]$?)`)

// PolicyInterpreter defines a rules engine for SysFlow data streams.
type PolicyInterpreter struct{}

// Compile parses and interprets an input policy defined in path.
func (pi PolicyInterpreter) compile(path string) {
	// Setup the input
	is, err := antlr.NewFileStream(path)
	if err != nil {
		logger.Error.Println(err)
	}

	// Create the Lexer
	lexer := parser.NewSfplLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSfplParser(stream)

	// Parse the policy
	antlr.ParseTreeWalkerDefault.Walk(&sfplListener{}, p.Policy())
}

// Compile parses and interprets a set of input policies defined in paths.
func (pi PolicyInterpreter) Compile(paths ...string) {
	for _, path := range paths {
		logger.Trace.Println("Parsing policy file ", path)
		pi.compile(path)
	}
}

// Process executes all compiled policies against record r.
func (pi PolicyInterpreter) Process(applyFilters bool, r Record) (bool, []Rule) {
	var rlist []Rule
	match := false
	if applyFilters && pi.evalFilters(r) {
		return match, rlist
	}
	for _, rule := range rules {
		if rule.condition.Eval(r) {
			rule.ctx["record"] = r
			rlist = append(rlist, rule)
			match = true
		}
	}
	return match, rlist
}

// ProcessRule executes compiled policy rule p against record r.
func (pi PolicyInterpreter) ProcessRule(applyFilters bool, r Record, ruleNames ...string) (bool, []Rule) {
	var rlist []Rule
	match := false
	if applyFilters && pi.evalFilters(r) {
		return match, rlist
	}
	for _, rname := range ruleNames {
		if rule, ok := rules[rname]; ok && rule.condition.Eval(r) {
			rlist = append(rlist, rule)
			match = true
		}
	}
	return match, rlist
}

// EvalFilters executes compiled policy filters against record r.
func (pi PolicyInterpreter) evalFilters(r Record) bool {
	for _, f := range filters {
		if f.condition.Eval(r) {
			return true
		}
	}
	return false
}

type sfplListener struct {
	*parser.BaseSfplListener
}

// ExitList is called when production list is exited.
func (listener *sfplListener) ExitPlist(ctx *parser.PlistContext) {
	logger.Trace.Println("Parsing list ", ctx.GetText())
	lists[ctx.ID().GetText()] = listener.extractListFromItems(ctx.Items())
}

// ExitMacro is called when production macro is exited.
func (listener *sfplListener) ExitPmacro(ctx *parser.PmacroContext) {
	logger.Trace.Println("Parsing macro ", ctx.GetText())
	macroCtxs[ctx.ID().GetText()] = ctx.Expression()
}

// ExitFilter is called when production filter is exited.
func (listener *sfplListener) ExitPfilter(ctx *parser.PfilterContext) {
	logger.Trace.Println("Parsing filter ", ctx.GetText())
	f := Filter{
		name:      ctx.ID().GetText(),
		condition: listener.visitExpression(ctx.Expression()),
	}
	filters[f.name] = f
}

// ExitFilter is called when production filter is exited.
func (listener *sfplListener) ExitPrule(ctx *parser.PruleContext) {
	logger.Trace.Println("Parsing rule ", ctx.GetText())
	r := Rule{
		name:      listener.getOffChannelText(ctx.Text(0)),
		desc:      listener.getOffChannelText(ctx.Text(1)),
		condition: listener.visitExpression(ctx.Expression()),
		actions:   listener.getActions(ctx.Text(2).GetText()),
		tags:      listener.getTags(ctx.Items()),
		priority:  listener.getPriority(ctx.SEVERITY().GetText()),
		ctx:       make(map[string]interface{}),
	}
	rules[r.name] = r
}

func (listener *sfplListener) getOffChannelText(ctx parser.ITextContext) string {
	a := ctx.GetStart().GetStart()
	b := ctx.GetStop().GetStop()
	interval := antlr.Interval{Start: a, Stop: b}
	return ctx.GetStart().GetInputStream().GetTextFromInterval(&interval)
}

func (listener *sfplListener) getTags(ctx parser.IItemsContext) []EnrichmentTag {
	var tags = make([]EnrichmentTag, 0)
	return append(tags, listener.extractListFromItems(ctx))
}

func (listener *sfplListener) getPriority(p string) Priority {
	switch strings.ToLower(p) {
	case Low.String():
		return Low
	case Medium.String():
		return Medium
	case High.String():
		return High
	default:
		logger.Warn.Printf("Unrecognized priority value %s. Deferring to %s\n", p, Low.String())
		break
	}
	return Low
}

func (listener *sfplListener) getActions(astr string) []Action {
	var actions []Action
	l := listener.extractList(astr)
	for _, v := range l {
		switch strings.ToLower(v) {
		case Alert.String():
			actions = append(actions, Alert)
		case Tag.String():
			actions = append(actions, Tag)
		default:
			logger.Warn.Println("Unrecognized action value ", v)
			break
		}
	}
	return actions
}

func (listener *sfplListener) extractList(str string) []string {
	s := []string{}
	ls := strings.Split(itemsre.ReplaceAllString(str, "$2"), LISTSEP)
	for _, v := range ls {
		s = append(s, v)
	}
	return s
}

func (listener *sfplListener) extractListFromItems(ctx parser.IItemsContext) []string {
	return listener.extractList(ctx.GetText())
}

func (listener *sfplListener) extractListFromAtoms(ctxs []parser.IAtomContext) []string {
	s := []string{}
	for _, v := range ctxs {
		s = append(s, listener.reduceList(v.GetText())...)
	}
	return s
}

func (listener *sfplListener) reduceList(sl string) []string {
	s := []string{}
	if l, ok := lists[sl]; ok {
		for _, v := range l {
			s = append(s, listener.reduceList(v)...)
		}
	} else {
		s = append(s, sl)
	}
	return s
}

func (listener *sfplListener) visitExpression(ctx parser.IExpressionContext) Criterion {
	orCtx := ctx.GetChild(0).(parser.IOr_expressionContext)
	orPreds := make([]Criterion, 0)
	for _, andCtx := range orCtx.GetChildren() {
		if andCtx.GetChildCount() > 0 {
			andPreds := make([]Criterion, 0)
			for _, termCtx := range andCtx.GetChildren() {
				t, isTermCtx := termCtx.(parser.ITermContext)
				if isTermCtx {
					c := listener.visitTerm(t)
					andPreds = append(andPreds, c)
				}
			}
			orPreds = append(orPreds, All(andPreds))
		}
	}
	return Any(orPreds)
}

func (listener *sfplListener) visitTerm(ctx parser.ITermContext) Criterion {
	termCtx := ctx.(*parser.TermContext)
	if termCtx.Variable() != nil {
		if m, ok := macroCtxs[termCtx.GetText()]; ok {
			return listener.visitExpression(m)
		}
		logger.Error.Println("Unrecognized reference ", termCtx.GetText())
	} else if termCtx.NOT() != nil {
		return listener.visitTerm(termCtx.GetChild(1).(parser.ITermContext)).Not()
	} else if opCtx, ok := termCtx.Unary_operator().(*parser.Unary_operatorContext); ok {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		if opCtx.EXISTS() != nil {
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
		return listener.visitExpression(termCtx.Expression())
	} else if termCtx.IN() != nil {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		rop := termCtx.AllAtom()[1:]
		return In(lop, listener.extractListFromAtoms(rop))
	} else if termCtx.PMATCH() != nil {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		rop := termCtx.AllAtom()[1:]
		return PMatch(lop, listener.extractListFromAtoms(rop))
	} else {
		logger.Warn.Println("Unrecognized term ", termCtx.GetText())
	}
	return False
}
