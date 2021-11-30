//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
// Andreas Schade <san@zurich.ibm.com>
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

// Package engine implements a rules engine for telemetry records.
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

// Regular expression for parsing lists.
var itemsre = regexp.MustCompile(`(^\[)(.*)(\]$?)`)

// PolicyInterpreter defines a rules engine for SysFlow data streams.
type PolicyInterpreter struct {
	*parser.BaseSfplListener

	// Parsed rule and filter object maps.
	rules []Rule
	filters []Filter

	// Accessory parsing maps.
	lists map[string][]string
	macroCtxs map[string]parser.IExpressionContext

	// Action Handler
	ah *ActionHandler
}

// NewPolicyInterpreter constructs a new interpreter instance.
func NewPolicyInterpreter(conf Config) *PolicyInterpreter {
	pi := new(PolicyInterpreter)
	pi.rules = make([]Rule, 0)
	pi.filters = make([]Filter, 0)
	pi.lists = make(map[string][]string)
	pi.macroCtxs = make(map[string]parser.IExpressionContext)

	pi.ah = NewActionHandler(conf)

	return pi
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
	antlr.ParseTreeWalkerDefault.Walk(pi, p.Defs())
	p.GetInputStream().Seek(0)

	// Parse the policy
	antlr.ParseTreeWalkerDefault.Walk(pi, p.Policy())

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
		return errors.New("errors found during compilation of policies. check logs for detail")
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
func (pi *PolicyInterpreter) ProcessAsync(mode Mode, r *Record, out func(r *Record), slots <-chan bool) {
	// Release the execution slot when the record has been processed
	defer func() {
		<-slots
	}()

	// Drop record if amy drop rule applied.
	if pi.EvalFilters(r)  {
		return
	}

	// Enrich mode is non-blocking: Push record even if no rule matches 
	match := (mode == EnrichMode)

	for _, rule := range pi.rules {
		if rule.Enabled && rule.isApplicable(r) && rule.condition.Eval(r) {
			r.Ctx.SetAlert(mode == AlertMode)
			r.Ctx.AddRule(rule)
			pi.ah.HandleActions(rule, r)
			match = true
		}
	}

	// Push record if a rule matched (or if we are in enrich mode)
	if match {
		out(r)
	}
}

// Process executes all compiled policies against record r.
func (pi *PolicyInterpreter) Process(mode Mode, r *Record) *Record {
	// Drop record if amy drop rule applied.
	if pi.EvalFilters(r)  {
		return nil
	}

	// Enrich mode is non-blocking: Push record even if no rule matches 
	match := (mode == EnrichMode)

	for _, rule := range pi.rules {
		if rule.Enabled && rule.isApplicable(r) && rule.condition.Eval(r) {
			r.Ctx.SetAlert(mode == AlertMode)
			r.Ctx.AddRule(rule)
			pi.ah.HandleActions(rule, r)
			match = true
		}
	}

	// Push record if a rule matched (or if we are in enrich mode)
	if match {
		return r
	}
	return nil
}

// EvalFilters executes compiled policy filters against record r.
func (pi *PolicyInterpreter) EvalFilters(r *Record) bool {
	for _, f := range pi.filters {
		if f.Enabled && f.condition.Eval(r) {
			return true
		}
	}
	return false
}

// ExitList is called when production list is exited.
func (pi *PolicyInterpreter) ExitPlist(ctx *parser.PlistContext) {
	logger.Trace.Println("Parsing list ", ctx.GetText())
	pi.lists[ctx.ID().GetText()] = pi.extractListFromItems(ctx.Items())
}

// ExitMacro is called when production macro is exited.
func (pi *PolicyInterpreter) ExitPmacro(ctx *parser.PmacroContext) {
	logger.Trace.Println("Parsing macro ", ctx.GetText())
	pi.macroCtxs[ctx.ID().GetText()] = ctx.Expression()
}

// ExitFilter is called when production filter is exited.
func (pi *PolicyInterpreter) ExitPfilter(ctx *parser.PfilterContext) {
	logger.Trace.Println("Parsing filter ", ctx.GetText())
	f := Filter{
		Name:      ctx.ID().GetText(),
		condition: pi.visitExpression(ctx.Expression()),
		Enabled:   ctx.ENABLED() == nil || pi.getEnabledFlag(ctx.Enabled()),
	}
	pi.filters = append(pi.filters, f)
}

// ExitFilter is called when production filter is exited.
func (pi *PolicyInterpreter) ExitPrule(ctx *parser.PruleContext) {
	logger.Trace.Println("Parsing rule ", ctx.GetText())
	r := Rule{
		Name:      pi.getOffChannelText(ctx.Text(0)),
		Desc:      pi.getOffChannelText(ctx.Text(1)),
		condition: pi.visitExpression(ctx.Expression()),
		Actions:   pi.getActions(ctx),
		Tags:      pi.getTags(ctx),
		Priority:  pi.getPriority(ctx),
		Prefilter: pi.getPrefilter(ctx),
		Enabled:   ctx.ENABLED(0) == nil || pi.getEnabledFlag(ctx.Enabled(0)),
	}
	pi.rules = append(pi.rules, r)
}

func (pi *PolicyInterpreter) getEnabledFlag(ctx parser.IEnabledContext) bool {
	flag := trimBoundingQuotes(ctx.GetText())
	if b, err := strconv.ParseBool(flag); err == nil {
		return b
	}
	logger.Warn.Println("Unrecognized enabled flag: ", flag)
	return true
}

func (pi *PolicyInterpreter) getOffChannelText(ctx parser.ITextContext) string {
	a := ctx.GetStart().GetStart()
	b := ctx.GetStop().GetStop()
	interval := antlr.Interval{Start: a, Stop: b}
	return ctx.GetStart().GetInputStream().GetTextFromInterval(&interval)
}

func (pi *PolicyInterpreter) getTags(ctx *parser.PruleContext) []EnrichmentTag {
	var tags = make([]EnrichmentTag, 0)
	ictx := ctx.Tags(0)
	if ictx != nil {
		return append(tags, pi.extractTags(ictx))
	}
	return tags
}

func (pi *PolicyInterpreter) getPrefilter(ctx *parser.PruleContext) []string {
	var pfs = make([]string, 0)
	ictx := ctx.Prefilter(0)
	if ictx != nil {
		return append(pfs, pi.extractList(ictx.GetText())...)
	}
	return pfs
}

func (pi *PolicyInterpreter) getPriority(ctx *parser.PruleContext) Priority {
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
		}
	}
	return Low
}

