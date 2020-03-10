// Code generated from Sfpl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Sfpl
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 170,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 3, 2,
	6, 2, 37, 10, 2, 13, 2, 14, 2, 38, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	7, 8, 95, 10, 8, 12, 8, 14, 8, 98, 11, 8, 3, 9, 3, 9, 3, 9, 7, 9, 103,
	10, 9, 12, 9, 14, 9, 106, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	123, 10, 10, 3, 10, 3, 10, 3, 10, 5, 10, 128, 10, 10, 7, 10, 130, 10, 10,
	12, 10, 14, 10, 133, 11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	5, 10, 141, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 147, 10, 11, 12,
	11, 14, 11, 150, 11, 11, 5, 11, 152, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 6, 14, 162, 10, 14, 13, 14, 14, 14, 163, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 16, 2, 2, 17, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 2, 6, 3, 2, 11, 12, 4, 2, 24, 24, 28, 28, 5, 2,
	18, 18, 20, 20, 38, 41, 4, 2, 18, 23, 25, 27, 2, 171, 2, 36, 3, 2, 2, 2,
	4, 42, 3, 2, 2, 2, 6, 62, 3, 2, 2, 2, 8, 73, 3, 2, 2, 2, 10, 81, 3, 2,
	2, 2, 12, 89, 3, 2, 2, 2, 14, 91, 3, 2, 2, 2, 16, 99, 3, 2, 2, 2, 18, 140,
	3, 2, 2, 2, 20, 142, 3, 2, 2, 2, 22, 155, 3, 2, 2, 2, 24, 157, 3, 2, 2,
	2, 26, 161, 3, 2, 2, 2, 28, 165, 3, 2, 2, 2, 30, 167, 3, 2, 2, 2, 32, 37,
	5, 4, 3, 2, 33, 37, 5, 6, 4, 2, 34, 37, 5, 8, 5, 2, 35, 37, 5, 10, 6, 2,
	36, 32, 3, 2, 2, 2, 36, 33, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 35, 3,
	2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39,
	40, 3, 2, 2, 2, 40, 41, 7, 2, 2, 3, 41, 3, 3, 2, 2, 2, 42, 43, 7, 35, 2,
	2, 43, 44, 7, 3, 2, 2, 44, 45, 7, 36, 2, 2, 45, 46, 5, 26, 14, 2, 46, 47,
	7, 10, 2, 2, 47, 48, 7, 36, 2, 2, 48, 49, 5, 26, 14, 2, 49, 50, 7, 9, 2,
	2, 50, 51, 7, 36, 2, 2, 51, 52, 5, 12, 7, 2, 52, 53, 9, 2, 2, 2, 53, 54,
	7, 36, 2, 2, 54, 55, 5, 26, 14, 2, 55, 56, 7, 13, 2, 2, 56, 57, 7, 36,
	2, 2, 57, 58, 7, 37, 2, 2, 58, 59, 7, 14, 2, 2, 59, 60, 7, 36, 2, 2, 60,
	61, 5, 20, 11, 2, 61, 5, 3, 2, 2, 2, 62, 63, 7, 35, 2, 2, 63, 64, 7, 4,
	2, 2, 64, 65, 7, 36, 2, 2, 65, 66, 5, 26, 14, 2, 66, 67, 7, 10, 2, 2, 67,
	68, 7, 36, 2, 2, 68, 69, 5, 26, 14, 2, 69, 70, 7, 9, 2, 2, 70, 71, 7, 36,
	2, 2, 71, 72, 5, 12, 7, 2, 72, 7, 3, 2, 2, 2, 73, 74, 7, 35, 2, 2, 74,
	75, 7, 5, 2, 2, 75, 76, 7, 36, 2, 2, 76, 77, 7, 38, 2, 2, 77, 78, 7, 9,
	2, 2, 78, 79, 7, 36, 2, 2, 79, 80, 5, 12, 7, 2, 80, 9, 3, 2, 2, 2, 81,
	82, 7, 35, 2, 2, 82, 83, 7, 6, 2, 2, 83, 84, 7, 36, 2, 2, 84, 85, 7, 38,
	2, 2, 85, 86, 7, 8, 2, 2, 86, 87, 7, 36, 2, 2, 87, 88, 5, 20, 11, 2, 88,
	11, 3, 2, 2, 2, 89, 90, 5, 14, 8, 2, 90, 13, 3, 2, 2, 2, 91, 96, 5, 16,
	9, 2, 92, 93, 7, 16, 2, 2, 93, 95, 5, 16, 9, 2, 94, 92, 3, 2, 2, 2, 95,
	98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 15, 3, 2, 2,
	2, 98, 96, 3, 2, 2, 2, 99, 104, 5, 18, 10, 2, 100, 101, 7, 15, 2, 2, 101,
	103, 5, 18, 10, 2, 102, 100, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102,
	3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 17, 3, 2, 2, 2, 106, 104, 3, 2,
	2, 2, 107, 141, 5, 22, 12, 2, 108, 109, 7, 17, 2, 2, 109, 141, 5, 18, 10,
	2, 110, 111, 5, 24, 13, 2, 111, 112, 5, 30, 16, 2, 112, 141, 3, 2, 2, 2,
	113, 114, 5, 24, 13, 2, 114, 115, 5, 28, 15, 2, 115, 116, 5, 24, 13, 2,
	116, 141, 3, 2, 2, 2, 117, 118, 5, 24, 13, 2, 118, 119, 9, 3, 2, 2, 119,
	122, 7, 32, 2, 2, 120, 123, 5, 24, 13, 2, 121, 123, 5, 20, 11, 2, 122,
	120, 3, 2, 2, 2, 122, 121, 3, 2, 2, 2, 123, 131, 3, 2, 2, 2, 124, 127,
	7, 34, 2, 2, 125, 128, 5, 24, 13, 2, 126, 128, 5, 20, 11, 2, 127, 125,
	3, 2, 2, 2, 127, 126, 3, 2, 2, 2, 128, 130, 3, 2, 2, 2, 129, 124, 3, 2,
	2, 2, 130, 133, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2,
	132, 134, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 134, 135, 7, 33, 2, 2, 135,
	141, 3, 2, 2, 2, 136, 137, 7, 32, 2, 2, 137, 138, 5, 12, 7, 2, 138, 139,
	7, 33, 2, 2, 139, 141, 3, 2, 2, 2, 140, 107, 3, 2, 2, 2, 140, 108, 3, 2,
	2, 2, 140, 110, 3, 2, 2, 2, 140, 113, 3, 2, 2, 2, 140, 117, 3, 2, 2, 2,
	140, 136, 3, 2, 2, 2, 141, 19, 3, 2, 2, 2, 142, 151, 7, 30, 2, 2, 143,
	148, 5, 24, 13, 2, 144, 145, 7, 34, 2, 2, 145, 147, 5, 24, 13, 2, 146,
	144, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149,
	3, 2, 2, 2, 149, 152, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151, 143, 3, 2,
	2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154, 7, 31, 2, 2,
	154, 21, 3, 2, 2, 2, 155, 156, 7, 38, 2, 2, 156, 23, 3, 2, 2, 2, 157, 158,
	9, 4, 2, 2, 158, 25, 3, 2, 2, 2, 159, 160, 6, 14, 2, 2, 160, 162, 11, 2,
	2, 2, 161, 159, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2,
	163, 164, 3, 2, 2, 2, 164, 27, 3, 2, 2, 2, 165, 166, 9, 5, 2, 2, 166, 29,
	3, 2, 2, 2, 167, 168, 7, 29, 2, 2, 168, 31, 3, 2, 2, 2, 13, 36, 38, 96,
	104, 122, 127, 131, 140, 148, 151, 163,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'rule'", "'filter'", "'macro'", "'list'", "'name'", "'items'", "'condition'",
	"'desc'", "'action'", "'output'", "'priority'", "'tags'", "'and'", "'or'",
	"'not'", "'<'", "'<='", "'>'", "'>='", "'='", "'!='", "'in'", "'contains'",
	"'icontains'", "'startswith'", "'pmatch'", "'exists'", "'['", "']'", "'('",
	"')'", "','", "'-'",
}
var symbolicNames = []string{
	"", "RULE", "FILTER", "MACRO", "LIST", "NAME", "ITEMS", "COND", "DESC",
	"ACTION", "OUTPUT", "PRIORITY", "TAGS", "AND", "OR", "NOT", "LT", "LE",
	"GT", "GE", "EQ", "NEQ", "IN", "CONTAINS", "ICONTAINS", "STARTSWITH", "PMATCH",
	"EXISTS", "LBRACK", "RBRACK", "LPAREN", "RPAREN", "LISTSEP", "DECL", "DEF",
	"SEVERITY", "ID", "NUMBER", "PATH", "STRING", "WS", "NL", "COMMENT", "ANY",
}

