//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package engine

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/lang/errorhandler"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/lang/parser"
)

// Parsed rule and filter object maps.
var rules = make([]Rule, 0)
var filters = make([]Filter, 0)

// Accessory parsing maps.
var lists = make(map[string][]string)
var macroCtxs = make(map[string]parser.IExpressionContext)

// Regular expression for pasting lists.
var itemsre = regexp.MustCompile(`(^\[)(.*)(\]$?)`)

// PolicyInterpreter defines a rules engine for SysFlow data streams.
type PolicyInterpreter struct {
	ahdl ActionHandler
}

// NewPolicyInterpreter constructs a new interpreter instance.
func NewPolicyInterpreter(conf Config) *PolicyInterpreter {
	ah := NewActionHandler(conf)
	return &PolicyInterpreter{ah}
}

// Compile parses and interprets an input policy defined in path.
func (pi *PolicyInterpreter) compile(path string) error {
	// Setup the input
	is, err := antlr.NewFileStream(path)
	if err != nil {
		logger.Error.Println("Error reading policy from path", path)
		return err
	}

	// Create the Lexer
	lexerErrors := &errorhandler.SfplErrorListener{}
	lexer := parser.NewSfplLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	parserErrors := &errorhandler.SfplErrorListener{}
	p := parser.NewSfplParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	// Pre-processing (to deal with usage before definitions of macros and lists)
	antlr.ParseTreeWalkerDefault.Walk(&sfplListener{}, p.Defs())
	p.GetInputStream().Seek(0)

	// Parse the policy
	antlr.ParseTreeWalkerDefault.Walk(&sfplListener{}, p.Policy())

	errFound := false
	if len(lexerErrors.Errors) > 0 {
		logger.Error.Printf("Lexer %d errors found\n", len(lexerErrors.Errors))
		for _, e := range lexerErrors.Errors {
			logger.Error.Println("\t", e.Error())
		}
		errFound = true
	}
	if len(parserErrors.Errors) > 0 {
		logger.Error.Printf("Parser %d errors found\n", len(parserErrors.Errors))
		for _, e := range parserErrors.Errors {
			logger.Error.Println("\t", e.Error())
		}
		errFound = true
	}

	if errFound {
		return errors.New("errors found during compilation of policies. check logs for detail.")
	}

	return nil
}

// Compile parses and interprets a set of input policies defined in paths.
func (pi *PolicyInterpreter) Compile(paths ...string) error {
	for _, path := range paths {
		logger.Trace.Println("Parsing policy file ", path)
		if err := pi.compile(path); err != nil {
			return err
		}
	}
	return nil
}

// ProcessAsync executes all compiled policies against record r.
func (pi *PolicyInterpreter) ProcessAsync(applyFilters bool, filterOnly bool, r *Record, out func(r *Record)) {
	if applyFilters && pi.EvalFilters(r) {
		return
	}
	if filterOnly {
		out(r)
	}
	match := false
	for _, rule := range rules {
		if rule.Enabled && rule.isApplicable(r) && rule.condition.Eval(r) {
			pi.ahdl.HandleActionAsync(rule, r, out)
			match = true
		}
	}
	if match {
		out(r)
	}
}

// Process executes all compiled policies against record r.
func (pi *PolicyInterpreter) Process(applyFilters bool, filterOnly bool, r *Record) (bool, *Record) {
	match := false
	if applyFilters && pi.EvalFilters(r) {
		return match, nil
	}
	if filterOnly {
		return true, r
	}
	for _, rule := range rules {
		if rule.Enabled && rule.isApplicable(r) && rule.condition.Eval(r) {
			pi.ahdl.HandleAction(rule, r)
			match = true
		}
	}
	return match, r
}

