// Code generated from QuACParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr4 // QuACParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type QuACParser struct {
	*antlr.BaseParser
}

var quacparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func quacparserParserInit() {
	staticData := &quacparserParserStaticData
	staticData.literalNames = []string{
		"", "", "'['", "']'", "';'", "':'", "'.'", "','", "'rz'", "'r0'", "'r1'",
		"'r2'", "'r3'", "'r4'", "'fl'", "'r5'", "'pc'", "'r7'", "'eq'", "'mov'",
		"'movl'", "'seth'", "'str'", "'ldr'", "'add'", "'sub'", "'and'", "'orr'",
		"'jpr'", "'cmp'", "'nop'", "'jpm'", "'jp'",
	}
	staticData.symbolicNames = []string{
		"", "IntegerLiteral", "LBRACK", "RBRACK", "SEMI", "COLON", "DOT", "COMMA",
		"RZ", "R0", "R1", "R2", "R3", "R4", "FL", "R5", "PC", "R7", "EQ", "MOV",
		"MOVL", "SETH", "STR", "LDR", "ADD", "SUB", "AND", "ORR", "JPR", "CMP",
		"NOP", "JPM", "JP", "Identifier", "WS", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"parse", "statement", "iFormat", "rMemFormat", "rALUFormat", "nop",
		"pseudo2Param", "jpr", "jpm", "jp", "register",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 35, 103, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 5, 0, 24, 8, 0, 10, 0, 12, 0, 27, 9, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 41, 8, 1, 1, 2,
		1, 2, 3, 2, 45, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 53, 8,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 63, 8, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 75, 8, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 3, 7, 83, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		3, 8, 89, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 97, 8, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 0, 0, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 0, 6, 1, 0, 20, 21, 1, 0, 22, 23, 1, 0, 24, 27, 2, 0, 19, 19, 29,
		29, 2, 0, 1, 1, 33, 33, 1, 0, 8, 17, 107, 0, 25, 1, 0, 0, 0, 2, 40, 1,
		0, 0, 0, 4, 42, 1, 0, 0, 0, 6, 50, 1, 0, 0, 0, 8, 60, 1, 0, 0, 0, 10, 70,
		1, 0, 0, 0, 12, 72, 1, 0, 0, 0, 14, 80, 1, 0, 0, 0, 16, 86, 1, 0, 0, 0,
		18, 94, 1, 0, 0, 0, 20, 100, 1, 0, 0, 0, 22, 24, 3, 2, 1, 0, 23, 22, 1,
		0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26,
		28, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29, 5, 0, 0, 1, 29, 1, 1, 0, 0,
		0, 30, 41, 3, 4, 2, 0, 31, 41, 3, 6, 3, 0, 32, 41, 3, 8, 4, 0, 33, 41,
		3, 10, 5, 0, 34, 41, 3, 12, 6, 0, 35, 41, 3, 14, 7, 0, 36, 41, 3, 16, 8,
		0, 37, 41, 3, 18, 9, 0, 38, 39, 5, 33, 0, 0, 39, 41, 5, 5, 0, 0, 40, 30,
		1, 0, 0, 0, 40, 31, 1, 0, 0, 0, 40, 32, 1, 0, 0, 0, 40, 33, 1, 0, 0, 0,
		40, 34, 1, 0, 0, 0, 40, 35, 1, 0, 0, 0, 40, 36, 1, 0, 0, 0, 40, 37, 1,
		0, 0, 0, 40, 38, 1, 0, 0, 0, 41, 3, 1, 0, 0, 0, 42, 44, 7, 0, 0, 0, 43,
		45, 5, 18, 0, 0, 44, 43, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 46, 1, 0,
		0, 0, 46, 47, 3, 20, 10, 0, 47, 48, 5, 7, 0, 0, 48, 49, 5, 1, 0, 0, 49,
		5, 1, 0, 0, 0, 50, 52, 7, 1, 0, 0, 51, 53, 5, 18, 0, 0, 52, 51, 1, 0, 0,
		0, 52, 53, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 3, 20, 10, 0, 55, 56,
		5, 7, 0, 0, 56, 57, 5, 2, 0, 0, 57, 58, 3, 20, 10, 0, 58, 59, 5, 3, 0,
		0, 59, 7, 1, 0, 0, 0, 60, 62, 7, 2, 0, 0, 61, 63, 5, 18, 0, 0, 62, 61,
		1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 3, 20, 10,
		0, 65, 66, 5, 7, 0, 0, 66, 67, 3, 20, 10, 0, 67, 68, 5, 7, 0, 0, 68, 69,
		3, 20, 10, 0, 69, 9, 1, 0, 0, 0, 70, 71, 5, 30, 0, 0, 71, 11, 1, 0, 0,
		0, 72, 74, 7, 3, 0, 0, 73, 75, 5, 18, 0, 0, 74, 73, 1, 0, 0, 0, 74, 75,
		1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 3, 20, 10, 0, 77, 78, 5, 7, 0,
		0, 78, 79, 3, 20, 10, 0, 79, 13, 1, 0, 0, 0, 80, 82, 5, 28, 0, 0, 81, 83,
		5, 18, 0, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0,
		84, 85, 3, 20, 10, 0, 85, 15, 1, 0, 0, 0, 86, 88, 5, 31, 0, 0, 87, 89,
		5, 18, 0, 0, 88, 87, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0,
		90, 91, 5, 2, 0, 0, 91, 92, 3, 20, 10, 0, 92, 93, 5, 3, 0, 0, 93, 17, 1,
		0, 0, 0, 94, 96, 5, 32, 0, 0, 95, 97, 5, 18, 0, 0, 96, 95, 1, 0, 0, 0,
		96, 97, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 7, 4, 0, 0, 99, 19, 1,
		0, 0, 0, 100, 101, 7, 5, 0, 0, 101, 21, 1, 0, 0, 0, 9, 25, 40, 44, 52,
		62, 74, 82, 88, 96,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// QuACParserInit initializes any static state used to implement QuACParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewQuACParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func QuACParserInit() {
	staticData := &quacparserParserStaticData
	staticData.once.Do(quacparserParserInit)
}

// NewQuACParser produces a new parser instance for the optional input antlr.TokenStream.
func NewQuACParser(input antlr.TokenStream) *QuACParser {
	QuACParserInit()
	this := new(QuACParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &quacparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "QuACParser.g4"

	return this
}

// QuACParser tokens.
const (
	QuACParserEOF            = antlr.TokenEOF
	QuACParserIntegerLiteral = 1
	QuACParserLBRACK         = 2
	QuACParserRBRACK         = 3
	QuACParserSEMI           = 4
	QuACParserCOLON          = 5
	QuACParserDOT            = 6
	QuACParserCOMMA          = 7
	QuACParserRZ             = 8
	QuACParserR0             = 9
	QuACParserR1             = 10
	QuACParserR2             = 11
	QuACParserR3             = 12
	QuACParserR4             = 13
	QuACParserFL             = 14
	QuACParserR5             = 15
	QuACParserPC             = 16
	QuACParserR7             = 17
	QuACParserEQ             = 18
	QuACParserMOV            = 19
	QuACParserMOVL           = 20
	QuACParserSETH           = 21
	QuACParserSTR            = 22
	QuACParserLDR            = 23
	QuACParserADD            = 24
	QuACParserSUB            = 25
	QuACParserAND            = 26
	QuACParserORR            = 27
	QuACParserJPR            = 28
	QuACParserCMP            = 29
	QuACParserNOP            = 30
	QuACParserJPM            = 31
	QuACParserJP             = 32
	QuACParserIdentifier     = 33
	QuACParserWS             = 34
	QuACParserLINE_COMMENT   = 35
)

// QuACParser rules.
const (
	QuACParserRULE_parse        = 0
	QuACParserRULE_statement    = 1
	QuACParserRULE_iFormat      = 2
	QuACParserRULE_rMemFormat   = 3
	QuACParserRULE_rALUFormat   = 4
	QuACParserRULE_nop          = 5
	QuACParserRULE_pseudo2Param = 6
	QuACParserRULE_jpr          = 7
	QuACParserRULE_jpm          = 8
	QuACParserRULE_jp           = 9
	QuACParserRULE_register     = 10
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuACParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuACParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(QuACParserEOF, 0)
}

func (s *ParseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *QuACParser) Parse() (localctx IParseContext) {
	this := p
	_ = this

	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QuACParserRULE_parse)
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(QuACParserMOV-19))|(1<<(QuACParserMOVL-19))|(1<<(QuACParserSETH-19))|(1<<(QuACParserSTR-19))|(1<<(QuACParserLDR-19))|(1<<(QuACParserADD-19))|(1<<(QuACParserSUB-19))|(1<<(QuACParserAND-19))|(1<<(QuACParserORR-19))|(1<<(QuACParserJPR-19))|(1<<(QuACParserCMP-19))|(1<<(QuACParserNOP-19))|(1<<(QuACParserJPM-19))|(1<<(QuACParserJP-19))|(1<<(QuACParserIdentifier-19)))) != 0 {
		{
			p.SetState(22)
			p.Statement()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(28)
		p.Match(QuACParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuACParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuACParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type JpmStatementContext struct {
	*StatementContext
}

func NewJpmStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JpmStatementContext {
	var p = new(JpmStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *JpmStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JpmStatementContext) Jpm() IJpmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJpmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJpmContext)
}

func (s *JpmStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterJpmStatement(s)
	}
}

func (s *JpmStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitJpmStatement(s)
	}
}