var ruleNames = []string{
	"policy", "f_rule", "f_filter", "f_macro", "f_list", "expression", "or_expression",
	"and_expression", "term", "items", "variable", "atom", "text", "binary_operator",
	"unary_operator",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SfplParser struct {
	*antlr.BaseParser
}

func NewSfplParser(input antlr.TokenStream) *SfplParser {
	this := new(SfplParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sfpl.g4"

	return this
}

// SfplParser tokens.
const (
	SfplParserEOF        = antlr.TokenEOF
	SfplParserRULE       = 1
	SfplParserFILTER     = 2
	SfplParserMACRO      = 3
	SfplParserLIST       = 4
	SfplParserNAME       = 5
	SfplParserITEMS      = 6
	SfplParserCOND       = 7
	SfplParserDESC       = 8
	SfplParserACTION     = 9
	SfplParserOUTPUT     = 10
	SfplParserPRIORITY   = 11
	SfplParserTAGS       = 12
	SfplParserAND        = 13
	SfplParserOR         = 14
	SfplParserNOT        = 15
	SfplParserLT         = 16
	SfplParserLE         = 17
	SfplParserGT         = 18
	SfplParserGE         = 19
	SfplParserEQ         = 20
	SfplParserNEQ        = 21
	SfplParserIN         = 22
	SfplParserCONTAINS   = 23
	SfplParserICONTAINS  = 24
	SfplParserSTARTSWITH = 25
	SfplParserPMATCH     = 26
	SfplParserEXISTS     = 27
	SfplParserLBRACK     = 28
	SfplParserRBRACK     = 29
	SfplParserLPAREN     = 30
	SfplParserRPAREN     = 31
	SfplParserLISTSEP    = 32
	SfplParserDECL       = 33
	SfplParserDEF        = 34
	SfplParserSEVERITY   = 35
	SfplParserID         = 36
	SfplParserNUMBER     = 37
	SfplParserPATH       = 38
	SfplParserSTRING     = 39
	SfplParserWS         = 40
	SfplParserNL         = 41
	SfplParserCOMMENT    = 42
	SfplParserANY        = 43
)

// SfplParser rules.
const (
	SfplParserRULE_policy          = 0
	SfplParserRULE_f_rule          = 1
	SfplParserRULE_f_filter        = 2
	SfplParserRULE_f_macro         = 3
	SfplParserRULE_f_list          = 4
	SfplParserRULE_expression      = 5
	SfplParserRULE_or_expression   = 6
	SfplParserRULE_and_expression  = 7
	SfplParserRULE_term            = 8
	SfplParserRULE_items           = 9
	SfplParserRULE_variable        = 10
	SfplParserRULE_atom            = 11
	SfplParserRULE_text            = 12
	SfplParserRULE_binary_operator = 13
	SfplParserRULE_unary_operator  = 14
)

// IPolicyContext is an interface to support dynamic dispatch.
type IPolicyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolicyContext differentiates from other interfaces.
	IsPolicyContext()
}

type PolicyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPolicyContext() *PolicyContext {
	var p = new(PolicyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_policy
	return p
}

func (*PolicyContext) IsPolicyContext() {}

func NewPolicyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolicyContext {
	var p = new(PolicyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_policy

	return p
}

func (s *PolicyContext) GetParser() antlr.Parser { return s.parser }

func (s *PolicyContext) EOF() antlr.TerminalNode {
	return s.GetToken(SfplParserEOF, 0)
}

func (s *PolicyContext) AllF_rule() []IF_ruleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IF_ruleContext)(nil)).Elem())
	var tst = make([]IF_ruleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IF_ruleContext)
		}
	}

	return tst
}

