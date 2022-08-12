// Generated from e:\QuAC-Compiler\antlr4\QuACParser.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class QuACParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		IntegerLiteral=1, LBRACK=2, RBRACK=3, SEMI=4, COLON=5, COMMA=6, NEWLINE=7, 
		RZ=8, R0=9, R1=10, R2=11, R3=12, R4=13, FL=14, R5=15, PC=16, R7=17, EQ=18, 
		MOV=19, MOVL=20, SETH=21, STR=22, LDR=23, ADD=24, SUB=25, AND=26, ORR=27, 
		JPR=28, CMP=29, NOP=30, JPM=31, JP=32, WORD=33, Identifier=34, WS=35, 
		LINE_COMMENT=36;
	public static final int
		RULE_parse = 0, RULE_statement = 1, RULE_iFormat = 2, RULE_rMemFormat = 3, 
		RULE_rALUFormat = 4, RULE_nop = 5, RULE_pseudo2Param = 6, RULE_jpr = 7, 
		RULE_jpm = 8, RULE_jp = 9, RULE_register = 10;
	private static String[] makeRuleNames() {
		return new String[] {
			"parse", "statement", "iFormat", "rMemFormat", "rALUFormat", "nop", "pseudo2Param", 
			"jpr", "jpm", "jp", "register"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'['", "']'", "';'", "':'", "','", "'\n'", "'rz'", "'r0'", 
			"'r1'", "'r2'", "'r3'", "'r4'", "'fl'", "'r5'", "'pc'", "'r7'", "'eq'", 
			"'mov'", "'movl'", "'seth'", "'str'", "'ldr'", "'add'", "'sub'", "'and'", 
			"'orr'", "'jpr'", "'cmp'", "'nop'", "'jpm'", "'jp'", "'.word'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "IntegerLiteral", "LBRACK", "RBRACK", "SEMI", "COLON", "COMMA", 
			"NEWLINE", "RZ", "R0", "R1", "R2", "R3", "R4", "FL", "R5", "PC", "R7", 
			"EQ", "MOV", "MOVL", "SETH", "STR", "LDR", "ADD", "SUB", "AND", "ORR", 
			"JPR", "CMP", "NOP", "JPM", "JP", "WORD", "Identifier", "WS", "LINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "QuACParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public QuACParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ParseContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(QuACParser.EOF, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public ParseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parse; }
	}

	public final ParseContext parse() throws RecognitionException {
		ParseContext _localctx = new ParseContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_parse);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MOV) | (1L << MOVL) | (1L << SETH) | (1L << STR) | (1L << LDR) | (1L << ADD) | (1L << SUB) | (1L << AND) | (1L << ORR) | (1L << JPR) | (1L << CMP) | (1L << NOP) | (1L << JPM) | (1L << JP) | (1L << WORD) | (1L << Identifier))) != 0)) {
				{
				{
				setState(22);
				statement();
				}
				}
				setState(27);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(28);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	 
		public StatementContext() { }
		public void copyFrom(StatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class JpmStatementContext extends StatementContext {
		public JpmContext jpm() {
			return getRuleContext(JpmContext.class,0);
		}
		public JpmStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class RALUFormatStatementContext extends StatementContext {
		public RALUFormatContext rALUFormat() {
			return getRuleContext(RALUFormatContext.class,0);
		}
		public RALUFormatStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class JpStatementContext extends StatementContext {
		public JpContext jp() {
			return getRuleContext(JpContext.class,0);
		}
		public JpStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class Pseudo2ParamStatementContext extends StatementContext {
		public Pseudo2ParamContext pseudo2Param() {
			return getRuleContext(Pseudo2ParamContext.class,0);
		}
		public Pseudo2ParamStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class NopStatementContext extends StatementContext {
		public NopContext nop() {
			return getRuleContext(NopContext.class,0);
		}
		public NopStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class JprStatementContext extends StatementContext {
		public JprContext jpr() {
			return getRuleContext(JprContext.class,0);
		}
		public JprStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class WordStatementContext extends StatementContext {
		public TerminalNode WORD() { return getToken(QuACParser.WORD, 0); }
		public TerminalNode IntegerLiteral() { return getToken(QuACParser.IntegerLiteral, 0); }
		public WordStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class RMemFormatStatementContext extends StatementContext {
		public RMemFormatContext rMemFormat() {
			return getRuleContext(RMemFormatContext.class,0);
		}
		public RMemFormatStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class LabelStatementContext extends StatementContext {
		public TerminalNode Identifier() { return getToken(QuACParser.Identifier, 0); }
		public TerminalNode COLON() { return getToken(QuACParser.COLON, 0); }
		public LabelStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class IFormatStatementContext extends StatementContext {
		public IFormatContext iFormat() {
			return getRuleContext(IFormatContext.class,0);
		}
		public IFormatStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(42);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MOVL:
			case SETH:
				_localctx = new IFormatStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(30);
				iFormat();
				}
				break;
			case STR:
			case LDR:
				_localctx = new RMemFormatStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(31);
				rMemFormat();
				}
				break;
			case ADD:
			case SUB:
			case AND:
			case ORR:
				_localctx = new RALUFormatStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(32);
				rALUFormat();
				}
				break;
			case NOP:
				_localctx = new NopStatementContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(33);
				nop();
				}
				break;
			case MOV:
			case CMP:
				_localctx = new Pseudo2ParamStatementContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(34);
				pseudo2Param();
				}
				break;
			case JPR:
				_localctx = new JprStatementContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(35);
				jpr();
				}
				break;
			case JPM:
				_localctx = new JpmStatementContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(36);
				jpm();
				}
				break;
			case JP:
				_localctx = new JpStatementContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(37);
				jp();
				}
				break;
			case WORD:
				_localctx = new WordStatementContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(38);
				match(WORD);
				setState(39);
				match(IntegerLiteral);
				}
				break;
			case Identifier:
				_localctx = new LabelStatementContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(40);
				match(Identifier);
				setState(41);
				match(COLON);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IFormatContext extends ParserRuleContext {
		public RegisterContext register() {
			return getRuleContext(RegisterContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(QuACParser.COMMA, 0); }
		public TerminalNode IntegerLiteral() { return getToken(QuACParser.IntegerLiteral, 0); }
		public TerminalNode MOVL() { return getToken(QuACParser.MOVL, 0); }
		public TerminalNode SETH() { return getToken(QuACParser.SETH, 0); }
		public TerminalNode EQ() { return getToken(QuACParser.EQ, 0); }
		public IFormatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_iFormat; }
	}

	public final IFormatContext iFormat() throws RecognitionException {
		IFormatContext _localctx = new IFormatContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_iFormat);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			_la = _input.LA(1);
			if ( !(_la==MOVL || _la==SETH) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(46);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQ) {
				{
				setState(45);
				match(EQ);
				}
			}

			setState(48);
			register();
			setState(49);
			match(COMMA);
			setState(50);
			match(IntegerLiteral);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RMemFormatContext extends ParserRuleContext {
		public List<RegisterContext> register() {
			return getRuleContexts(RegisterContext.class);
		}
		public RegisterContext register(int i) {
			return getRuleContext(RegisterContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(QuACParser.COMMA, 0); }
		public TerminalNode LBRACK() { return getToken(QuACParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(QuACParser.RBRACK, 0); }
		public TerminalNode STR() { return getToken(QuACParser.STR, 0); }
		public TerminalNode LDR() { return getToken(QuACParser.LDR, 0); }
		public TerminalNode EQ() { return getToken(QuACParser.EQ, 0); }
		public RMemFormatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rMemFormat; }
	}

	public final RMemFormatContext rMemFormat() throws RecognitionException {
		RMemFormatContext _localctx = new RMemFormatContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_rMemFormat);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			_la = _input.LA(1);
			if ( !(_la==STR || _la==LDR) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(54);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQ) {
				{
				setState(53);
				match(EQ);
				}
			}

			setState(56);
			register();
			setState(57);
			match(COMMA);
			setState(58);
			match(LBRACK);
			setState(59);
			register();
			setState(60);
			match(RBRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RALUFormatContext extends ParserRuleContext {
		public List<RegisterContext> register() {
			return getRuleContexts(RegisterContext.class);
		}
		public RegisterContext register(int i) {
			return getRuleContext(RegisterContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(QuACParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(QuACParser.COMMA, i);
		}
		public TerminalNode ADD() { return getToken(QuACParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(QuACParser.SUB, 0); }
		public TerminalNode AND() { return getToken(QuACParser.AND, 0); }
		public TerminalNode ORR() { return getToken(QuACParser.ORR, 0); }
		public TerminalNode EQ() { return getToken(QuACParser.EQ, 0); }
		public RALUFormatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rALUFormat; }
	}

	public final RALUFormatContext rALUFormat() throws RecognitionException {
		RALUFormatContext _localctx = new RALUFormatContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_rALUFormat);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ADD) | (1L << SUB) | (1L << AND) | (1L << ORR))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(64);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQ) {
				{
				setState(63);
				match(EQ);
				}
			}

			setState(66);
			register();
			setState(67);
			match(COMMA);
			setState(68);
			register();
			setState(69);
			match(COMMA);
			setState(70);
			register();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NopContext extends ParserRuleContext {
		public TerminalNode NOP() { return getToken(QuACParser.NOP, 0); }
		public NopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nop; }
	}

	public final NopContext nop() throws RecognitionException {
		NopContext _localctx = new NopContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_nop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			match(NOP);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Pseudo2ParamContext extends ParserRuleContext {
		public List<RegisterContext> register() {
			return getRuleContexts(RegisterContext.class);
		}
		public RegisterContext register(int i) {
			return getRuleContext(RegisterContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(QuACParser.COMMA, 0); }
		public TerminalNode MOV() { return getToken(QuACParser.MOV, 0); }
		public TerminalNode CMP() { return getToken(QuACParser.CMP, 0); }
		public TerminalNode EQ() { return getToken(QuACParser.EQ, 0); }
		public Pseudo2ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pseudo2Param; }
	}

	public final Pseudo2ParamContext pseudo2Param() throws RecognitionException {
		Pseudo2ParamContext _localctx = new Pseudo2ParamContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_pseudo2Param);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(74);
			_la = _input.LA(1);
			if ( !(_la==MOV || _la==CMP) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(76);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQ) {
				{
				setState(75);
				match(EQ);
				}
			}

			setState(78);
			register();
			setState(79);
			match(COMMA);
			setState(80);
			register();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class JprContext extends ParserRuleContext {
		public TerminalNode JPR() { return getToken(QuACParser.JPR, 0); }
		public RegisterContext register() {
			return getRuleContext(RegisterContext.class,0);
		}
		public TerminalNode EQ() { return getToken(QuACParser.EQ, 0); }
		public JprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_jpr; }
	}

	public final JprContext jpr() throws RecognitionException {
		JprContext _localctx = new JprContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_jpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(82);
			match(JPR);
			setState(84);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQ) {
				{
				setState(83);
				match(EQ);
				}
			}

			setState(86);
			register();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class JpmContext extends ParserRuleContext {
		public TerminalNode JPM() { return getToken(QuACParser.JPM, 0); }
		public TerminalNode LBRACK() { return getToken(QuACParser.LBRACK, 0); }
		public RegisterContext register() {
			return getRuleContext(RegisterContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(QuACParser.RBRACK, 0); }
		public TerminalNode EQ() { return getToken(QuACParser.EQ, 0); }
		public JpmContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_jpm; }
	}

	public final JpmContext jpm() throws RecognitionException {
		JpmContext _localctx = new JpmContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_jpm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(88);
			match(JPM);
			setState(90);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQ) {
				{
				setState(89);
				match(EQ);
				}
			}

			setState(92);
			match(LBRACK);
			setState(93);
			register();
			setState(94);
			match(RBRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class JpContext extends ParserRuleContext {
		public TerminalNode JP() { return getToken(QuACParser.JP, 0); }
		public TerminalNode IntegerLiteral() { return getToken(QuACParser.IntegerLiteral, 0); }
		public TerminalNode Identifier() { return getToken(QuACParser.Identifier, 0); }
		public TerminalNode EQ() { return getToken(QuACParser.EQ, 0); }
		public JpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_jp; }
	}

	public final JpContext jp() throws RecognitionException {
		JpContext _localctx = new JpContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_jp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(96);
			match(JP);
			setState(98);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQ) {
				{
				setState(97);
				match(EQ);
				}
			}

			setState(100);
			_la = _input.LA(1);
			if ( !(_la==IntegerLiteral || _la==Identifier) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RegisterContext extends ParserRuleContext {
		public TerminalNode RZ() { return getToken(QuACParser.RZ, 0); }
		public TerminalNode R0() { return getToken(QuACParser.R0, 0); }
		public TerminalNode R1() { return getToken(QuACParser.R1, 0); }
		public TerminalNode R2() { return getToken(QuACParser.R2, 0); }
		public TerminalNode R3() { return getToken(QuACParser.R3, 0); }
		public TerminalNode R4() { return getToken(QuACParser.R4, 0); }
		public TerminalNode FL() { return getToken(QuACParser.FL, 0); }
		public TerminalNode R5() { return getToken(QuACParser.R5, 0); }
		public TerminalNode PC() { return getToken(QuACParser.PC, 0); }
		public TerminalNode R7() { return getToken(QuACParser.R7, 0); }
		public RegisterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_register; }
	}

	public final RegisterContext register() throws RecognitionException {
		RegisterContext _localctx = new RegisterContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_register);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RZ) | (1L << R0) | (1L << R1) | (1L << R2) | (1L << R3) | (1L << R4) | (1L << FL) | (1L << R5) | (1L << PC) | (1L << R7))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3&k\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4\f"+
		"\t\f\3\2\7\2\32\n\2\f\2\16\2\35\13\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\5\3-\n\3\3\4\3\4\5\4\61\n\4\3\4\3\4\3\4\3\4\3\5"+
		"\3\5\5\59\n\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\5\6C\n\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\7\3\7\3\b\3\b\5\bO\n\b\3\b\3\b\3\b\3\b\3\t\3\t\5\tW\n\t\3\t"+
		"\3\t\3\n\3\n\5\n]\n\n\3\n\3\n\3\n\3\n\3\13\3\13\5\13e\n\13\3\13\3\13\3"+
		"\f\3\f\3\f\2\2\r\2\4\6\b\n\f\16\20\22\24\26\2\b\3\2\26\27\3\2\30\31\3"+
		"\2\32\35\4\2\25\25\37\37\4\2\3\3$$\3\2\n\23\2p\2\33\3\2\2\2\4,\3\2\2\2"+
		"\6.\3\2\2\2\b\66\3\2\2\2\n@\3\2\2\2\fJ\3\2\2\2\16L\3\2\2\2\20T\3\2\2\2"+
		"\22Z\3\2\2\2\24b\3\2\2\2\26h\3\2\2\2\30\32\5\4\3\2\31\30\3\2\2\2\32\35"+
		"\3\2\2\2\33\31\3\2\2\2\33\34\3\2\2\2\34\36\3\2\2\2\35\33\3\2\2\2\36\37"+
		"\7\2\2\3\37\3\3\2\2\2 -\5\6\4\2!-\5\b\5\2\"-\5\n\6\2#-\5\f\7\2$-\5\16"+
		"\b\2%-\5\20\t\2&-\5\22\n\2\'-\5\24\13\2()\7#\2\2)-\7\3\2\2*+\7$\2\2+-"+
		"\7\7\2\2, \3\2\2\2,!\3\2\2\2,\"\3\2\2\2,#\3\2\2\2,$\3\2\2\2,%\3\2\2\2"+
		",&\3\2\2\2,\'\3\2\2\2,(\3\2\2\2,*\3\2\2\2-\5\3\2\2\2.\60\t\2\2\2/\61\7"+
		"\24\2\2\60/\3\2\2\2\60\61\3\2\2\2\61\62\3\2\2\2\62\63\5\26\f\2\63\64\7"+
		"\b\2\2\64\65\7\3\2\2\65\7\3\2\2\2\668\t\3\2\2\679\7\24\2\28\67\3\2\2\2"+
		"89\3\2\2\29:\3\2\2\2:;\5\26\f\2;<\7\b\2\2<=\7\4\2\2=>\5\26\f\2>?\7\5\2"+
		"\2?\t\3\2\2\2@B\t\4\2\2AC\7\24\2\2BA\3\2\2\2BC\3\2\2\2CD\3\2\2\2DE\5\26"+
		"\f\2EF\7\b\2\2FG\5\26\f\2GH\7\b\2\2HI\5\26\f\2I\13\3\2\2\2JK\7 \2\2K\r"+
		"\3\2\2\2LN\t\5\2\2MO\7\24\2\2NM\3\2\2\2NO\3\2\2\2OP\3\2\2\2PQ\5\26\f\2"+
		"QR\7\b\2\2RS\5\26\f\2S\17\3\2\2\2TV\7\36\2\2UW\7\24\2\2VU\3\2\2\2VW\3"+
		"\2\2\2WX\3\2\2\2XY\5\26\f\2Y\21\3\2\2\2Z\\\7!\2\2[]\7\24\2\2\\[\3\2\2"+
		"\2\\]\3\2\2\2]^\3\2\2\2^_\7\4\2\2_`\5\26\f\2`a\7\5\2\2a\23\3\2\2\2bd\7"+
		"\"\2\2ce\7\24\2\2dc\3\2\2\2de\3\2\2\2ef\3\2\2\2fg\t\6\2\2g\25\3\2\2\2"+
		"hi\t\7\2\2i\27\3\2\2\2\13\33,\608BNV\\d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}