type RALUFormatStatementContext struct {
	*StatementContext
}

func NewRALUFormatStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RALUFormatStatementContext {
	var p = new(RALUFormatStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *RALUFormatStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RALUFormatStatementContext) RALUFormat() IRALUFormatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRALUFormatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRALUFormatContext)
}

func (s *RALUFormatStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterRALUFormatStatement(s)
	}
}

func (s *RALUFormatStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitRALUFormatStatement(s)
	}
}

type JpStatementContext struct {
	*StatementContext
}

func NewJpStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JpStatementContext {
	var p = new(JpStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *JpStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JpStatementContext) Jp() IJpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJpContext)
}

func (s *JpStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterJpStatement(s)
	}
}

func (s *JpStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitJpStatement(s)
	}
}

type Pseudo2ParamStatementContext struct {
	*StatementContext
}

func NewPseudo2ParamStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Pseudo2ParamStatementContext {
	var p = new(Pseudo2ParamStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *Pseudo2ParamStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pseudo2ParamStatementContext) Pseudo2Param() IPseudo2ParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPseudo2ParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPseudo2ParamContext)
}

func (s *Pseudo2ParamStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterPseudo2ParamStatement(s)
	}
}

func (s *Pseudo2ParamStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitPseudo2ParamStatement(s)
	}
}

