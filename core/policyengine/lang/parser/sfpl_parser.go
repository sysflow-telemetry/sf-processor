// Code generated from Sfpl.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 222,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 3, 2,
	3, 2, 3, 2, 6, 2, 49, 10, 2, 13, 2, 14, 2, 50, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 86, 10, 3, 12, 3, 14, 3, 89, 11, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 101,
	10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 124,
	10, 8, 12, 8, 14, 8, 127, 11, 8, 3, 9, 3, 9, 3, 9, 7, 9, 132, 10, 9, 12,
	9, 14, 9, 135, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 152, 10,
	10, 3, 10, 3, 10, 3, 10, 5, 10, 157, 10, 10, 7, 10, 159, 10, 10, 12, 10,
	14, 10, 162, 11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 170,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 176, 10, 11, 12, 11, 14, 11,
	179, 11, 11, 5, 11, 181, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 7, 12, 189, 10, 12, 12, 12, 14, 12, 192, 11, 12, 5, 12, 194, 10, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 6, 20, 214, 10, 20,
	13, 20, 14, 20, 215, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 2, 2, 23, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	2, 6, 3, 2, 11, 12, 4, 2, 28, 28, 32, 32, 5, 2, 22, 22, 24, 24, 44, 48,
	4, 2, 22, 27, 29, 31, 2, 225, 2, 48, 3, 2, 2, 2, 4, 54, 3, 2, 2, 2, 6,
	90, 3, 2, 2, 2, 8, 102, 3, 2, 2, 2, 10, 110, 3, 2, 2, 2, 12, 118, 3, 2,
	2, 2, 14, 120, 3, 2, 2, 2, 16, 128, 3, 2, 2, 2, 18, 169, 3, 2, 2, 2, 20,
	171, 3, 2, 2, 2, 22, 184, 3, 2, 2, 2, 24, 197, 3, 2, 2, 2, 26, 199, 3,
	2, 2, 2, 28, 201, 3, 2, 2, 2, 30, 203, 3, 2, 2, 2, 32, 205, 3, 2, 2, 2,
	34, 207, 3, 2, 2, 2, 36, 209, 3, 2, 2, 2, 38, 213, 3, 2, 2, 2, 40, 217,
	3, 2, 2, 2, 42, 219, 3, 2, 2, 2, 44, 49, 5, 4, 3, 2, 45, 49, 5, 6, 4, 2,
	46, 49, 5, 8, 5, 2, 47, 49, 5, 10, 6, 2, 48, 44, 3, 2, 2, 2, 48, 45, 3,
	2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50,
	48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 53, 7, 2, 2,
	3, 53, 3, 3, 2, 2, 2, 54, 55, 7, 39, 2, 2, 55, 56, 7, 3, 2, 2, 56, 57,
	7, 40, 2, 2, 57, 58, 5, 38, 20, 2, 58, 59, 7, 10, 2, 2, 59, 60, 7, 40,
	2, 2, 60, 61, 5, 38, 20, 2, 61, 62, 7, 9, 2, 2, 62, 63, 7, 40, 2, 2, 63,
	64, 5, 12, 7, 2, 64, 65, 9, 2, 2, 2, 65, 66, 7, 40, 2, 2, 66, 67, 5, 38,
	20, 2, 67, 68, 7, 13, 2, 2, 68, 69, 7, 40, 2, 2, 69, 87, 5, 26, 14, 2,
	70, 71, 7, 14, 2, 2, 71, 72, 7, 40, 2, 2, 72, 86, 5, 22, 12, 2, 73, 74,
	7, 15, 2, 2, 74, 75, 7, 40, 2, 2, 75, 86, 5, 24, 13, 2, 76, 77, 7, 16,
	2, 2, 77, 78, 7, 40, 2, 2, 78, 86, 5, 28, 15, 2, 79, 80, 7, 17, 2, 2, 80,
	81, 7, 40, 2, 2, 81, 86, 5, 30, 16, 2, 82, 83, 7, 18, 2, 2, 83, 84, 7,
	40, 2, 2, 84, 86, 5, 32, 17, 2, 85, 70, 3, 2, 2, 2, 85, 73, 3, 2, 2, 2,
	85, 76, 3, 2, 2, 2, 85, 79, 3, 2, 2, 2, 85, 82, 3, 2, 2, 2, 86, 89, 3,
	2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 5, 3, 2, 2, 2, 89,
	87, 3, 2, 2, 2, 90, 91, 7, 39, 2, 2, 91, 92, 7, 4, 2, 2, 92, 93, 7, 40,
	2, 2, 93, 94, 7, 44, 2, 2, 94, 95, 7, 9, 2, 2, 95, 96, 7, 40, 2, 2, 96,
	100, 5, 12, 7, 2, 97, 98, 7, 16, 2, 2, 98, 99, 7, 40, 2, 2, 99, 101, 5,
	28, 15, 2, 100, 97, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 7, 3, 2, 2,
	2, 102, 103, 7, 39, 2, 2, 103, 104, 7, 5, 2, 2, 104, 105, 7, 40, 2, 2,
	105, 106, 7, 44, 2, 2, 106, 107, 7, 9, 2, 2, 107, 108, 7, 40, 2, 2, 108,
	109, 5, 12, 7, 2, 109, 9, 3, 2, 2, 2, 110, 111, 7, 39, 2, 2, 111, 112,
	7, 6, 2, 2, 112, 113, 7, 40, 2, 2, 113, 114, 7, 44, 2, 2, 114, 115, 7,
	8, 2, 2, 115, 116, 7, 40, 2, 2, 116, 117, 5, 20, 11, 2, 117, 11, 3, 2,
	2, 2, 118, 119, 5, 14, 8, 2, 119, 13, 3, 2, 2, 2, 120, 125, 5, 16, 9, 2,
	121, 122, 7, 20, 2, 2, 122, 124, 5, 16, 9, 2, 123, 121, 3, 2, 2, 2, 124,
	127, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 15, 3,
	2, 2, 2, 127, 125, 3, 2, 2, 2, 128, 133, 5, 18, 10, 2, 129, 130, 7, 19,
	2, 2, 130, 132, 5, 18, 10, 2, 131, 129, 3, 2, 2, 2, 132, 135, 3, 2, 2,
	2, 133, 131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 17, 3, 2, 2, 2, 135,
	133, 3, 2, 2, 2, 136, 170, 5, 34, 18, 2, 137, 138, 7, 21, 2, 2, 138, 170,
	5, 18, 10, 2, 139, 140, 5, 36, 19, 2, 140, 141, 5, 42, 22, 2, 141, 170,
	3, 2, 2, 2, 142, 143, 5, 36, 19, 2, 143, 144, 5, 40, 21, 2, 144, 145, 5,
	36, 19, 2, 145, 170, 3, 2, 2, 2, 146, 147, 5, 36, 19, 2, 147, 148, 9, 3,
	2, 2, 148, 151, 7, 36, 2, 2, 149, 152, 5, 36, 19, 2, 150, 152, 5, 20, 11,
	2, 151, 149, 3, 2, 2, 2, 151, 150, 3, 2, 2, 2, 152, 160, 3, 2, 2, 2, 153,
	156, 7, 38, 2, 2, 154, 157, 5, 36, 19, 2, 155, 157, 5, 20, 11, 2, 156,
	154, 3, 2, 2, 2, 156, 155, 3, 2, 2, 2, 157, 159, 3, 2, 2, 2, 158, 153,
	3, 2, 2, 2, 159, 162, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2,
	2, 2, 161, 163, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 163, 164, 7, 37, 2, 2,
	164, 170, 3, 2, 2, 2, 165, 166, 7, 36, 2, 2, 166, 167, 5, 12, 7, 2, 167,
	168, 7, 37, 2, 2, 168, 170, 3, 2, 2, 2, 169, 136, 3, 2, 2, 2, 169, 137,
	3, 2, 2, 2, 169, 139, 3, 2, 2, 2, 169, 142, 3, 2, 2, 2, 169, 146, 3, 2,
	2, 2, 169, 165, 3, 2, 2, 2, 170, 19, 3, 2, 2, 2, 171, 180, 7, 34, 2, 2,
	172, 177, 5, 36, 19, 2, 173, 174, 7, 38, 2, 2, 174, 176, 5, 36, 19, 2,
	175, 173, 3, 2, 2, 2, 176, 179, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177,
	178, 3, 2, 2, 2, 178, 181, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 180, 172,
	3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 183, 7, 35,
	2, 2, 183, 21, 3, 2, 2, 2, 184, 193, 7, 34, 2, 2, 185, 190, 5, 36, 19,
	2, 186, 187, 7, 38, 2, 2, 187, 189, 5, 36, 19, 2, 188, 186, 3, 2, 2, 2,
	189, 192, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191,
	194, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 193, 185, 3, 2, 2, 2, 193, 194,
	3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 7, 35, 2, 2, 196, 23, 3, 2,
	2, 2, 197, 198, 5, 20, 11, 2, 198, 25, 3, 2, 2, 2, 199, 200, 7, 41, 2,
	2, 200, 27, 3, 2, 2, 2, 201, 202, 5, 36, 19, 2, 202, 29, 3, 2, 2, 2, 203,
	204, 5, 36, 19, 2, 204, 31, 3, 2, 2, 2, 205, 206, 5, 36, 19, 2, 206, 33,
	3, 2, 2, 2, 207, 208, 7, 44, 2, 2, 208, 35, 3, 2, 2, 2, 209, 210, 9, 4,
	2, 2, 210, 37, 3, 2, 2, 2, 211, 212, 6, 20, 2, 2, 212, 214, 11, 2, 2, 2,
	213, 211, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215,
	216, 3, 2, 2, 2, 216, 39, 3, 2, 2, 2, 217, 218, 9, 5, 2, 2, 218, 41, 3,
	2, 2, 2, 219, 220, 7, 33, 2, 2, 220, 43, 3, 2, 2, 2, 18, 48, 50, 85, 87,
	100, 125, 133, 151, 156, 160, 169, 177, 180, 190, 193, 215,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'rule'", "'filter'", "'macro'", "'list'", "'name'", "'items'", "'condition'",
	"'desc'", "'action'", "'output'", "'priority'", "'tags'", "'prefilter'",
	"'enabled'", "'warn_evttypes'", "'skip-if-unknown-filter'", "'and'", "'or'",
	"'not'", "'<'", "'<='", "'>'", "'>='", "'='", "'!='", "'in'", "'contains'",
	"'icontains'", "'startswith'", "'pmatch'", "'exists'", "'['", "']'", "'('",
	"')'", "','", "'-'",
}
var symbolicNames = []string{
	"", "RULE", "FILTER", "MACRO", "LIST", "NAME", "ITEMS", "COND", "DESC",
	"ACTION", "OUTPUT", "PRIORITY", "TAGS", "PREFILTER", "ENABLED", "WARNEVTTYPE",
	"SKIPUNKNOWN", "AND", "OR", "NOT", "LT", "LE", "GT", "GE", "EQ", "NEQ",
	"IN", "CONTAINS", "ICONTAINS", "STARTSWITH", "PMATCH", "EXISTS", "LBRACK",
	"RBRACK", "LPAREN", "RPAREN", "LISTSEP", "DECL", "DEF", "SEVERITY", "SFSEVERITY",
	"FSEVERITY", "ID", "NUMBER", "PATH", "STRING", "TAG", "WS", "NL", "COMMENT",
	"ANY",
}