func (s *PolicyContext) F_rule(i int) IF_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IF_ruleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IF_ruleContext)
}

func (s *PolicyContext) AllF_filter() []IF_filterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IF_filterContext)(nil)).Elem())
	var tst = make([]IF_filterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IF_filterContext)
		}
	}

	return tst
}

func (s *PolicyContext) F_filter(i int) IF_filterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IF_filterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IF_filterContext)
}

func (s *PolicyContext) AllF_macro() []IF_macroContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IF_macroContext)(nil)).Elem())
	var tst = make([]IF_macroContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IF_macroContext)
		}
	}

	return tst
}

func (s *PolicyContext) F_macro(i int) IF_macroContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IF_macroContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IF_macroContext)
}

func (s *PolicyContext) AllF_list() []IF_listContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IF_listContext)(nil)).Elem())
	var tst = make([]IF_listContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IF_listContext)
		}
	}

	return tst
}

func (s *PolicyContext) F_list(i int) IF_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IF_listContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IF_listContext)
}

func (s *PolicyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolicyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolicyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPolicy(s)
	}
}

func (s *PolicyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPolicy(s)
	}
}

func (s *PolicyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPolicy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Policy() (localctx IPolicyContext) {
	localctx = NewPolicyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SfplParserRULE_policy)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SfplParserDECL {
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(30)
				p.F_rule()
			}

		case 2:
			{
				p.SetState(31)
				p.F_filter()
			}

		case 3:
			{
				p.SetState(32)
				p.F_macro()
			}

		case 4:
			{
				p.SetState(33)
				p.F_list()
			}

		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(SfplParserEOF)
	}

	return localctx
}

