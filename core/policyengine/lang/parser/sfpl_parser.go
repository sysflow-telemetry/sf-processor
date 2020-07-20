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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 167,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 3, 2,
	6, 2, 37, 10, 2, 13, 2, 14, 2, 38, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 92, 10, 8,
	12, 8, 14, 8, 95, 11, 8, 3, 9, 3, 9, 3, 9, 7, 9, 100, 10, 9, 12, 9, 14,
	9, 103, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 120, 10, 10, 3,
	10, 3, 10, 3, 10, 5, 10, 125, 10, 10, 7, 10, 127, 10, 10, 12, 10, 14, 10,
	130, 11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 138, 10,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 144, 10, 11, 12, 11, 14, 11, 147,
	11, 11, 5, 11, 149, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	14, 3, 14, 6, 14, 159, 10, 14, 13, 14, 14, 14, 160, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 2, 2, 17, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 2, 6, 3, 2, 11, 12, 4, 2, 24, 24, 28, 28, 5, 2, 18, 18, 20, 20,
	38, 41, 4, 2, 18, 23, 25, 27, 2, 168, 2, 36, 3, 2, 2, 2, 4, 42, 3, 2, 2,
	2, 6, 62, 3, 2, 2, 2, 8, 70, 3, 2, 2, 2, 10, 78, 3, 2, 2, 2, 12, 86, 3,
	2, 2, 2, 14, 88, 3, 2, 2, 2, 16, 96, 3, 2, 2, 2, 18, 137, 3, 2, 2, 2, 20,
	139, 3, 2, 2, 2, 22, 152, 3, 2, 2, 2, 24, 154, 3, 2, 2, 2, 26, 158, 3,
	2, 2, 2, 28, 162, 3, 2, 2, 2, 30, 164, 3, 2, 2, 2, 32, 37, 5, 4, 3, 2,
	33, 37, 5, 6, 4, 2, 34, 37, 5, 8, 5, 2, 35, 37, 5, 10, 6, 2, 36, 32, 3,
	2, 2, 2, 36, 33, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 35, 3, 2, 2, 2, 37,
	38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 40, 3, 2, 2,
	2, 40, 41, 7, 2, 2, 3, 41, 3, 3, 2, 2, 2, 42, 43, 7, 35, 2, 2, 43, 44,
	7, 3, 2, 2, 44, 45, 7, 36, 2, 2, 45, 46, 5, 26, 14, 2, 46, 47, 7, 10, 2,
	2, 47, 48, 7, 36, 2, 2, 48, 49, 5, 26, 14, 2, 49, 50, 7, 9, 2, 2, 50, 51,
	7, 36, 2, 2, 51, 52, 5, 12, 7, 2, 52, 53, 9, 2, 2, 2, 53, 54, 7, 36, 2,
	2, 54, 55, 5, 26, 14, 2, 55, 56, 7, 13, 2, 2, 56, 57, 7, 36, 2, 2, 57,
	58, 7, 37, 2, 2, 58, 59, 7, 14, 2, 2, 59, 60, 7, 36, 2, 2, 60, 61, 5, 20,
	11, 2, 61, 5, 3, 2, 2, 2, 62, 63, 7, 35, 2, 2, 63, 64, 7, 4, 2, 2, 64,
	65, 7, 36, 2, 2, 65, 66, 7, 38, 2, 2, 66, 67, 7, 9, 2, 2, 67, 68, 7, 36,
	2, 2, 68, 69, 5, 12, 7, 2, 69, 7, 3, 2, 2, 2, 70, 71, 7, 35, 2, 2, 71,
	72, 7, 5, 2, 2, 72, 73, 7, 36, 2, 2, 73, 74, 7, 38, 2, 2, 74, 75, 7, 9,
	2, 2, 75, 76, 7, 36, 2, 2, 76, 77, 5, 12, 7, 2, 77, 9, 3, 2, 2, 2, 78,
	79, 7, 35, 2, 2, 79, 80, 7, 6, 2, 2, 80, 81, 7, 36, 2, 2, 81, 82, 7, 38,
	2, 2, 82, 83, 7, 8, 2, 2, 83, 84, 7, 36, 2, 2, 84, 85, 5, 20, 11, 2, 85,
	11, 3, 2, 2, 2, 86, 87, 5, 14, 8, 2, 87, 13, 3, 2, 2, 2, 88, 93, 5, 16,
	9, 2, 89, 90, 7, 16, 2, 2, 90, 92, 5, 16, 9, 2, 91, 89, 3, 2, 2, 2, 92,
	95, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 15, 3, 2, 2,
	2, 95, 93, 3, 2, 2, 2, 96, 101, 5, 18, 10, 2, 97, 98, 7, 15, 2, 2, 98,
	100, 5, 18, 10, 2, 99, 97, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99, 3,
	2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 17, 3, 2, 2, 2, 103, 101, 3, 2, 2,
	2, 104, 138, 5, 22, 12, 2, 105, 106, 7, 17, 2, 2, 106, 138, 5, 18, 10,
	2, 107, 108, 5, 24, 13, 2, 108, 109, 5, 30, 16, 2, 109, 138, 3, 2, 2, 2,
	110, 111, 5, 24, 13, 2, 111, 112, 5, 28, 15, 2, 112, 113, 5, 24, 13, 2,
	113, 138, 3, 2, 2, 2, 114, 115, 5, 24, 13, 2, 115, 116, 9, 3, 2, 2, 116,
	119, 7, 32, 2, 2, 117, 120, 5, 24, 13, 2, 118, 120, 5, 20, 11, 2, 119,
	117, 3, 2, 2, 2, 119, 118, 3, 2, 2, 2, 120, 128, 3, 2, 2, 2, 121, 124,
	7, 34, 2, 2, 122, 125, 5, 24, 13, 2, 123, 125, 5, 20, 11, 2, 124, 122,
	3, 2, 2, 2, 124, 123, 3, 2, 2, 2, 125, 127, 3, 2, 2, 2, 126, 121, 3, 2,
	2, 2, 127, 130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2,
	129, 131, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 132, 7, 33, 2, 2, 132,
	138, 3, 2, 2, 2, 133, 134, 7, 32, 2, 2, 134, 135, 5, 12, 7, 2, 135, 136,
	7, 33, 2, 2, 136, 138, 3, 2, 2, 2, 137, 104, 3, 2, 2, 2, 137, 105, 3, 2,
	2, 2, 137, 107, 3, 2, 2, 2, 137, 110, 3, 2, 2, 2, 137, 114, 3, 2, 2, 2,
	137, 133, 3, 2, 2, 2, 138, 19, 3, 2, 2, 2, 139, 148, 7, 30, 2, 2, 140,
	145, 5, 24, 13, 2, 141, 142, 7, 34, 2, 2, 142, 144, 5, 24, 13, 2, 143,
	141, 3, 2, 2, 2, 144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146,
	3, 2, 2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148, 140, 3, 2,
	2, 2, 148, 149, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 151, 7, 31, 2, 2,
	151, 21, 3, 2, 2, 2, 152, 153, 7, 38, 2, 2, 153, 23, 3, 2, 2, 2, 154, 155,
	9, 4, 2, 2, 155, 25, 3, 2, 2, 2, 156, 157, 6, 14, 2, 2, 157, 159, 11, 2,
	2, 2, 158, 156, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2,
	160, 161, 3, 2, 2, 2, 161, 27, 3, 2, 2, 2, 162, 163, 9, 5, 2, 2, 163, 29,
	3, 2, 2, 2, 164, 165, 7, 29, 2, 2, 165, 31, 3, 2, 2, 2, 13, 36, 38, 93,
	101, 119, 124, 128, 137, 145, 148, 160,
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
	"policy", "prule", "pfilter", "pmacro", "plist", "expression", "or_expression",
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
	SfplParserRULE_prule           = 1
	SfplParserRULE_pfilter         = 2
	SfplParserRULE_pmacro          = 3
	SfplParserRULE_plist           = 4
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