var ruleNames = []string{
	"policy", "prule", "pfilter", "pmacro", "plist", "expression", "or_expression",
	"and_expression", "term", "items", "tags", "prefilter", "severity", "enabled",
	"warnevttype", "skipunknown", "variable", "atom", "text", "binary_operator",
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
	SfplParserEOF         = antlr.TokenEOF
	SfplParserRULE        = 1
	SfplParserFILTER      = 2
	SfplParserMACRO       = 3
	SfplParserLIST        = 4
	SfplParserNAME        = 5
	SfplParserITEMS       = 6
	SfplParserCOND        = 7
	SfplParserDESC        = 8
	SfplParserACTION      = 9
	SfplParserOUTPUT      = 10
	SfplParserPRIORITY    = 11
	SfplParserTAGS        = 12
	SfplParserPREFILTER   = 13
	SfplParserENABLED     = 14
	SfplParserWARNEVTTYPE = 15
	SfplParserSKIPUNKNOWN = 16
	SfplParserAND         = 17
	SfplParserOR          = 18
	SfplParserNOT         = 19
	SfplParserLT          = 20
	SfplParserLE          = 21
	SfplParserGT          = 22
	SfplParserGE          = 23
	SfplParserEQ          = 24
	SfplParserNEQ         = 25
	SfplParserIN          = 26
	SfplParserCONTAINS    = 27
	SfplParserICONTAINS   = 28
	SfplParserSTARTSWITH  = 29
	SfplParserPMATCH      = 30
	SfplParserEXISTS      = 31
	SfplParserLBRACK      = 32
	SfplParserRBRACK      = 33
	SfplParserLPAREN      = 34
	SfplParserRPAREN      = 35
	SfplParserLISTSEP     = 36
	SfplParserDECL        = 37
	SfplParserDEF         = 38
	SfplParserSEVERITY    = 39
	SfplParserSFSEVERITY  = 40
	SfplParserFSEVERITY   = 41
	SfplParserID          = 42
	SfplParserNUMBER      = 43
	SfplParserPATH        = 44
	SfplParserSTRING      = 45
	SfplParserTAG         = 46
	SfplParserWS          = 47
	SfplParserNL          = 48
	SfplParserCOMMENT     = 49
	SfplParserANY         = 50
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
	SfplParserRULE_tags            = 10
	SfplParserRULE_prefilter       = 11
	SfplParserRULE_severity        = 12
	SfplParserRULE_enabled         = 13
	SfplParserRULE_warnevttype     = 14
	SfplParserRULE_skipunknown     = 15
	SfplParserRULE_variable        = 16
	SfplParserRULE_atom            = 17
	SfplParserRULE_text            = 18
	SfplParserRULE_binary_operator = 19
	SfplParserRULE_unary_operator  = 20
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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SfplParserDECL {
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(42)
				p.Prule()
			}

		case 2:
			{
				p.SetState(43)
				p.Pfilter()
			}

		case 3:
			{
				p.SetState(44)
				p.Pmacro()
			}

		case 4:
			{
				p.SetState(45)
				p.Plist()
			}

		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
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

func (s *PruleContext) Severity() ISeverityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeverityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeverityContext)
}