type NopStatementContext struct {
	*StatementContext
}

func NewNopStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NopStatementContext {
	var p = new(NopStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *NopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NopStatementContext) Nop() INopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INopContext)
}

func (s *NopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterNopStatement(s)
	}
}

func (s *NopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitNopStatement(s)
	}
}

type JprStatementContext struct {
	*StatementContext
}

func NewJprStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JprStatementContext {
	var p = new(JprStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *JprStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JprStatementContext) Jpr() IJprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJprContext)
}

func (s *JprStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterJprStatement(s)
	}
}

func (s *JprStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitJprStatement(s)
	}
}

type RMemFormatStatementContext struct {
	*StatementContext
}

func NewRMemFormatStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RMemFormatStatementContext {
	var p = new(RMemFormatStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *RMemFormatStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RMemFormatStatementContext) RMemFormat() IRMemFormatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRMemFormatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRMemFormatContext)
}

func (s *RMemFormatStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterRMemFormatStatement(s)
	}
}

func (s *RMemFormatStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitRMemFormatStatement(s)
	}
}

type LabelStatementContext struct {
	*StatementContext
}

func NewLabelStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LabelStatementContext {
	var p = new(LabelStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *LabelStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(QuACParserIdentifier, 0)
}

func (s *LabelStatementContext) COLON() antlr.TerminalNode {
	return s.GetToken(QuACParserCOLON, 0)
}

func (s *LabelStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterLabelStatement(s)
	}
}