func (s *PolicyContext) AllPrule() []IPruleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPruleContext)(nil)).Elem())
	var tst = make([]IPruleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPruleContext)
		}
	}

	return tst
}

func (s *PolicyContext) Prule(i int) IPruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPruleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPruleContext)
}

func (s *PolicyContext) AllPfilter() []IPfilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPfilterContext)(nil)).Elem())
	var tst = make([]IPfilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPfilterContext)
		}
	}

	return tst
}

func (s *PolicyContext) Pfilter(i int) IPfilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPfilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPfilterContext)
}

func (s *PolicyContext) AllPmacro() []IPmacroContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPmacroContext)(nil)).Elem())
	var tst = make([]IPmacroContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPmacroContext)
		}
	}

	return tst
}

func (s *PolicyContext) Pmacro(i int) IPmacroContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmacroContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPmacroContext)
}

func (s *PolicyContext) AllPlist() []IPlistContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPlistContext)(nil)).Elem())
	var tst = make([]IPlistContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPlistContext)
		}
	}

	return tst
}

func (s *PolicyContext) Plist(i int) IPlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlistContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPlistContext)
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
				p.Prule()
			}

		case 2:
			{
				p.SetState(31)
				p.Pfilter()
			}

		case 3:
			{
				p.SetState(32)
				p.Pmacro()
			}

		case 4:
			{
				p.SetState(33)
				p.Plist()
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

// IPruleContext is an interface to support dynamic dispatch.
type IPruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPruleContext differentiates from other interfaces.
	IsPruleContext()
}

type PruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPruleContext() *PruleContext {
	var p = new(PruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_prule
	return p
}

func (*PruleContext) IsPruleContext() {}

func NewPruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PruleContext {
	var p = new(PruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_prule

	return p
}

func (s *PruleContext) GetParser() antlr.Parser { return s.parser }

func (s *PruleContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *PruleContext) RULE() antlr.TerminalNode {
	return s.GetToken(SfplParserRULE, 0)
}

func (s *PruleContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *PruleContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *PruleContext) AllText() []ITextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITextContext)(nil)).Elem())
	var tst = make([]ITextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITextContext)
		}
	}

	return tst
}

func (s *PruleContext) Text(i int) ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *PruleContext) DESC() antlr.TerminalNode {
	return s.GetToken(SfplParserDESC, 0)
}

func (s *PruleContext) COND() antlr.TerminalNode {
	return s.GetToken(SfplParserCOND, 0)
}

func (s *PruleContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PruleContext) PRIORITY() antlr.TerminalNode {
	return s.GetToken(SfplParserPRIORITY, 0)
}

func (s *PruleContext) SEVERITY() antlr.TerminalNode {
	return s.GetToken(SfplParserSEVERITY, 0)
}

func (s *PruleContext) TAGS() antlr.TerminalNode {
	return s.GetToken(SfplParserTAGS, 0)
}

func (s *PruleContext) Items() IItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemsContext)
}

func (s *PruleContext) ACTION() antlr.TerminalNode {
	return s.GetToken(SfplParserACTION, 0)
}

func (s *PruleContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(SfplParserOUTPUT, 0)
}

func (s *PruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPrule(s)
	}
}

func (s *PruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPrule(s)
	}
}

func (s *PruleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPrule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Prule() (localctx IPruleContext) {
	localctx = NewPruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SfplParserRULE_prule)
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

// IPfilterContext is an interface to support dynamic dispatch.
type IPfilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPfilterContext differentiates from other interfaces.
	IsPfilterContext()
}

type PfilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPfilterContext() *PfilterContext {
	var p = new(PfilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_pfilter
	return p
}

func (*PfilterContext) IsPfilterContext() {}

func NewPfilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PfilterContext {
	var p = new(PfilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_pfilter

	return p
}

func (s *PfilterContext) GetParser() antlr.Parser { return s.parser }

func (s *PfilterContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *PfilterContext) FILTER() antlr.TerminalNode {
	return s.GetToken(SfplParserFILTER, 0)
}

func (s *PfilterContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *PfilterContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *PfilterContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *PfilterContext) COND() antlr.TerminalNode {
	return s.GetToken(SfplParserCOND, 0)
}

func (s *PfilterContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PfilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PfilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PfilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPfilter(s)
	}
}

func (s *PfilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPfilter(s)
	}
}

func (s *PfilterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPfilter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Pfilter() (localctx IPfilterContext) {
	localctx = NewPfilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SfplParserRULE_pfilter)

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
		p.Match(SfplParserID)
	}
	{
		p.SetState(64)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(65)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(66)
		p.Expression()
	}

	return localctx
}

// IPmacroContext is an interface to support dynamic dispatch.
type IPmacroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPmacroContext differentiates from other interfaces.
	IsPmacroContext()
}

type PmacroContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPmacroContext() *PmacroContext {
	var p = new(PmacroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_pmacro
	return p
}

func (*PmacroContext) IsPmacroContext() {}

func NewPmacroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmacroContext {
	var p = new(PmacroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_pmacro

	return p
}

func (s *PmacroContext) GetParser() antlr.Parser { return s.parser }

func (s *PmacroContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *PmacroContext) MACRO() antlr.TerminalNode {
	return s.GetToken(SfplParserMACRO, 0)
}

func (s *PmacroContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *PmacroContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *PmacroContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *PmacroContext) COND() antlr.TerminalNode {
	return s.GetToken(SfplParserCOND, 0)
}

func (s *PmacroContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PmacroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmacroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PmacroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPmacro(s)
	}
}

func (s *PmacroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPmacro(s)
	}
}

func (s *PmacroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPmacro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Pmacro() (localctx IPmacroContext) {
	localctx = NewPmacroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SfplParserRULE_pmacro)

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
		p.SetState(68)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(69)
		p.Match(SfplParserMACRO)
	}
	{
		p.SetState(70)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(71)
		p.Match(SfplParserID)
	}
	{
		p.SetState(72)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(73)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(74)
		p.Expression()
	}

	return localctx
}

// IPlistContext is an interface to support dynamic dispatch.
type IPlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlistContext differentiates from other interfaces.
	IsPlistContext()
}

type PlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlistContext() *PlistContext {
	var p = new(PlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_plist
	return p
}

func (*PlistContext) IsPlistContext() {}

func NewPlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlistContext {
	var p = new(PlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_plist

	return p
}

func (s *PlistContext) GetParser() antlr.Parser { return s.parser }

func (s *PlistContext) DECL() antlr.TerminalNode {
	return s.GetToken(SfplParserDECL, 0)
}

func (s *PlistContext) LIST() antlr.TerminalNode {
	return s.GetToken(SfplParserLIST, 0)
}

func (s *PlistContext) AllDEF() []antlr.TerminalNode {
	return s.GetTokens(SfplParserDEF)
}

func (s *PlistContext) DEF(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserDEF, i)
}

func (s *PlistContext) ID() antlr.TerminalNode {
	return s.GetToken(SfplParserID, 0)
}

func (s *PlistContext) ITEMS() antlr.TerminalNode {
	return s.GetToken(SfplParserITEMS, 0)
}

func (s *PlistContext) Items() IItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemsContext)
}

func (s *PlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPlist(s)
	}
}

func (s *PlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPlist(s)
	}
}