func (s *PruleContext) ACTION() antlr.TerminalNode {
	return s.GetToken(SfplParserACTION, 0)
}

func (s *PruleContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(SfplParserOUTPUT, 0)
}

func (s *PruleContext) AllTAGS() []antlr.TerminalNode {
	return s.GetTokens(SfplParserTAGS)
}

func (s *PruleContext) TAGS(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserTAGS, i)
}

func (s *PruleContext) AllTags() []ITagsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagsContext)(nil)).Elem())
	var tst = make([]ITagsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagsContext)
		}
	}

	return tst
}

func (s *PruleContext) Tags(i int) ITagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagsContext)
}

func (s *PruleContext) AllPREFILTER() []antlr.TerminalNode {
	return s.GetTokens(SfplParserPREFILTER)
}

func (s *PruleContext) PREFILTER(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserPREFILTER, i)
}

func (s *PruleContext) AllPrefilter() []IPrefilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrefilterContext)(nil)).Elem())
	var tst = make([]IPrefilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrefilterContext)
		}
	}

	return tst
}

func (s *PruleContext) Prefilter(i int) IPrefilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrefilterContext)
}

func (s *PruleContext) AllENABLED() []antlr.TerminalNode {
	return s.GetTokens(SfplParserENABLED)
}

func (s *PruleContext) ENABLED(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserENABLED, i)
}

