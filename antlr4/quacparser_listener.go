// Code generated from QuACParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr4 // QuACParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// QuACParserListener is a complete listener for a parse tree produced by QuACParser.
type QuACParserListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterIFormatStatement is called when entering the iFormatStatement production.
	EnterIFormatStatement(c *IFormatStatementContext)

	// EnterRMemFormatStatement is called when entering the rMemFormatStatement production.
	EnterRMemFormatStatement(c *RMemFormatStatementContext)

	// EnterRALUFormatStatement is called when entering the rALUFormatStatement production.
	EnterRALUFormatStatement(c *RALUFormatStatementContext)

	// EnterNopStatement is called when entering the nopStatement production.
	EnterNopStatement(c *NopStatementContext)

	// EnterPseudo2ParamStatement is called when entering the pseudo2ParamStatement production.
	EnterPseudo2ParamStatement(c *Pseudo2ParamStatementContext)

	// EnterJprStatement is called when entering the jprStatement production.
	EnterJprStatement(c *JprStatementContext)

	// EnterJpmStatement is called when entering the jpmStatement production.
	EnterJpmStatement(c *JpmStatementContext)

	// EnterJpStatement is called when entering the jpStatement production.
	EnterJpStatement(c *JpStatementContext)

	// EnterLabelStatement is called when entering the labelStatement production.
	EnterLabelStatement(c *LabelStatementContext)

	// EnterIFormat is called when entering the iFormat production.
	EnterIFormat(c *IFormatContext)

	// EnterRMemFormat is called when entering the rMemFormat production.
	EnterRMemFormat(c *RMemFormatContext)

	// EnterRALUFormat is called when entering the rALUFormat production.
	EnterRALUFormat(c *RALUFormatContext)

	// EnterNop is called when entering the nop production.
	EnterNop(c *NopContext)

	// EnterPseudo2Param is called when entering the pseudo2Param production.
	EnterPseudo2Param(c *Pseudo2ParamContext)

	// EnterJpr is called when entering the jpr production.
	EnterJpr(c *JprContext)

	// EnterJpm is called when entering the jpm production.
	EnterJpm(c *JpmContext)

	// EnterJp is called when entering the jp production.
	EnterJp(c *JpContext)

	// EnterRegister is called when entering the register production.
	EnterRegister(c *RegisterContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitIFormatStatement is called when exiting the iFormatStatement production.
	ExitIFormatStatement(c *IFormatStatementContext)

	// ExitRMemFormatStatement is called when exiting the rMemFormatStatement production.
	ExitRMemFormatStatement(c *RMemFormatStatementContext)

	// ExitRALUFormatStatement is called when exiting the rALUFormatStatement production.
	ExitRALUFormatStatement(c *RALUFormatStatementContext)

	// ExitNopStatement is called when exiting the nopStatement production.
	ExitNopStatement(c *NopStatementContext)

	// ExitPseudo2ParamStatement is called when exiting the pseudo2ParamStatement production.
	ExitPseudo2ParamStatement(c *Pseudo2ParamStatementContext)

	// ExitJprStatement is called when exiting the jprStatement production.
	ExitJprStatement(c *JprStatementContext)

	// ExitJpmStatement is called when exiting the jpmStatement production.
	ExitJpmStatement(c *JpmStatementContext)

	// ExitJpStatement is called when exiting the jpStatement production.
	ExitJpStatement(c *JpStatementContext)

	// ExitLabelStatement is called when exiting the labelStatement production.
	ExitLabelStatement(c *LabelStatementContext)

	// ExitIFormat is called when exiting the iFormat production.
	ExitIFormat(c *IFormatContext)

	// ExitRMemFormat is called when exiting the rMemFormat production.
	ExitRMemFormat(c *RMemFormatContext)

	// ExitRALUFormat is called when exiting the rALUFormat production.
	ExitRALUFormat(c *RALUFormatContext)

	// ExitNop is called when exiting the nop production.
	ExitNop(c *NopContext)

	// ExitPseudo2Param is called when exiting the pseudo2Param production.
	ExitPseudo2Param(c *Pseudo2ParamContext)

	// ExitJpr is called when exiting the jpr production.
	ExitJpr(c *JprContext)

	// ExitJpm is called when exiting the jpm production.
	ExitJpm(c *JpmContext)

	// ExitJp is called when exiting the jp production.
	ExitJp(c *JpContext)

	// ExitRegister is called when exiting the register production.
	ExitRegister(c *RegisterContext)
}