func (s *LabelStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitLabelStatement(s)
	}
}

type IFormatStatementContext struct {
	*StatementContext
}

func NewIFormatStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IFormatStatementContext {
	var p = new(IFormatStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IFormatStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IFormatStatementContext) IFormat() IIFormatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIFormatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIFormatContext)
}

func (s *IFormatStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterIFormatStatement(s)
	}
}

func (s *IFormatStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitIFormatStatement(s)
	}
}

func (p *QuACParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QuACParserRULE_statement)

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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QuACParserMOVL, QuACParserSETH:
		localctx = NewIFormatStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.IFormat()
		}

	case QuACParserSTR, QuACParserLDR:
		localctx = NewRMemFormatStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.RMemFormat()
		}

	case QuACParserADD, QuACParserSUB, QuACParserAND, QuACParserORR:
		localctx = NewRALUFormatStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)
			p.RALUFormat()
		}

	case QuACParserNOP:
		localctx = NewNopStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(33)
			p.Nop()
		}

	case QuACParserMOV, QuACParserCMP:
		localctx = NewPseudo2ParamStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(34)
			p.Pseudo2Param()
		}

	case QuACParserJPR:
		localctx = NewJprStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(35)
			p.Jpr()
		}

	case QuACParserJPM:
		localctx = NewJpmStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(36)
			p.Jpm()
		}

	case QuACParserJP:
		localctx = NewJpStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(37)
			p.Jp()
		}

	case QuACParserIdentifier:
		localctx = NewLabelStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(38)
			p.Match(QuACParserIdentifier)
		}
		{
			p.SetState(39)
			p.Match(QuACParserCOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIFormatContext is an interface to support dynamic dispatch.
type IIFormatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIFormatContext differentiates from other interfaces.
	IsIFormatContext()
}

type IFormatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIFormatContext() *IFormatContext {
	var p = new(IFormatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuACParserRULE_iFormat
	return p
}

func (*IFormatContext) IsIFormatContext() {}

func NewIFormatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IFormatContext {
	var p = new(IFormatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuACParserRULE_iFormat

	return p
}

func (s *IFormatContext) GetParser() antlr.Parser { return s.parser }

func (s *IFormatContext) Register() IRegisterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *IFormatContext) COMMA() antlr.TerminalNode {
	return s.GetToken(QuACParserCOMMA, 0)
}

func (s *IFormatContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(QuACParserIntegerLiteral, 0)
}

func (s *IFormatContext) MOVL() antlr.TerminalNode {
	return s.GetToken(QuACParserMOVL, 0)
}

func (s *IFormatContext) SETH() antlr.TerminalNode {
	return s.GetToken(QuACParserSETH, 0)
}

func (s *IFormatContext) EQ() antlr.TerminalNode {
	return s.GetToken(QuACParserEQ, 0)
}

func (s *IFormatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IFormatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IFormatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterIFormat(s)
	}
}

func (s *IFormatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitIFormat(s)
	}
}