func (s *PruleContext) AllEnabled() []IEnabledContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnabledContext)(nil)).Elem())
	var tst = make([]IEnabledContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnabledContext)
		}
	}

	return tst
}

func (s *PruleContext) Enabled(i int) IEnabledContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnabledContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnabledContext)
}

func (s *PruleContext) AllWARNEVTTYPE() []antlr.TerminalNode {
	return s.GetTokens(SfplParserWARNEVTTYPE)
}

func (s *PruleContext) WARNEVTTYPE(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserWARNEVTTYPE, i)
}

func (s *PruleContext) AllWarnevttype() []IWarnevttypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWarnevttypeContext)(nil)).Elem())
	var tst = make([]IWarnevttypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWarnevttypeContext)
		}
	}

	return tst
}

func (s *PruleContext) Warnevttype(i int) IWarnevttypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWarnevttypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWarnevttypeContext)
}

func (s *PruleContext) AllSKIPUNKNOWN() []antlr.TerminalNode {
	return s.GetTokens(SfplParserSKIPUNKNOWN)
}

func (s *PruleContext) SKIPUNKNOWN(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserSKIPUNKNOWN, i)
}

func (s *PruleContext) AllSkipunknown() []ISkipunknownContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISkipunknownContext)(nil)).Elem())
	var tst = make([]ISkipunknownContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISkipunknownContext)
		}
	}

	return tst
}

