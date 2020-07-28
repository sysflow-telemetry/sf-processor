//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Code generated from Sfpl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Sfpl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSfplListener is a complete listener for a parse tree produced by SfplParser.
type BaseSfplListener struct{}

var _ SfplListener = &BaseSfplListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSfplListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSfplListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSfplListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSfplListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPolicy is called when production policy is entered.
func (s *BaseSfplListener) EnterPolicy(ctx *PolicyContext) {}

// ExitPolicy is called when production policy is exited.
func (s *BaseSfplListener) ExitPolicy(ctx *PolicyContext) {}

// EnterPrule is called when production prule is entered.
func (s *BaseSfplListener) EnterPrule(ctx *PruleContext) {}

// ExitPrule is called when production prule is exited.
func (s *BaseSfplListener) ExitPrule(ctx *PruleContext) {}

// EnterPfilter is called when production pfilter is entered.
func (s *BaseSfplListener) EnterPfilter(ctx *PfilterContext) {}

// ExitPfilter is called when production pfilter is exited.
func (s *BaseSfplListener) ExitPfilter(ctx *PfilterContext) {}

// EnterPmacro is called when production pmacro is entered.
func (s *BaseSfplListener) EnterPmacro(ctx *PmacroContext) {}

// ExitPmacro is called when production pmacro is exited.
func (s *BaseSfplListener) ExitPmacro(ctx *PmacroContext) {}

// EnterPlist is called when production plist is entered.
func (s *BaseSfplListener) EnterPlist(ctx *PlistContext) {}

// ExitPlist is called when production plist is exited.
func (s *BaseSfplListener) ExitPlist(ctx *PlistContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSfplListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSfplListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOr_expression is called when production or_expression is entered.
func (s *BaseSfplListener) EnterOr_expression(ctx *Or_expressionContext) {}

// ExitOr_expression is called when production or_expression is exited.
func (s *BaseSfplListener) ExitOr_expression(ctx *Or_expressionContext) {}

// EnterAnd_expression is called when production and_expression is entered.
func (s *BaseSfplListener) EnterAnd_expression(ctx *And_expressionContext) {}

// ExitAnd_expression is called when production and_expression is exited.
func (s *BaseSfplListener) ExitAnd_expression(ctx *And_expressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseSfplListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseSfplListener) ExitTerm(ctx *TermContext) {}

// EnterItems is called when production items is entered.
func (s *BaseSfplListener) EnterItems(ctx *ItemsContext) {}

// ExitItems is called when production items is exited.
func (s *BaseSfplListener) ExitItems(ctx *ItemsContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSfplListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSfplListener) ExitVariable(ctx *VariableContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseSfplListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseSfplListener) ExitAtom(ctx *AtomContext) {}

// EnterText is called when production text is entered.
func (s *BaseSfplListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseSfplListener) ExitText(ctx *TextContext) {}

// EnterBinary_operator is called when production binary_operator is entered.
func (s *BaseSfplListener) EnterBinary_operator(ctx *Binary_operatorContext) {}

// ExitBinary_operator is called when production binary_operator is exited.
func (s *BaseSfplListener) ExitBinary_operator(ctx *Binary_operatorContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseSfplListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseSfplListener) ExitUnary_operator(ctx *Unary_operatorContext) {}