// IF_ruleContext is an interface to support dynamic dispatch.
type IF_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsF_ruleContext differentiates from other interfaces.
	IsF_ruleContext()
}

type F_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyF_ruleContext() *F_ruleContext {
	var p = new(F_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_f_rule
	return p
}

func (*F_ruleContext) IsF_ruleContext() {}

func NewF_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *F_ruleContext {
	var p = new(F_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_f_rule

	return p
}

func (s *F_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *F_ruleContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *F_ruleContext) RULE() antlr.TerminalNode {
	return s.GetToken(SfplParserRULE, 0)
}

func (s *F_ruleContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *F_ruleContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *F_ruleContext) AllText() []ITextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITextContext)(nil)).Elem())
	var tst = make([]ITextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITextContext)
		}
	}

	return tst
}

func (s *F_ruleContext) Text(i int) ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *F_ruleContext) DESC() antlr.TerminalNode {
	return s.GetToken(SfplParserDESC, 0)
}

func (s *F_ruleContext) COND() antlr.TerminalNode {
	return s.GetToken(SfplParserCOND, 0)
}

func (s *F_ruleContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *F_ruleContext) PRIORITY() antlr.TerminalNode {
	return s.GetToken(SfplParserPRIORITY, 0)
}

func (s *F_ruleContext) SEVERITY() antlr.TerminalNode {
	return s.GetToken(SfplParserSEVERITY, 0)
}

func (s *F_ruleContext) TAGS() antlr.TerminalNode {
	return s.GetToken(SfplParserTAGS, 0)
}

func (s *F_ruleContext) Items() IItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemsContext)
}

func (s *F_ruleContext) ACTION() antlr.TerminalNode {
	return s.GetToken(SfplParserACTION, 0)
}

func (s *F_ruleContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(SfplParserOUTPUT, 0)
}

func (s *F_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *F_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterF_rule(s)
	}
}

func (s *F_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitF_rule(s)
	}
}

func (s *F_ruleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitF_rule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) F_rule() (localctx IF_ruleContext) {
	localctx = NewF_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SfplParserRULE_f_rule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(41)
		p.Match(SfplParserRULE)
	}
	{
		p.SetState(42)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(43)
		p.Text()
	}
	{
		p.SetState(44)
		p.Match(SfplParserDESC)
	}
	{
		p.SetState(45)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(46)
		p.Text()
	}
	{
		p.SetState(47)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(48)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(49)
		p.Expression()
	}
	{
		p.SetState(50)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SfplParserACTION || _la == SfplParserOUTPUT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(51)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(52)
		p.Text()
	}
	{
		p.SetState(53)
		p.Match(SfplParserPRIORITY)
	}
	{
		p.SetState(54)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(55)
		p.Match(SfplParserSEVERITY)
	}
	{
		p.SetState(56)
		p.Match(SfplParserTAGS)
	}
	{
		p.SetState(57)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(58)
		p.Items()
	}

	return localctx
}