func (p *QuACParser) IFormat() (localctx IIFormatContext) {
	this := p
	_ = this

	localctx = NewIFormatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QuACParserRULE_iFormat)
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
		p.SetState(42)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QuACParserMOVL || _la == QuACParserSETH) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuACParserEQ {
		{
			p.SetState(43)
			p.Match(QuACParserEQ)
		}

	}
	{
		p.SetState(46)
		p.Register()
	}
	{
		p.SetState(47)
		p.Match(QuACParserCOMMA)
	}
	{
		p.SetState(48)
		p.Match(QuACParserIntegerLiteral)
	}

	return localctx
}

// IRMemFormatContext is an interface to support dynamic dispatch.
type IRMemFormatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRMemFormatContext differentiates from other interfaces.
	IsRMemFormatContext()
}

type RMemFormatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRMemFormatContext() *RMemFormatContext {
	var p = new(RMemFormatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuACParserRULE_rMemFormat
	return p
}

func (*RMemFormatContext) IsRMemFormatContext() {}

func NewRMemFormatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RMemFormatContext {
	var p = new(RMemFormatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuACParserRULE_rMemFormat

	return p
}

func (s *RMemFormatContext) GetParser() antlr.Parser { return s.parser }

func (s *RMemFormatContext) AllRegister() []IRegisterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegisterContext); ok {
			len++
		}
	}

	tst := make([]IRegisterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegisterContext); ok {
			tst[i] = t.(IRegisterContext)
			i++
		}
	}

	return tst
}

func (s *RMemFormatContext) Register(i int) IRegisterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *RMemFormatContext) COMMA() antlr.TerminalNode {
	return s.GetToken(QuACParserCOMMA, 0)
}

func (s *RMemFormatContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(QuACParserLBRACK, 0)
}

func (s *RMemFormatContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(QuACParserRBRACK, 0)
}

func (s *RMemFormatContext) STR() antlr.TerminalNode {
	return s.GetToken(QuACParserSTR, 0)
}

func (s *RMemFormatContext) LDR() antlr.TerminalNode {
	return s.GetToken(QuACParserLDR, 0)
}

func (s *RMemFormatContext) EQ() antlr.TerminalNode {
	return s.GetToken(QuACParserEQ, 0)
}

func (s *RMemFormatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RMemFormatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RMemFormatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterRMemFormat(s)
	}
}

func (s *RMemFormatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitRMemFormat(s)
	}
}

func (p *QuACParser) RMemFormat() (localctx IRMemFormatContext) {
	this := p
	_ = this

	localctx = NewRMemFormatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QuACParserRULE_rMemFormat)
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
		p.SetState(50)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QuACParserSTR || _la == QuACParserLDR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuACParserEQ {
		{
			p.SetState(51)
			p.Match(QuACParserEQ)
		}

	}
	{
		p.SetState(54)
		p.Register()
	}
	{
		p.SetState(55)
		p.Match(QuACParserCOMMA)
	}
	{
		p.SetState(56)
		p.Match(QuACParserLBRACK)
	}
	{
		p.SetState(57)
		p.Register()
	}
	{
		p.SetState(58)
		p.Match(QuACParserRBRACK)
	}

	return localctx
}

// IRALUFormatContext is an interface to support dynamic dispatch.
type IRALUFormatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRALUFormatContext differentiates from other interfaces.
	IsRALUFormatContext()
}

type RALUFormatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRALUFormatContext() *RALUFormatContext {
	var p = new(RALUFormatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuACParserRULE_rALUFormat
	return p
}

func (*RALUFormatContext) IsRALUFormatContext() {}

func NewRALUFormatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RALUFormatContext {
	var p = new(RALUFormatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuACParserRULE_rALUFormat

	return p
}

func (s *RALUFormatContext) GetParser() antlr.Parser { return s.parser }

func (s *RALUFormatContext) AllRegister() []IRegisterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegisterContext); ok {
			len++
		}
	}

	tst := make([]IRegisterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegisterContext); ok {
			tst[i] = t.(IRegisterContext)
			i++
		}
	}

	return tst
}

