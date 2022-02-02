// Code generated from Sfpl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sfpl
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSfplVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSfplVisitor) VisitPolicy(ctx *PolicyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitDefs(ctx *DefsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitPrule(ctx *PruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitSrule(ctx *SruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitPfilter(ctx *PfilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitSfilter(ctx *SfilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitDrop_keyword(ctx *Drop_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitPmacro(ctx *PmacroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitPlist(ctx *PlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitPreq(ctx *PreqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitOr_expression(ctx *Or_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitAnd_expression(ctx *And_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitItems(ctx *ItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitTags(ctx *TagsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitPrefilter(ctx *PrefilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitSeverity(ctx *SeverityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitEnabled(ctx *EnabledContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitWarnevttype(ctx *WarnevttypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitSkipunknown(ctx *SkipunknownContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitFappend(ctx *FappendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitText(ctx *TextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitBinary_operator(ctx *Binary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitUnary_operator(ctx *Unary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}
