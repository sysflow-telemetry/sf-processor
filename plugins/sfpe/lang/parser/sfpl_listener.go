// Code generated from Sfpl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Sfpl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SfplListener is a complete listener for a parse tree produced by SfplParser.
type SfplListener interface {
	antlr.ParseTreeListener

	// EnterPolicy is called when entering the policy production.
	EnterPolicy(c *PolicyContext)

	// EnterF_rule is called when entering the f_rule production.
	EnterF_rule(c *F_ruleContext)

	// EnterF_filter is called when entering the f_filter production.
	EnterF_filter(c *F_filterContext)

	// EnterF_macro is called when entering the f_macro production.
	EnterF_macro(c *F_macroContext)

	// EnterF_list is called when entering the f_list production.
	EnterF_list(c *F_listContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOr_expression is called when entering the or_expression production.
	EnterOr_expression(c *Or_expressionContext)

	// EnterAnd_expression is called when entering the and_expression production.
	EnterAnd_expression(c *And_expressionContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterItems is called when entering the items production.
	EnterItems(c *ItemsContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterBinary_operator is called when entering the binary_operator production.
	EnterBinary_operator(c *Binary_operatorContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// ExitPolicy is called when exiting the policy production.
	ExitPolicy(c *PolicyContext)

	// ExitF_rule is called when exiting the f_rule production.
	ExitF_rule(c *F_ruleContext)

	// ExitF_filter is called when exiting the f_filter production.
	ExitF_filter(c *F_filterContext)

	// ExitF_macro is called when exiting the f_macro production.
	ExitF_macro(c *F_macroContext)

	// ExitF_list is called when exiting the f_list production.
	ExitF_list(c *F_listContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOr_expression is called when exiting the or_expression production.
	ExitOr_expression(c *Or_expressionContext)

	// ExitAnd_expression is called when exiting the and_expression production.
	ExitAnd_expression(c *And_expressionContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitItems is called when exiting the items production.
	ExitItems(c *ItemsContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitBinary_operator is called when exiting the binary_operator production.
	ExitBinary_operator(c *Binary_operatorContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)
}