func (s *RALUFormatContext) Register(i int) IRegisterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *RALUFormatContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(QuACParserCOMMA)
}

func (s *RALUFormatContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(QuACParserCOMMA, i)
}

func (s *RALUFormatContext) ADD() antlr.TerminalNode {
	return s.GetToken(QuACParserADD, 0)
}

func (s *RALUFormatContext) SUB() antlr.TerminalNode {
	return s.GetToken(QuACParserSUB, 0)
}

func (s *RALUFormatContext) AND() antlr.TerminalNode {
	return s.GetToken(QuACParserAND, 0)
}

func (s *RALUFormatContext) ORR() antlr.TerminalNode {
	return s.GetToken(QuACParserORR, 0)
}

func (s *RALUFormatContext) EQ() antlr.TerminalNode {
	return s.GetToken(QuACParserEQ, 0)
}

func (s *RALUFormatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RALUFormatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RALUFormatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterRALUFormat(s)
	}
}

func (s *RALUFormatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitRALUFormat(s)
	}
}

func (p *QuACParser) RALUFormat() (localctx IRALUFormatContext) {
	this := p
	_ = this

	localctx = NewRALUFormatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QuACParserRULE_rALUFormat)
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
		p.SetState(60)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuACParserADD)|(1<<QuACParserSUB)|(1<<QuACParserAND)|(1<<QuACParserORR))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuACParserEQ {
		{
			p.SetState(61)
			p.Match(QuACParserEQ)
		}

	}
	{
		p.SetState(64)
		p.Register()
	}
	{
		p.SetState(65)
		p.Match(QuACParserCOMMA)
	}
	{
		p.SetState(66)
		p.Register()
	}
	{
		p.SetState(67)
		p.Match(QuACParserCOMMA)
	}
	{
		p.SetState(68)
		p.Register()
	}

	return localctx
}

// INopContext is an interface to support dynamic dispatch.
type INopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNopContext differentiates from other interfaces.
	IsNopContext()
}

type NopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNopContext() *NopContext {
	var p = new(NopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuACParserRULE_nop
	return p
}

func (*NopContext) IsNopContext() {}

func NewNopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NopContext {
	var p = new(NopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuACParserRULE_nop

	return p
}

func (s *NopContext) GetParser() antlr.Parser { return s.parser }

func (s *NopContext) NOP() antlr.TerminalNode {
	return s.GetToken(QuACParserNOP, 0)
}

func (s *NopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterNop(s)
	}
}

func (s *NopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitNop(s)
	}
}

func (p *QuACParser) Nop() (localctx INopContext) {
	this := p
	_ = this

	localctx = NewNopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QuACParserRULE_nop)

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
		p.SetState(70)
		p.Match(QuACParserNOP)
	}

	return localctx
}

// IPseudo2ParamContext is an interface to support dynamic dispatch.
type IPseudo2ParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPseudo2ParamContext differentiates from other interfaces.
	IsPseudo2ParamContext()
}

type Pseudo2ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPseudo2ParamContext() *Pseudo2ParamContext {
	var p = new(Pseudo2ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuACParserRULE_pseudo2Param
	return p
}

func (*Pseudo2ParamContext) IsPseudo2ParamContext() {}

func NewPseudo2ParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pseudo2ParamContext {
	var p = new(Pseudo2ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuACParserRULE_pseudo2Param

	return p
}

func (s *Pseudo2ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *Pseudo2ParamContext) AllRegister() []IRegisterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegisterContext); ok {
			len++
		}
	}

	tst := make([]IRegisterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegisterContext); ok {
			tst[i] = t.(IRegisterContext)
			i++
		}
	}

	return tst
}

func (s *Pseudo2ParamContext) Register(i int) IRegisterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Pseudo2ParamContext) COMMA() antlr.TerminalNode {
	return s.GetToken(QuACParserCOMMA, 0)
}

func (s *Pseudo2ParamContext) MOV() antlr.TerminalNode {
	return s.GetToken(QuACParserMOV, 0)
}