func (s *PruleContext) Skipunknown(i int) ISkipunknownContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISkipunknownContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISkipunknownContext)
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
		p.SetState(52)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(53)
		p.Match(SfplParserRULE)
	}
	{
		p.SetState(54)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(55)
		p.Text()
	}
	{
		p.SetState(56)
		p.Match(SfplParserDESC)
	}
	{
		p.SetState(57)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(58)
		p.Text()
	}
	{
		p.SetState(59)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(60)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(61)
		p.Expression()
	}
	{
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SfplParserACTION || _la == SfplParserOUTPUT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(63)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(64)
		p.Text()
	}
	{
		p.SetState(65)
		p.Match(SfplParserPRIORITY)
	}
	{
		p.SetState(66)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(67)
		p.Severity()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SfplParserTAGS)|(1<<SfplParserPREFILTER)|(1<<SfplParserENABLED)|(1<<SfplParserWARNEVTTYPE)|(1<<SfplParserSKIPUNKNOWN))) != 0 {
		p.SetState(83)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SfplParserTAGS:
			{
				p.SetState(68)
				p.Match(SfplParserTAGS)
			}
			{
				p.SetState(69)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(70)
				p.Tags()
			}

		case SfplParserPREFILTER:
			{
				p.SetState(71)
				p.Match(SfplParserPREFILTER)
			}
			{
				p.SetState(72)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(73)
				p.Prefilter()
			}

		case SfplParserENABLED:
			{
				p.SetState(74)
				p.Match(SfplParserENABLED)
			}
			{
				p.SetState(75)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(76)
				p.Enabled()
			}

		case SfplParserWARNEVTTYPE:
			{
				p.SetState(77)
				p.Match(SfplParserWARNEVTTYPE)
			}
			{
				p.SetState(78)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(79)
				p.Warnevttype()
			}

		case SfplParserSKIPUNKNOWN:
			{
				p.SetState(80)
				p.Match(SfplParserSKIPUNKNOWN)
			}
			{
				p.SetState(81)
				p.Match(SfplParserDEF)
			}
			{
				p.SetState(82)
				p.Skipunknown()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *PfilterContext) ENABLED() antlr.TerminalNode {
	return s.GetToken(SfplParserENABLED, 0)
}

func (s *PfilterContext) Enabled() IEnabledContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnabledContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnabledContext)
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
		p.SetState(88)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(89)
		p.Match(SfplParserFILTER)
	}
	{
		p.SetState(90)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(91)
		p.Match(SfplParserID)
	}
	{
		p.SetState(92)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(93)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(94)
		p.Expression()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SfplParserENABLED {
		{
			p.SetState(95)
			p.Match(SfplParserENABLED)
		}
		{
			p.SetState(96)
			p.Match(SfplParserDEF)
		}
		{
			p.SetState(97)
			p.Enabled()
		}

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
		p.SetState(100)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(101)
		p.Match(SfplParserMACRO)
	}
	{
		p.SetState(102)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(103)
		p.Match(SfplParserID)
	}
	{
		p.SetState(104)
		p.Match(SfplParserCOND)
	}
	{
		p.SetState(105)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(106)
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
		p.SetState(108)
		p.Match(SfplParserDECL)
	}
	{
		p.SetState(109)
		p.Match(SfplParserLIST)
	}
	{
		p.SetState(110)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(111)
		p.Match(SfplParserID)
	}
	{
		p.SetState(112)
		p.Match(SfplParserITEMS)
	}
	{
		p.SetState(113)
		p.Match(SfplParserDEF)
	}
	{
		p.SetState(114)
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
		p.SetState(116)
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
		p.SetState(118)
		p.And_expression()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SfplParserOR {
		{
			p.SetState(119)
			p.Match(SfplParserOR)
		}
		{
			p.SetState(120)
			p.And_expression()
		}

		p.SetState(125)
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
		p.SetState(126)
		p.Term()
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SfplParserAND {
		{
			p.SetState(127)
			p.Match(SfplParserAND)
		}
		{
			p.SetState(128)
			p.Term()
		}

		p.SetState(133)
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

	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(SfplParserNOT)
		}
		{
			p.SetState(136)
			p.Term()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)
			p.Atom()
		}
		{
			p.SetState(138)
			p.Unary_operator()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(140)
			p.Atom()
		}
		{
			p.SetState(141)
			p.Binary_operator()
		}
		{
			p.SetState(142)
			p.Atom()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(144)
			p.Atom()
		}
		{
			p.SetState(145)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SfplParserIN || _la == SfplParserPMATCH) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(146)
			p.Match(SfplParserLPAREN)
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SfplParserLT, SfplParserGT, SfplParserID, SfplParserNUMBER, SfplParserPATH, SfplParserSTRING, SfplParserTAG:
			{
				p.SetState(147)
				p.Atom()
			}

		case SfplParserLBRACK:
			{
				p.SetState(148)
				p.Items()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SfplParserLISTSEP {
			{
				p.SetState(151)
				p.Match(SfplParserLISTSEP)
			}
			p.SetState(154)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case SfplParserLT, SfplParserGT, SfplParserID, SfplParserNUMBER, SfplParserPATH, SfplParserSTRING, SfplParserTAG:
				{
					p.SetState(152)
					p.Atom()
				}

			case SfplParserLBRACK:
				{
					p.SetState(153)
					p.Items()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(160)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(161)
			p.Match(SfplParserRPAREN)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(163)
			p.Match(SfplParserLPAREN)
		}
		{
			p.SetState(164)
			p.Expression()
		}
		{
			p.SetState(165)
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
		p.SetState(169)
		p.Match(SfplParserLBRACK)
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(SfplParserLT-20))|(1<<(SfplParserGT-20))|(1<<(SfplParserID-20))|(1<<(SfplParserNUMBER-20))|(1<<(SfplParserPATH-20))|(1<<(SfplParserSTRING-20))|(1<<(SfplParserTAG-20)))) != 0 {
		{
			p.SetState(170)
			p.Atom()
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SfplParserLISTSEP {
			{
				p.SetState(171)
				p.Match(SfplParserLISTSEP)
			}
			{
				p.SetState(172)
				p.Atom()
			}

			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(180)
		p.Match(SfplParserRBRACK)
	}

	return localctx
}

// ITagsContext is an interface to support dynamic dispatch.
type ITagsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagsContext differentiates from other interfaces.
	IsTagsContext()
}

type TagsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagsContext() *TagsContext {
	var p = new(TagsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_tags
	return p
}

func (*TagsContext) IsTagsContext() {}

func NewTagsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagsContext {
	var p = new(TagsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_tags

	return p
}

func (s *TagsContext) GetParser() antlr.Parser { return s.parser }

func (s *TagsContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(SfplParserLBRACK, 0)
}

func (s *TagsContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(SfplParserRBRACK, 0)
}

func (s *TagsContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *TagsContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *TagsContext) AllLISTSEP() []antlr.TerminalNode {
	return s.GetTokens(SfplParserLISTSEP)
}

func (s *TagsContext) LISTSEP(i int) antlr.TerminalNode {
	return s.GetToken(SfplParserLISTSEP, i)
}

func (s *TagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterTags(s)
	}
}

func (s *TagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitTags(s)
	}
}

func (s *TagsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitTags(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Tags() (localctx ITagsContext) {
	localctx = NewTagsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SfplParserRULE_tags)
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
		p.SetState(182)
		p.Match(SfplParserLBRACK)
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(SfplParserLT-20))|(1<<(SfplParserGT-20))|(1<<(SfplParserID-20))|(1<<(SfplParserNUMBER-20))|(1<<(SfplParserPATH-20))|(1<<(SfplParserSTRING-20))|(1<<(SfplParserTAG-20)))) != 0 {
		{
			p.SetState(183)
			p.Atom()
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SfplParserLISTSEP {
			{
				p.SetState(184)
				p.Match(SfplParserLISTSEP)
			}
			{
				p.SetState(185)
				p.Atom()
			}

			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(193)
		p.Match(SfplParserRBRACK)
	}

	return localctx
}

// IPrefilterContext is an interface to support dynamic dispatch.
type IPrefilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefilterContext differentiates from other interfaces.
	IsPrefilterContext()
}

type PrefilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefilterContext() *PrefilterContext {
	var p = new(PrefilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_prefilter
	return p
}

func (*PrefilterContext) IsPrefilterContext() {}

func NewPrefilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefilterContext {
	var p = new(PrefilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_prefilter

	return p
}

func (s *PrefilterContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefilterContext) Items() IItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemsContext)
}

func (s *PrefilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterPrefilter(s)
	}
}

