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
// Code generated from Sfpl.g4 by ANTLR 4.7.2. DO NOT EDIT.
//
package parser // Sfpl
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSfplVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSfplVisitor) VisitPolicy(ctx *PolicyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitPrule(ctx *PruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitPfilter(ctx *PfilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitPmacro(ctx *PmacroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSfplVisitor) VisitPlist(ctx *PlistContext) interface{} {
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