func (s *Pseudo2ParamContext) CMP() antlr.TerminalNode {
	return s.GetToken(QuACParserCMP, 0)
}

func (s *Pseudo2ParamContext) EQ() antlr.TerminalNode {
	return s.GetToken(QuACParserEQ, 0)
}

func (s *Pseudo2ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pseudo2ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pseudo2ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterPseudo2Param(s)
	}
}

func (s *Pseudo2ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitPseudo2Param(s)
	}
}

func (p *QuACParser) Pseudo2Param() (localctx IPseudo2ParamContext) {
	this := p
	_ = this

	localctx = NewPseudo2ParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QuACParserRULE_pseudo2Param)
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
		p.SetState(72)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QuACParserMOV || _la == QuACParserCMP) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuACParserEQ {
		{
			p.SetState(73)
			p.Match(QuACParserEQ)
		}

	}
	{
		p.SetState(76)
		p.Register()
	}
	{
		p.SetState(77)
		p.Match(QuACParserCOMMA)
	}
	{
		p.SetState(78)
		p.Register()
	}

	return localctx
}

// IJprContext is an interface to support dynamic dispatch.
type IJprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJprContext differentiates from other interfaces.
	IsJprContext()
}

type JprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJprContext() *JprContext {
	var p = new(JprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuACParserRULE_jpr
	return p
}

func (*JprContext) IsJprContext() {}

func NewJprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JprContext {
	var p = new(JprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuACParserRULE_jpr

	return p
}

func (s *JprContext) GetParser() antlr.Parser { return s.parser }

func (s *JprContext) JPR() antlr.TerminalNode {
	return s.GetToken(QuACParserJPR, 0)
}

func (s *JprContext) Register() IRegisterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *JprContext) EQ() antlr.TerminalNode {
	return s.GetToken(QuACParserEQ, 0)
}

func (s *JprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterJpr(s)
	}
}

func (s *JprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitJpr(s)
	}
}

func (p *QuACParser) Jpr() (localctx IJprContext) {
	this := p
	_ = this

	localctx = NewJprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QuACParserRULE_jpr)
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
		p.SetState(80)
		p.Match(QuACParserJPR)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuACParserEQ {
		{
			p.SetState(81)
			p.Match(QuACParserEQ)
		}

	}
	{
		p.SetState(84)
		p.Register()
	}

	return localctx
}

// IJpmContext is an interface to support dynamic dispatch.
type IJpmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJpmContext differentiates from other interfaces.
	IsJpmContext()
}

type JpmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJpmContext() *JpmContext {
	var p = new(JpmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuACParserRULE_jpm
	return p
}

func (*JpmContext) IsJpmContext() {}

func NewJpmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JpmContext {
	var p = new(JpmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuACParserRULE_jpm

	return p
}

func (s *JpmContext) GetParser() antlr.Parser { return s.parser }

func (s *JpmContext) JPM() antlr.TerminalNode {
	return s.GetToken(QuACParserJPM, 0)
}

func (s *JpmContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(QuACParserLBRACK, 0)
}

func (s *JpmContext) Register() IRegisterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegisterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *JpmContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(QuACParserRBRACK, 0)
}

func (s *JpmContext) EQ() antlr.TerminalNode {
	return s.GetToken(QuACParserEQ, 0)
}

func (s *JpmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JpmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JpmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterJpm(s)
	}
}

func (s *JpmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitJpm(s)
	}
}

func (p *QuACParser) Jpm() (localctx IJpmContext) {
	this := p
	_ = this

	localctx = NewJpmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, QuACParserRULE_jpm)
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
		p.Match(QuACParserJPM)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuACParserEQ {
		{
			p.SetState(87)
			p.Match(QuACParserEQ)
		}

	}
	{
		p.SetState(90)
		p.Match(QuACParserLBRACK)
	}
	{
		p.SetState(91)
		p.Register()
	}
	{
		p.SetState(92)
		p.Match(QuACParserRBRACK)
	}

	return localctx
}