// EvalFilters executes compiled policy filters against record r.
func (pi *PolicyInterpreter) EvalFilters(r *Record) bool {
	for _, f := range filters {
		if f.Enabled && f.condition.Eval(r) {
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
		Name:      ctx.ID().GetText(),
		condition: listener.visitExpression(ctx.Expression()),
		Enabled:   ctx.ENABLED() == nil || listener.getEnabledFlag(ctx.Enabled()),
	}
	filters = append(filters, f)
}

// ExitFilter is called when production filter is exited.
func (listener *sfplListener) ExitPrule(ctx *parser.PruleContext) {
	logger.Trace.Println("Parsing rule ", ctx.GetText())
	r := Rule{
		Name:      listener.getOffChannelText(ctx.Text(0)),
		Desc:      listener.getOffChannelText(ctx.Text(1)),
		condition: listener.visitExpression(ctx.Expression()),
		Actions:   listener.getActions(ctx),
		Tags:      listener.getTags(ctx),
		Priority:  listener.getPriority(ctx),
		Prefilter: listener.getPrefilter(ctx),
		Enabled:   ctx.ENABLED(0) == nil || listener.getEnabledFlag(ctx.Enabled(0)),
	}
	rules = append(rules, r)
}

func (listener *sfplListener) getEnabledFlag(ctx parser.IEnabledContext) bool {
	flag := trimBoundingQuotes(ctx.GetText())
	if b, err := strconv.ParseBool(flag); err == nil {
		return b
	}
	logger.Warn.Println("Unrecognized enabled flag: ", flag)
	return true
}

func (listener *sfplListener) getOffChannelText(ctx parser.ITextContext) string {
	a := ctx.GetStart().GetStart()
	b := ctx.GetStop().GetStop()
	interval := antlr.Interval{Start: a, Stop: b}
	return ctx.GetStart().GetInputStream().GetTextFromInterval(&interval)
}

func (listener *sfplListener) getTags(ctx *parser.PruleContext) []EnrichmentTag {
	var tags = make([]EnrichmentTag, 0)
	ictx := ctx.Tags(0)
	if ictx != nil {
		return append(tags, listener.extractTags(ictx))
	}
	return tags
}

func (listener *sfplListener) getPrefilter(ctx *parser.PruleContext) []string {
	var pfs = make([]string, 0)
	ictx := ctx.Prefilter(0)
	if ictx != nil {
		return append(pfs, listener.extractList(ictx.GetText())...)
	}
	return pfs
}

func (listener *sfplListener) getPriority(ctx *parser.PruleContext) Priority {
	ictx := ctx.Severity(0)
	if ictx != nil {
		p := ictx.GetText()
		switch strings.ToLower(p) {
		case Low.String():
			return Low
		case Medium.String():
			return Medium
		case High.String():
			return High
		case FPriorityDebug:
			return Low
		case FPriorityInfo:
			return Low
		case FPriorityInformational:
			return Low
		case FPriorityNotice:
			return Low
		case FPriorityWarning:
			return Medium
		case FPriorityError:
			return High
		case FPriorityCritical:
			return High
		case FPriorityEmergency:
			return High
		default:
			logger.Warn.Printf("Unrecognized priority value %s. Deferring to %s\n", p, Low.String())
			break
		}
	}
	return Low
}

func (listener *sfplListener) getActions(ctx *parser.PruleContext) []Action {
	var actions []Action
	if ctx.OUTPUT(0) != nil {
		actions = append(actions, Alert)
	} else if ctx.ACTION(0) != nil {
		astr := ctx.Text(2).GetText()
		l := listener.extractList(astr)
		for _, v := range l {
			switch strings.ToLower(v) {
			case Alert.String():
				actions = append(actions, Alert)
			case Tag.String():
				actions = append(actions, Tag)
			case Hash.String():
				actions = append(actions, Hash)
			default:
				logger.Warn.Println("Unrecognized action value ", v)
				break
			}
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
	if ctx != nil {
		return listener.extractList(ctx.GetText())
	}
	return []string{}
}

func (listener *sfplListener) extractTags(ctx parser.ITagsContext) []string {
	if ctx != nil {
		return listener.extractList(ctx.GetText())
	}
	return []string{}
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
		} else if opCtx.ENDSWITH() != nil {
			return EndsWith(lop, rop)
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
		} else if opCtx.LE() != nil {
			return Le(lop, rop)
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