// IF_filterContext is an interface to support dynamic dispatch.
type IF_filterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsF_filterContext differentiates from other interfaces.
	IsF_filterContext()
}

type F_filterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyF_filterContext() *F_filterContext {
	var p = new(F_filterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_f_filter
	return p
}

func (*F_filterContext) IsF_filterContext() {}

func NewF_filterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *F_filterContext {
	var p = new(F_filterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_f_filter

	return p
}

func (s *F_filterContext) GetParser() antlr.Parser { return s.parser }

func (s *F_filterContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *F_filterContext) FILTER() antlr.TerminalNode {
	return s.GetToken(SfplParserFILTER, 0)
}

func (s *F_filterContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *F_filterContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *F_filterContext) AllText() []ITextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITextContext)(nil)).Elem())
	var tst = make([]ITextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITextContext)
		}
	}

	return tst
}

func (s *F_filterContext) Text(i int) ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *F_filterContext) DESC() antlr.TerminalNode {
	return s.GetToken(SfplParserDESC, 0)
}

func (s *F_filterContext) COND() antlr.TerminalNode {
	return s.GetToken(SfplParserCOND, 0)
}

func (s *F_filterContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *F_filterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F_filterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *F_filterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterF_filter(s)
	}
}

func (s *F_filterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitF_filter(s)
	}
}

func (s *F_filterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitF_filter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) F_filter() (localctx IF_filterContext) {
	localctx = NewF_filterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SfplParserRULE_f_filter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(61)
		p.Match(SfplParserFILTER)
	}
	{
		p.SetState(62)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(63)
		p.Text()
	}
	{
		p.SetState(64)
		p.Match(SfplParserDESC)
	}
	{
		p.SetState(65)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(66)
		p.Text()
	}
	{
		p.SetState(67)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(68)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(69)
		p.Expression()
	}

	return localctx
}

// IF_macroContext is an interface to support dynamic dispatch.
type IF_macroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsF_macroContext differentiates from other interfaces.
	IsF_macroContext()
}

type F_macroContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyF_macroContext() *F_macroContext {
	var p = new(F_macroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_f_macro
	return p
}

func (*F_macroContext) IsF_macroContext() {}

func NewF_macroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *F_macroContext {
	var p = new(F_macroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_f_macro

	return p
}

func (s *F_macroContext) GetParser() antlr.Parser { return s.parser }

func (s *F_macroContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *F_macroContext) MACRO() antlr.TerminalNode {
	return s.GetToken(SfplParserMACRO, 0)
}

func (s *F_macroContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *F_macroContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *F_macroContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *F_macroContext) COND() antlr.TerminalNode {
	return s.GetToken(SfplParserCOND, 0)
}

func (s *F_macroContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *F_macroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F_macroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *F_macroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterF_macro(s)
	}
}

func (s *F_macroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitF_macro(s)
	}
}

func (s *F_macroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitF_macro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) F_macro() (localctx IF_macroContext) {
	localctx = NewF_macroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SfplParserRULE_f_macro)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(72)
		p.Match(SfplParserMACRO)
	}
	{
		p.SetState(73)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(74)
		p.Match(SfplParserID)
	}
	{
		p.SetState(75)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(76)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(77)
		p.Expression()
	}

	return localctx
}

// IF_listContext is an interface to support dynamic dispatch.
type IF_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsF_listContext differentiates from other interfaces.
	IsF_listContext()
}

type F_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyF_listContext() *F_listContext {
	var p = new(F_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_f_list
	return p
}

func (*F_listContext) IsF_listContext() {}

func NewF_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *F_listContext {
	var p = new(F_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_f_list

	return p
}

func (s *F_listContext) GetParser() antlr.Parser { return s.parser }

func (s *F_listContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *F_listContext) LIST() antlr.TerminalNode {
	return s.GetToken(SfplParserLIST, 0)
}

func (s *F_listContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *F_listContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *F_listContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *F_listContext) ITEMS() antlr.TerminalNode {
	return s.GetToken(SfplParserITEMS, 0)
}

func (s *F_listContext) Items() IItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemsContext)
}