// IJpContext is an interface to support dynamic dispatch.
type IJpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJpContext differentiates from other interfaces.
	IsJpContext()
}

type JpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJpContext() *JpContext {
	var p = new(JpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuACParserRULE_jp
	return p
}

func (*JpContext) IsJpContext() {}

func NewJpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JpContext {
	var p = new(JpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuACParserRULE_jp

	return p
}

func (s *JpContext) GetParser() antlr.Parser { return s.parser }

func (s *JpContext) JP() antlr.TerminalNode {
	return s.GetToken(QuACParserJP, 0)
}

func (s *JpContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(QuACParserIntegerLiteral, 0)
}

func (s *JpContext) Identifier() antlr.TerminalNode {
	return s.GetToken(QuACParserIdentifier, 0)
}

func (s *JpContext) EQ() antlr.TerminalNode {
	return s.GetToken(QuACParserEQ, 0)
}

func (s *JpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterJp(s)
	}
}

func (s *JpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitJp(s)
	}
}

func (p *QuACParser) Jp() (localctx IJpContext) {
	this := p
	_ = this

	localctx = NewJpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, QuACParserRULE_jp)
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
		p.Match(QuACParserJP)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QuACParserEQ {
		{
			p.SetState(95)
			p.Match(QuACParserEQ)
		}

	}
	{
		p.SetState(98)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QuACParserIntegerLiteral || _la == QuACParserIdentifier) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRegisterContext is an interface to support dynamic dispatch.
type IRegisterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegisterContext differentiates from other interfaces.
	IsRegisterContext()
}

type RegisterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegisterContext() *RegisterContext {
	var p = new(RegisterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuACParserRULE_register
	return p
}

func (*RegisterContext) IsRegisterContext() {}

func NewRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterContext {
	var p = new(RegisterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuACParserRULE_register

	return p
}

func (s *RegisterContext) GetParser() antlr.Parser { return s.parser }

func (s *RegisterContext) RZ() antlr.TerminalNode {
	return s.GetToken(QuACParserRZ, 0)
}

func (s *RegisterContext) R0() antlr.TerminalNode {
	return s.GetToken(QuACParserR0, 0)
}

func (s *RegisterContext) R1() antlr.TerminalNode {
	return s.GetToken(QuACParserR1, 0)
}

func (s *RegisterContext) R2() antlr.TerminalNode {
	return s.GetToken(QuACParserR2, 0)
}

func (s *RegisterContext) R3() antlr.TerminalNode {
	return s.GetToken(QuACParserR3, 0)
}

func (s *RegisterContext) R4() antlr.TerminalNode {
	return s.GetToken(QuACParserR4, 0)
}

func (s *RegisterContext) FL() antlr.TerminalNode {
	return s.GetToken(QuACParserFL, 0)
}

func (s *RegisterContext) R5() antlr.TerminalNode {
	return s.GetToken(QuACParserR5, 0)
}

func (s *RegisterContext) PC() antlr.TerminalNode {
	return s.GetToken(QuACParserPC, 0)
}

func (s *RegisterContext) R7() antlr.TerminalNode {
	return s.GetToken(QuACParserR7, 0)
}

func (s *RegisterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegisterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegisterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.EnterRegister(s)
	}
}

func (s *RegisterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuACParserListener); ok {
		listenerT.ExitRegister(s)
	}
}

func (p *QuACParser) Register() (localctx IRegisterContext) {
	this := p
	_ = this

	localctx = NewRegisterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, QuACParserRULE_register)
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
		p.SetState(100)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QuACParserRZ)|(1<<QuACParserR0)|(1<<QuACParserR1)|(1<<QuACParserR2)|(1<<QuACParserR3)|(1<<QuACParserR4)|(1<<QuACParserFL)|(1<<QuACParserR5)|(1<<QuACParserPC)|(1<<QuACParserR7))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
