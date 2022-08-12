// Code generated from QuACParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr4 // QuACParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by QuACParser.
type QuACParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QuACParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by QuACParser#iFormatStatement.
	VisitIFormatStatement(ctx *IFormatStatementContext) interface{}

	// Visit a parse tree produced by QuACParser#rMemFormatStatement.
	VisitRMemFormatStatement(ctx *RMemFormatStatementContext) interface{}

	// Visit a parse tree produced by QuACParser#rALUFormatStatement.
	VisitRALUFormatStatement(ctx *RALUFormatStatementContext) interface{}

	// Visit a parse tree produced by QuACParser#nopStatement.
	VisitNopStatement(ctx *NopStatementContext) interface{}

	// Visit a parse tree produced by QuACParser#pseudo2ParamStatement.
	VisitPseudo2ParamStatement(ctx *Pseudo2ParamStatementContext) interface{}

	// Visit a parse tree produced by QuACParser#jprStatement.
	VisitJprStatement(ctx *JprStatementContext) interface{}

	// Visit a parse tree produced by QuACParser#jpmStatement.
	VisitJpmStatement(ctx *JpmStatementContext) interface{}

	// Visit a parse tree produced by QuACParser#jpStatement.
	VisitJpStatement(ctx *JpStatementContext) interface{}

	// Visit a parse tree produced by QuACParser#wordStatement.
	VisitWordStatement(ctx *WordStatementContext) interface{}

	// Visit a parse tree produced by QuACParser#labelStatement.
	VisitLabelStatement(ctx *LabelStatementContext) interface{}

	// Visit a parse tree produced by QuACParser#iFormat.
	VisitIFormat(ctx *IFormatContext) interface{}

	// Visit a parse tree produced by QuACParser#rMemFormat.
	VisitRMemFormat(ctx *RMemFormatContext) interface{}

	// Visit a parse tree produced by QuACParser#rALUFormat.
	VisitRALUFormat(ctx *RALUFormatContext) interface{}

	// Visit a parse tree produced by QuACParser#nop.
	VisitNop(ctx *NopContext) interface{}

	// Visit a parse tree produced by QuACParser#pseudo2Param.
	VisitPseudo2Param(ctx *Pseudo2ParamContext) interface{}

	// Visit a parse tree produced by QuACParser#jpr.
	VisitJpr(ctx *JprContext) interface{}

	// Visit a parse tree produced by QuACParser#jpm.
	VisitJpm(ctx *JpmContext) interface{}

	// Visit a parse tree produced by QuACParser#jp.
	VisitJp(ctx *JpContext) interface{}

	// Visit a parse tree produced by QuACParser#register.
	VisitRegister(ctx *RegisterContext) interface{}
}