func (s *F_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *F_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterF_list(s)
	}
}

func (s *F_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitF_list(s)
	}
}

func (s *F_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitF_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) F_list() (localctx IF_listContext) {
	localctx = NewF_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SfplParserRULE_f_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(80)
		p.Match(SfplParserLIST)
	}
	{
		p.SetState(81)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(82)
		p.Match(SfplParserID)
	}
	{
		p.SetState(83)
		p.Match(SfplParserITEMS)
	}
	{
		p.SetState(84)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(85)
		p.Items()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Or_expression() IOr_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_expressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SfplParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Or_expression()
	}

	return localctx
}

// IOr_expressionContext is an interface to support dynamic dispatch.
type IOr_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOr_expressionContext differentiates from other interfaces.
	IsOr_expressionContext()
}

type Or_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOr_expressionContext() *Or_expressionContext {
	var p = new(Or_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_or_expression
	return p
}

func (*Or_expressionContext) IsOr_expressionContext() {}

func NewOr_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_expressionContext {
	var p = new(Or_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_or_expression

	return p
}

func (s *Or_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_expressionContext) AllAnd_expression() []IAnd_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnd_expressionContext)(nil)).Elem())
	var tst = make([]IAnd_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnd_expressionContext)
		}
	}

	return tst
}

func (s *Or_expressionContext) And_expression(i int) IAnd_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnd_expressionContext)
}

func (s *Or_expressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(SfplParserOR)
}

func (s *Or_expressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserOR, i)
}

func (s *Or_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterOr_expression(s)
	}
}

func (s *Or_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitOr_expression(s)
	}
}

func (s *Or_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitOr_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Or_expression() (localctx IOr_expressionContext) {
	localctx = NewOr_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SfplParserRULE_or_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.And_expression()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SfplParserOR {
		{
			p.SetState(90)
			p.Match(SfplParserOR)
		}
		{
			p.SetState(91)
			p.And_expression()
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAnd_expressionContext is an interface to support dynamic dispatch.
type IAnd_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnd_expressionContext differentiates from other interfaces.
	IsAnd_expressionContext()
}

type And_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_expressionContext() *And_expressionContext {
	var p = new(And_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_and_expression
	return p
}

func (*And_expressionContext) IsAnd_expressionContext() {}

func NewAnd_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_expressionContext {
	var p = new(And_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_and_expression

	return p
}

func (s *And_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *And_expressionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *And_expressionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *And_expressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(SfplParserAND)
}

func (s *And_expressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserAND, i)
}

func (s *And_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *And_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterAnd_expression(s)
	}
}

func (s *And_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitAnd_expression(s)
	}
}

func (s *And_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitAnd_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) And_expression() (localctx IAnd_expressionContext) {
	localctx = NewAnd_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SfplParserRULE_and_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Term()
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SfplParserAND {
		{
			p.SetState(98)
			p.Match(SfplParserAND)
		}
		{
			p.SetState(99)
			p.Term()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *TermContext) NOT() antlr.TerminalNode {
	return s.GetToken(SfplParserNOT, 0)
}

func (s *TermContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *TermContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *TermContext) Unary_operator() IUnary_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnary_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnary_operatorContext)
}

func (s *TermContext) Binary_operator() IBinary_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_operatorContext)
}

func (s *TermContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SfplParserLPAREN, 0)
}

func (s *TermContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SfplParserRPAREN, 0)
}

func (s *TermContext) IN() antlr.TerminalNode {
	return s.GetToken(SfplParserIN, 0)
}

func (s *TermContext) PMATCH() antlr.TerminalNode {
	return s.GetToken(SfplParserPMATCH, 0)
}

func (s *TermContext) AllItems() []IItemsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IItemsContext)(nil)).Elem())
	var tst = make([]IItemsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IItemsContext)
		}
	}

	return tst
}

func (s *TermContext) Items(i int) IItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IItemsContext)
}

func (s *TermContext) AllLISTSEP() []antlr.TerminalNode {
	return s.GetTokens(SfplParserLISTSEP)
}

func (s *TermContext) LISTSEP(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserLISTSEP, i)
}