func (s *PrefilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitPrefilter(s)
	}
}

func (s *PrefilterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitPrefilter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Prefilter() (localctx IPrefilterContext) {
	localctx = NewPrefilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SfplParserRULE_prefilter)

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
		p.SetState(195)
		p.Items()
	}

	return localctx
}

// ISeverityContext is an interface to support dynamic dispatch.
type ISeverityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSeverityContext differentiates from other interfaces.
	IsSeverityContext()
}

type SeverityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeverityContext() *SeverityContext {
	var p = new(SeverityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_severity
	return p
}

func (*SeverityContext) IsSeverityContext() {}

func NewSeverityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeverityContext {
	var p = new(SeverityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_severity

	return p
}

func (s *SeverityContext) GetParser() antlr.Parser { return s.parser }

func (s *SeverityContext) SEVERITY() antlr.TerminalNode {
	return s.GetToken(SfplParserSEVERITY, 0)
}

func (s *SeverityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeverityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeverityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterSeverity(s)
	}
}

func (s *SeverityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitSeverity(s)
	}
}

func (s *SeverityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitSeverity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Severity() (localctx ISeverityContext) {
	localctx = NewSeverityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SfplParserRULE_severity)

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
		p.SetState(197)
		p.Match(SfplParserSEVERITY)
	}

	return localctx
}