func (pi *PolicyInterpreter) getActions(ctx *parser.PruleContext) []string {
	var actions []string
//	if ctx.OUTPUT(0) != nil {
//		actions = append(actions, Alert)
//	} else if ctx.ACTION(0) != nil {
//		astr := ctx.Text(2).GetText()
//		l := pi.extractList(astr)
//		for _, v := range l {
//			switch strings.ToLower(v) {
//			case Alert.String():
//				actions = append(actions, Alert)
//			case Tag.String():
//				actions = append(actions, Tag)
//			case Hash.String():
//				actions = append(actions, Hash)
//			default:
//				logger.Warn.Println("Unrecognized action value ", v)
//			}
//		}
//	}
	if ctx.ACTION(0) != nil || ctx.OUTPUT(0) != nil {
              astr := ctx.Text(2).GetText()
              l := pi.extractList(astr)
              for _, v := range l {
                      actions = append(actions, strings.ToLower(v))
              }
	}
	return actions
}

func (pi *PolicyInterpreter) extractList(str string) []string {
	return strings.Split(itemsre.ReplaceAllString(str, "$2"), LISTSEP)
}

func (pi *PolicyInterpreter) extractListFromItems(ctx parser.IItemsContext) []string {
	if ctx != nil {
		return pi.extractList(ctx.GetText())
	}
	return []string{}
}

func (pi *PolicyInterpreter) extractTags(ctx parser.ITagsContext) []string {
	if ctx != nil {
		return pi.extractList(ctx.GetText())
	}
	return []string{}
}

func (pi *PolicyInterpreter) extractListFromAtoms(ctxs []parser.IAtomContext) []string {
	s := []string{}
	for _, v := range ctxs {
		s = append(s, pi.reduceList(v.GetText())...)
	}
	return s
}

func (pi *PolicyInterpreter) reduceList(sl string) []string {
	s := []string{}
	if l, ok := pi.lists[sl]; ok {
		for _, v := range l {
			s = append(s, pi.reduceList(v)...)
		}
	} else {
		s = append(s, trimBoundingQuotes(sl))
	}
	return s
}

func (pi *PolicyInterpreter) visitExpression(ctx parser.IExpressionContext) Criterion {
	orCtx := ctx.GetChild(0).(parser.IOr_expressionContext)
	orPreds := make([]Criterion, 0)
	for _, andCtx := range orCtx.GetChildren() {
		if andCtx.GetChildCount() > 0 {
			andPreds := make([]Criterion, 0)
			for _, termCtx := range andCtx.GetChildren() {
				t, isTermCtx := termCtx.(parser.ITermContext)
				if isTermCtx {
					c := pi.visitTerm(t)
					andPreds = append(andPreds, c)
				}
			}
			orPreds = append(orPreds, All(andPreds))
		}
	}
	return Any(orPreds)
}

func (pi *PolicyInterpreter) visitTerm(ctx parser.ITermContext) Criterion {
	termCtx := ctx.(*parser.TermContext)
	if termCtx.Variable() != nil {
		if m, ok := pi.macroCtxs[termCtx.GetText()]; ok {
			return pi.visitExpression(m)
		}
		logger.Error.Println("Unrecognized reference ", termCtx.GetText())
	} else if termCtx.NOT() != nil {
		return pi.visitTerm(termCtx.GetChild(1).(parser.ITermContext)).Not()
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
		return pi.visitExpression(termCtx.Expression())
	} else if termCtx.IN() != nil {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		rop := termCtx.AllAtom()[1:]
		return In(lop, pi.extractListFromAtoms(rop))
	} else if termCtx.PMATCH() != nil {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		rop := termCtx.AllAtom()[1:]
		return PMatch(lop, pi.extractListFromAtoms(rop))
	} else {
		logger.Warn.Println("Unrecognized term ", termCtx.GetText())
	}
	return False
}