func (s *TermContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SfplParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.Match(SfplParserNOT)
		}
		{
			p.SetState(107)
			p.Term()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.Atom()
		}
		{
			p.SetState(109)
			p.Unary_operator()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.Atom()
		}
		{
			p.SetState(112)
			p.Binary_operator()
		}
		{
			p.SetState(113)
			p.Atom()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(115)
			p.Atom()
		}
		{
			p.SetState(116)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SfplParserIN || _la == SfplParserPMATCH) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(117)
			p.Match(SfplParserLPAREN)
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SfplParserLT, SfplParserGT, SfplParserID, SfplParserNUMBER, SfplParserPATH, SfplParserSTRING:
			{
				p.SetState(118)
				p.Atom()
			}

		case SfplParserLBRACK:
			{
				p.SetState(119)
				p.Items()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SfplParserLISTSEP {
			{
				p.SetState(122)
				p.Match(SfplParserLISTSEP)
			}
			p.SetState(125)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case SfplParserLT, SfplParserGT, SfplParserID, SfplParserNUMBER, SfplParserPATH, SfplParserSTRING:
				{
					p.SetState(123)
					p.Atom()
				}

			case SfplParserLBRACK:
				{
					p.SetState(124)
					p.Items()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(132)
			p.Match(SfplParserRPAREN)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(134)
			p.Match(SfplParserLPAREN)
		}
		{
			p.SetState(135)
			p.Expression()
		}
		{
			p.SetState(136)
			p.Match(SfplParserRPAREN)
		}

	}

	return localctx
}

// IItemsContext is an interface to support dynamic dispatch.
type IItemsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemsContext differentiates from other interfaces.
	IsItemsContext()
}

type ItemsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemsContext() *ItemsContext {
	var p = new(ItemsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_items
	return p
}

func (*ItemsContext) IsItemsContext() {}

func NewItemsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemsContext {
	var p = new(ItemsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_items

	return p
}

func (s *ItemsContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemsContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(SfplParserLBRACK, 0)
}

func (s *ItemsContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(SfplParserRBRACK, 0)
}

func (s *ItemsContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *ItemsContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ItemsContext) AllLISTSEP() []antlr.TerminalNode {
	return s.GetTokens(SfplParserLISTSEP)
}

func (s *ItemsContext) LISTSEP(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserLISTSEP, i)
}

func (s *ItemsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterItems(s)
	}
}

func (s *ItemsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitItems(s)
	}
}

func (s *ItemsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitItems(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Items() (localctx IItemsContext) {
	localctx = NewItemsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SfplParserRULE_items)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(SfplParserLBRACK)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(SfplParserLT-16))|(1<<(SfplParserGT-16))|(1<<(SfplParserID-16))|(1<<(SfplParserNUMBER-16))|(1<<(SfplParserPATH-16))|(1<<(SfplParserSTRING-16)))) != 0 {
		{
			p.SetState(141)
			p.Atom()
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SfplParserLISTSEP {
			{
				p.SetState(142)
				p.Match(SfplParserLISTSEP)
			}
			{
				p.SetState(143)
				p.Atom()
			}

			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(151)
		p.Match(SfplParserRBRACK)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SfplParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(SfplParserID)
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *AtomContext) PATH() antlr.TerminalNode {
	return s.GetToken(SfplParserPATH, 0)
}

func (s *AtomContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SfplParserNUMBER, 0)
}

func (s *AtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(SfplParserSTRING, 0)
}

func (s *AtomContext) LT() antlr.TerminalNode {
	return s.GetToken(SfplParserLT, 0)
}

func (s *AtomContext) GT() antlr.TerminalNode {
	return s.GetToken(SfplParserGT, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SfplParserRULE_atom)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(SfplParserLT-16))|(1<<(SfplParserGT-16))|(1<<(SfplParserID-16))|(1<<(SfplParserNUMBER-16))|(1<<(SfplParserPATH-16))|(1<<(SfplParserSTRING-16)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }
func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitText(s)
	}
}

