// Code generated from QuACParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr4 // QuACParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQuACParserListener is a complete listener for a parse tree produced by QuACParser.
type BaseQuACParserListener struct{}

var _ QuACParserListener = &BaseQuACParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQuACParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQuACParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQuACParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQuACParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseQuACParserListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseQuACParserListener) ExitParse(ctx *ParseContext) {}

// EnterIFormatStatement is called when production iFormatStatement is entered.
func (s *BaseQuACParserListener) EnterIFormatStatement(ctx *IFormatStatementContext) {}

// ExitIFormatStatement is called when production iFormatStatement is exited.
func (s *BaseQuACParserListener) ExitIFormatStatement(ctx *IFormatStatementContext) {}

// EnterRMemFormatStatement is called when production rMemFormatStatement is entered.
func (s *BaseQuACParserListener) EnterRMemFormatStatement(ctx *RMemFormatStatementContext) {}

// ExitRMemFormatStatement is called when production rMemFormatStatement is exited.
func (s *BaseQuACParserListener) ExitRMemFormatStatement(ctx *RMemFormatStatementContext) {}

// EnterRALUFormatStatement is called when production rALUFormatStatement is entered.
func (s *BaseQuACParserListener) EnterRALUFormatStatement(ctx *RALUFormatStatementContext) {}

// ExitRALUFormatStatement is called when production rALUFormatStatement is exited.
func (s *BaseQuACParserListener) ExitRALUFormatStatement(ctx *RALUFormatStatementContext) {}

// EnterNopStatement is called when production nopStatement is entered.
func (s *BaseQuACParserListener) EnterNopStatement(ctx *NopStatementContext) {}

// ExitNopStatement is called when production nopStatement is exited.
func (s *BaseQuACParserListener) ExitNopStatement(ctx *NopStatementContext) {}

// EnterPseudo2ParamStatement is called when production pseudo2ParamStatement is entered.
func (s *BaseQuACParserListener) EnterPseudo2ParamStatement(ctx *Pseudo2ParamStatementContext) {}

// ExitPseudo2ParamStatement is called when production pseudo2ParamStatement is exited.
func (s *BaseQuACParserListener) ExitPseudo2ParamStatement(ctx *Pseudo2ParamStatementContext) {}

// EnterJprStatement is called when production jprStatement is entered.
func (s *BaseQuACParserListener) EnterJprStatement(ctx *JprStatementContext) {}

// ExitJprStatement is called when production jprStatement is exited.
func (s *BaseQuACParserListener) ExitJprStatement(ctx *JprStatementContext) {}

// EnterJpmStatement is called when production jpmStatement is entered.
func (s *BaseQuACParserListener) EnterJpmStatement(ctx *JpmStatementContext) {}

// ExitJpmStatement is called when production jpmStatement is exited.
func (s *BaseQuACParserListener) ExitJpmStatement(ctx *JpmStatementContext) {}

// EnterJpStatement is called when production jpStatement is entered.
func (s *BaseQuACParserListener) EnterJpStatement(ctx *JpStatementContext) {}

// ExitJpStatement is called when production jpStatement is exited.
func (s *BaseQuACParserListener) ExitJpStatement(ctx *JpStatementContext) {}

// EnterLabelStatement is called when production labelStatement is entered.
func (s *BaseQuACParserListener) EnterLabelStatement(ctx *LabelStatementContext) {}

// ExitLabelStatement is called when production labelStatement is exited.
func (s *BaseQuACParserListener) ExitLabelStatement(ctx *LabelStatementContext) {}

// EnterIFormat is called when production iFormat is entered.
func (s *BaseQuACParserListener) EnterIFormat(ctx *IFormatContext) {}

// ExitIFormat is called when production iFormat is exited.
func (s *BaseQuACParserListener) ExitIFormat(ctx *IFormatContext) {}

// EnterRMemFormat is called when production rMemFormat is entered.
func (s *BaseQuACParserListener) EnterRMemFormat(ctx *RMemFormatContext) {}

// ExitRMemFormat is called when production rMemFormat is exited.
func (s *BaseQuACParserListener) ExitRMemFormat(ctx *RMemFormatContext) {}

// EnterRALUFormat is called when production rALUFormat is entered.
func (s *BaseQuACParserListener) EnterRALUFormat(ctx *RALUFormatContext) {}

// ExitRALUFormat is called when production rALUFormat is exited.
func (s *BaseQuACParserListener) ExitRALUFormat(ctx *RALUFormatContext) {}

// EnterNop is called when production nop is entered.
func (s *BaseQuACParserListener) EnterNop(ctx *NopContext) {}

// ExitNop is called when production nop is exited.
func (s *BaseQuACParserListener) ExitNop(ctx *NopContext) {}

// EnterPseudo2Param is called when production pseudo2Param is entered.
func (s *BaseQuACParserListener) EnterPseudo2Param(ctx *Pseudo2ParamContext) {}

// ExitPseudo2Param is called when production pseudo2Param is exited.
func (s *BaseQuACParserListener) ExitPseudo2Param(ctx *Pseudo2ParamContext) {}

// EnterJpr is called when production jpr is entered.
func (s *BaseQuACParserListener) EnterJpr(ctx *JprContext) {}

// ExitJpr is called when production jpr is exited.
func (s *BaseQuACParserListener) ExitJpr(ctx *JprContext) {}

// EnterJpm is called when production jpm is entered.
func (s *BaseQuACParserListener) EnterJpm(ctx *JpmContext) {}

// ExitJpm is called when production jpm is exited.
func (s *BaseQuACParserListener) ExitJpm(ctx *JpmContext) {}

// EnterJp is called when production jp is entered.
func (s *BaseQuACParserListener) EnterJp(ctx *JpContext) {}

// ExitJp is called when production jp is exited.
func (s *BaseQuACParserListener) ExitJp(ctx *JpContext) {}

// EnterRegister is called when production register is entered.
func (s *BaseQuACParserListener) EnterRegister(ctx *RegisterContext) {}

// ExitRegister is called when production register is exited.
func (s *BaseQuACParserListener) ExitRegister(ctx *RegisterContext) {}