func (s *PlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Plist() (localctx IPlistContext) {
	localctx = NewPlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SfplParserRULE_plist)

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
		p.SetState(76)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(77)
		p.Match(SfplParserLIST)
	}
	{
		p.SetState(78)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(79)
		p.Match(SfplParserID)
	}
	{
		p.SetState(80)
		p.Match(SfplParserITEMS)
	}
	{
		p.SetState(81)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(82)
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
		p.SetState(84)
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
		p.SetState(86)
		p.And_expression()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SfplParserOR {
		{
			p.SetState(87)
			p.Match(SfplParserOR)
		}
		{
			p.SetState(88)
			p.And_expression()
		}

		p.SetState(93)
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
		p.SetState(94)
		p.Term()
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SfplParserAND {
		{
			p.SetState(95)
			p.Match(SfplParserAND)
		}
		{
			p.SetState(96)
			p.Term()
		}

		p.SetState(101)
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

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(SfplParserNOT)
		}
		{
			p.SetState(104)
			p.Term()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.Atom()
		}
		{
			p.SetState(106)
			p.Unary_operator()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(108)
			p.Atom()
		}
		{
			p.SetState(109)
			p.Binary_operator()
		}
		{
			p.SetState(110)
			p.Atom()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.Atom()
		}
		{
			p.SetState(113)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SfplParserIN || _la == SfplParserPMATCH) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(114)
			p.Match(SfplParserLPAREN)
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SfplParserLT, SfplParserGT, SfplParserID, SfplParserNUMBER, SfplParserPATH, SfplParserSTRING:
			{
				p.SetState(115)
				p.Atom()
			}

		case SfplParserLBRACK:
			{
				p.SetState(116)
				p.Items()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SfplParserLISTSEP {
			{
				p.SetState(119)
				p.Match(SfplParserLISTSEP)
			}
			p.SetState(122)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case SfplParserLT, SfplParserGT, SfplParserID, SfplParserNUMBER, SfplParserPATH, SfplParserSTRING:
				{
					p.SetState(120)
					p.Atom()
				}

			case SfplParserLBRACK:
				{
					p.SetState(121)
					p.Items()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(129)
			p.Match(SfplParserRPAREN)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(131)
			p.Match(SfplParserLPAREN)
		}
		{
			p.SetState(132)
			p.Expression()
		}
		{
			p.SetState(133)
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
		p.SetState(137)
		p.Match(SfplParserLBRACK)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(SfplParserLT-16))|(1<<(SfplParserGT-16))|(1<<(SfplParserID-16))|(1<<(SfplParserNUMBER-16))|(1<<(SfplParserPATH-16))|(1<<(SfplParserSTRING-16)))) != 0 {
		{
			p.SetState(138)
			p.Atom()
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SfplParserLISTSEP {
			{
				p.SetState(139)
				p.Match(SfplParserLISTSEP)
			}
			{
				p.SetState(140)
				p.Atom()
			}

			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(148)
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
		p.SetState(150)
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
		p.SetState(152)
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
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(154)

			if !(!(p.GetCurrentToken().GetText() == "desc" ||
				p.GetCurrentToken().GetText() == "condition" ||
				p.GetCurrentToken().GetText() == "action" ||
				p.GetCurrentToken().GetText() == "output" ||
				p.GetCurrentToken().GetText() == "priority" ||
				p.GetCurrentToken().GetText() == "tags")) {
				panic(antlr.NewFailedPredicateException(p, "!(p.GetCurrentToken().GetText() == \"desc\" ||\n\t      p.GetCurrentToken().GetText() == \"condition\" ||\n\t      p.GetCurrentToken().GetText() == \"action\" ||\n\t      p.GetCurrentToken().GetText() == \"output\" ||\n\t      p.GetCurrentToken().GetText() == \"priority\" ||\n\t      p.GetCurrentToken().GetText() == \"tags\")", ""))
			}
			p.SetState(155)
			p.MatchWildcard()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(158)
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
		p.SetState(160)
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
		p.SetState(162)
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