func (s *TextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SfplParserRULE_text)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(157)

			if !(!(p.GetCurrentToken().GetText() == "desc" ||
				p.GetCurrentToken().GetText() == "condition" ||
				p.GetCurrentToken().GetText() == "action" ||
				p.GetCurrentToken().GetText() == "output" ||
				p.GetCurrentToken().GetText() == "priority" ||
				p.GetCurrentToken().GetText() == "tags")) {
				panic(antlr.NewFailedPredicateException(p, "!(p.GetCurrentToken().GetText() == \"desc\" ||\n\t      p.GetCurrentToken().GetText() == \"condition\" ||\n\t      p.GetCurrentToken().GetText() == \"action\" ||\n\t      p.GetCurrentToken().GetText() == \"output\" ||\n\t      p.GetCurrentToken().GetText() == \"priority\" ||\n\t      p.GetCurrentToken().GetText() == \"tags\")", ""))
			}
			p.SetState(158)
			p.MatchWildcard()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IBinary_operatorContext is an interface to support dynamic dispatch.
type IBinary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_operatorContext differentiates from other interfaces.
	IsBinary_operatorContext()
}

type Binary_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_operatorContext() *Binary_operatorContext {
	var p = new(Binary_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_binary_operator
	return p
}

func (*Binary_operatorContext) IsBinary_operatorContext() {}

func NewBinary_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_operatorContext {
	var p = new(Binary_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_binary_operator

	return p
}

func (s *Binary_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Binary_operatorContext) LT() antlr.TerminalNode {
	return s.GetToken(SfplParserLT, 0)
}

func (s *Binary_operatorContext) LE() antlr.TerminalNode {
	return s.GetToken(SfplParserLE, 0)
}

func (s *Binary_operatorContext) GT() antlr.TerminalNode {
	return s.GetToken(SfplParserGT, 0)
}

func (s *Binary_operatorContext) GE() antlr.TerminalNode {
	return s.GetToken(SfplParserGE, 0)
}

func (s *Binary_operatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(SfplParserEQ, 0)
}

func (s *Binary_operatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(SfplParserNEQ, 0)
}

func (s *Binary_operatorContext) CONTAINS() antlr.TerminalNode {
	return s.GetToken(SfplParserCONTAINS, 0)
}

func (s *Binary_operatorContext) ICONTAINS() antlr.TerminalNode {
	return s.GetToken(SfplParserICONTAINS, 0)
}

func (s *Binary_operatorContext) STARTSWITH() antlr.TerminalNode {
	return s.GetToken(SfplParserSTARTSWITH, 0)
}

func (s *Binary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Binary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterBinary_operator(s)
	}
}

func (s *Binary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitBinary_operator(s)
	}
}

func (s *Binary_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitBinary_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Binary_operator() (localctx IBinary_operatorContext) {
	localctx = NewBinary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SfplParserRULE_binary_operator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfplParserLT)|(1<<SfplParserLE)|(1<<SfplParserGT)|(1<<SfplParserGE)|(1<<SfplParserEQ)|(1<<SfplParserNEQ)|(1<<SfplParserCONTAINS)|(1<<SfplParserICONTAINS)|(1<<SfplParserSTARTSWITH))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnary_operatorContext is an interface to support dynamic dispatch.
type IUnary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnary_operatorContext differentiates from other interfaces.
	IsUnary_operatorContext()
}

type Unary_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_operatorContext() *Unary_operatorContext {
	var p = new(Unary_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_unary_operator
	return p
}

func (*Unary_operatorContext) IsUnary_operatorContext() {}

func NewUnary_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_operatorContext {
	var p = new(Unary_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_unary_operator

	return p
}

func (s *Unary_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_operatorContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(SfplParserEXISTS, 0)
}

func (s *Unary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterUnary_operator(s)
	}
}

func (s *Unary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitUnary_operator(s)
	}
}

func (s *Unary_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitUnary_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Unary_operator() (localctx IUnary_operatorContext) {
	localctx = NewUnary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SfplParserRULE_unary_operator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(SfplParserEXISTS)
	}

	return localctx
}

func (p *SfplParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *TextContext = nil
		if localctx != nil {
			t = localctx.(*TextContext)
		}
		return p.Text_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SfplParser) Text_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return !(p.GetCurrentToken().GetText() == "desc" ||
			p.GetCurrentToken().GetText() == "condition" ||
			p.GetCurrentToken().GetText() == "action" ||
			p.GetCurrentToken().GetText() == "output" ||
			p.GetCurrentToken().GetText() == "priority" ||
			p.GetCurrentToken().GetText() == "tags")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
