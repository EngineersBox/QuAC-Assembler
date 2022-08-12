package insn

import "github.com/EngineersBox/QuAC-Compiler/antlr4"

var (
	offset uint16 = 0
)

type LabelListener struct {
	*antlr4.BaseQuACParserListener
	Labels map[string]uint16
}

func NewLabelListener() *LabelListener {
	var listener *LabelListener = new(LabelListener)
	listener.Labels = make(map[string]uint16)
	return listener
}

func (this *LabelListener) EnterIFormatStatement(ctx *antlr4.IFormatStatementContext) {
	offset += 16
}

func (this *LabelListener) EnterRMemFormatStatement(ctx *antlr4.RMemFormatStatementContext) {
	offset += 16
}

func (this *LabelListener) EnterRALUFormatStatement(ctx *antlr4.RALUFormatStatementContext) {
	offset += 16
}

func (this *LabelListener) EnterNopStatement(ctx *antlr4.NopStatementContext) {
	offset += 16
}

func (this *LabelListener) EnterPseudo2ParamStatement(ctx *antlr4.Pseudo2ParamStatementContext) {
	offset += 16
}

func (this *LabelListener) EnterJprStatement(ctx *antlr4.JprStatementContext) {
	offset += 16
}

func (this *LabelListener) EnterJpmStatement(ctx *antlr4.JpmStatementContext) {
	offset += 16
}

func (this *LabelListener) EnterJpStatement(ctx *antlr4.JpStatementContext) {
	offset += 16
}

func (this *LabelListener) EnterWordStatement(ctx *antlr4.WordStatementContext) {
	offset += 16
}

func (this *LabelListener) EnterLabelStatement(ctx *antlr4.LabelStatementContext) {
	this.Labels[ctx.Identifier().GetText()] = offset
}