// IEnabledContext is an interface to support dynamic dispatch.
type IEnabledContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnabledContext differentiates from other interfaces.
	IsEnabledContext()
}

type EnabledContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnabledContext() *EnabledContext {
	var p = new(EnabledContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_enabled
	return p
}

func (*EnabledContext) IsEnabledContext() {}

func NewEnabledContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnabledContext {
	var p = new(EnabledContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_enabled

	return p
}

func (s *EnabledContext) GetParser() antlr.Parser { return s.parser }

func (s *EnabledContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *EnabledContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnabledContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnabledContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterEnabled(s)
	}
}

func (s *EnabledContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitEnabled(s)
	}
}

func (s *EnabledContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitEnabled(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Enabled() (localctx IEnabledContext) {
	localctx = NewEnabledContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SfplParserRULE_enabled)

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
		p.SetState(199)
		p.Atom()
	}

	return localctx
}

// IWarnevttypeContext is an interface to support dynamic dispatch.
type IWarnevttypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWarnevttypeContext differentiates from other interfaces.
	IsWarnevttypeContext()
}

type WarnevttypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWarnevttypeContext() *WarnevttypeContext {
	var p = new(WarnevttypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_warnevttype
	return p
}

func (*WarnevttypeContext) IsWarnevttypeContext() {}

func NewWarnevttypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WarnevttypeContext {
	var p = new(WarnevttypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_warnevttype

	return p
}

func (s *WarnevttypeContext) GetParser() antlr.Parser { return s.parser }

func (s *WarnevttypeContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *WarnevttypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WarnevttypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WarnevttypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterWarnevttype(s)
	}
}

func (s *WarnevttypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitWarnevttype(s)
	}
}

func (s *WarnevttypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitWarnevttype(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Warnevttype() (localctx IWarnevttypeContext) {
	localctx = NewWarnevttypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SfplParserRULE_warnevttype)

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
		p.SetState(201)
		p.Atom()
	}

	return localctx
}

// ISkipunknownContext is an interface to support dynamic dispatch.
type ISkipunknownContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSkipunknownContext differentiates from other interfaces.
	IsSkipunknownContext()
}

type SkipunknownContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySkipunknownContext() *SkipunknownContext {
	var p = new(SkipunknownContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SfplParserRULE_skipunknown
	return p
}

func (*SkipunknownContext) IsSkipunknownContext() {}

func NewSkipunknownContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SkipunknownContext {
	var p = new(SkipunknownContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SfplParserRULE_skipunknown

	return p
}

func (s *SkipunknownContext) GetParser() antlr.Parser { return s.parser }

func (s *SkipunknownContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *SkipunknownContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkipunknownContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SkipunknownContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.EnterSkipunknown(s)
	}
}

func (s *SkipunknownContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SfplListener); ok {
		listenerT.ExitSkipunknown(s)
	}
}

func (s *SkipunknownContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SfplVisitor:
		return t.VisitSkipunknown(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SfplParser) Skipunknown() (localctx ISkipunknownContext) {
	localctx = NewSkipunknownContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SfplParserRULE_skipunknown)

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
		p.SetState(203)
		p.Atom()
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
	p.EnterRule(localctx, 32, SfplParserRULE_variable)

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
		p.SetState(205)
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

func (s *AtomContext) TAG() antlr.TerminalNode {
	return s.GetToken(SfplParserTAG, 0)
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
	p.EnterRule(localctx, 34, SfplParserRULE_atom)
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
		p.SetState(207)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(SfplParserLT-20))|(1<<(SfplParserGT-20))|(1<<(SfplParserID-20))|(1<<(SfplParserNUMBER-20))|(1<<(SfplParserPATH-20))|(1<<(SfplParserSTRING-20))|(1<<(SfplParserTAG-20)))) != 0) {
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
	p.EnterRule(localctx, 36, SfplParserRULE_text)

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
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(209)

			if !(!(p.GetCurrentToken().GetText() == "desc" ||
				p.GetCurrentToken().GetText() == "condition" ||
				p.GetCurrentToken().GetText() == "action" ||
				p.GetCurrentToken().GetText() == "output" ||
				p.GetCurrentToken().GetText() == "priority" ||
				p.GetCurrentToken().GetText() == "tags" ||
				p.GetCurrentToken().GetText() == "prefilter" ||
				p.GetCurrentToken().GetText() == "enabled" ||
				p.GetCurrentToken().GetText() == "warn_evttypes" ||
				p.GetCurrentToken().GetText() == "skip-if-unknown-filter")) {
				panic(antlr.NewFailedPredicateException(p, "!(p.GetCurrentToken().GetText() == \"desc\" ||\n\t      p.GetCurrentToken().GetText() == \"condition\" ||\n\t      p.GetCurrentToken().GetText() == \"action\" ||\n\t      p.GetCurrentToken().GetText() == \"output\" ||\n\t      p.GetCurrentToken().GetText() == \"priority\" ||\n\t      p.GetCurrentToken().GetText() == \"tags\" ||\n\t\t  p.GetCurrentToken().GetText() == \"prefilter\" ||\n\t\t  p.GetCurrentToken().GetText() == \"enabled\" ||\n\t\t  p.GetCurrentToken().GetText() == \"warn_evttypes\" ||\n\t\t  p.GetCurrentToken().GetText() == \"skip-if-unknown-filter\")", ""))
			}
			p.SetState(210)
			p.MatchWildcard()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 38, SfplParserRULE_binary_operator)
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
		p.SetState(215)
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
	p.EnterRule(localctx, 40, SfplParserRULE_unary_operator)

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
		p.SetState(217)
		p.Match(SfplParserEXISTS)
	}

	return localctx
}

func (p *SfplParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
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
			p.GetCurrentToken().GetText() == "tags" ||
			p.GetCurrentToken().GetText() == "prefilter" ||
			p.GetCurrentToken().GetText() == "enabled" ||
			p.GetCurrentToken().GetText() == "warn_evttypes" ||
			p.GetCurrentToken().GetText() == "skip-if-unknown-filter")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
