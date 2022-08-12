// Generated from e:\QuAC-Compiler\antlr4\QuACLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class QuACLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"IntegerLiteral", "DecimalIntegerLiteral", "HexIntegerLiteral", "OctalIntegerLiteral", 
			"BinaryIntegerLiteral", "IntegerTypeSuffix", "DecimalNumeral", "Digits", 
			"Digit", "NonZeroDigit", "DigitsAndUnderscores", "DigitOrUnderscore", 
			"Underscores", "HexNumeral", "HexDigits", "HexDigit", "HexDigitsAndUnderscores", 
			"HexDigitOrUnderscore", "OctalNumeral", "OctalDigits", "OctalDigit", 
			"OctalDigitsAndUnderscores", "OctalDigitOrUnderscore", "BinaryNumeral", 
			"BinaryDigits", "BinaryDigit", "BinaryDigitsAndUnderscores", "BinaryDigitOrUnderscore", 
			"LBRACK", "RBRACK", "SEMI", "COLON", "COMMA", "NEWLINE", "RZ", "R0", 
			"R1", "R2", "R3", "R4", "FL", "R5", "PC", "R7", "EQ", "MOV", "MOVL", 
			"SETH", "STR", "LDR", "ADD", "SUB", "AND", "ORR", "JPR", "CMP", "NOP", 
			"JPM", "JP", "WORD", "Identifier", "Letter", "LetterOrDigit", "WS", "LINE_COMMENT"
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


	public QuACLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "QuACLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2&\u018f\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\3\2\3\2\3\2\3\2\5\2\u008a\n\2\3\3\3\3\5"+
		"\3\u008e\n\3\3\4\3\4\5\4\u0092\n\4\3\5\3\5\5\5\u0096\n\5\3\6\3\6\5\6\u009a"+
		"\n\6\3\7\3\7\3\b\3\b\3\b\5\b\u00a1\n\b\3\b\3\b\3\b\5\b\u00a6\n\b\5\b\u00a8"+
		"\n\b\3\t\3\t\5\t\u00ac\n\t\3\t\5\t\u00af\n\t\3\n\3\n\5\n\u00b3\n\n\3\13"+
		"\3\13\3\f\6\f\u00b8\n\f\r\f\16\f\u00b9\3\r\3\r\5\r\u00be\n\r\3\16\6\16"+
		"\u00c1\n\16\r\16\16\16\u00c2\3\17\3\17\3\17\3\17\3\20\3\20\5\20\u00cb"+
		"\n\20\3\20\5\20\u00ce\n\20\3\21\3\21\3\22\6\22\u00d3\n\22\r\22\16\22\u00d4"+
		"\3\23\3\23\5\23\u00d9\n\23\3\24\3\24\5\24\u00dd\n\24\3\24\3\24\3\25\3"+
		"\25\5\25\u00e3\n\25\3\25\5\25\u00e6\n\25\3\26\3\26\3\27\6\27\u00eb\n\27"+
		"\r\27\16\27\u00ec\3\30\3\30\5\30\u00f1\n\30\3\31\3\31\3\31\3\31\3\32\3"+
		"\32\5\32\u00f9\n\32\3\32\5\32\u00fc\n\32\3\33\3\33\3\34\6\34\u0101\n\34"+
		"\r\34\16\34\u0102\3\35\3\35\5\35\u0107\n\35\3\36\3\36\3\37\3\37\3 \3 "+
		"\3!\3!\3\"\3\"\3#\3#\3$\3$\3$\3%\3%\3%\3&\3&\3&\3\'\3\'\3\'\3(\3(\3(\3"+
		")\3)\3)\3*\3*\3*\3+\3+\3+\3,\3,\3,\3-\3-\3-\3.\3.\3.\3/\3/\3/\3/\3\60"+
		"\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3\62\3\63"+
		"\3\63\3\63\3\63\3\64\3\64\3\64\3\64\3\65\3\65\3\65\3\65\3\66\3\66\3\66"+
		"\3\66\3\67\3\67\3\67\3\67\38\38\38\38\39\39\39\39\3:\3:\3:\3:\3;\3;\3"+
		";\3;\3<\3<\3<\3=\3=\3=\3=\3=\3=\3>\3>\7>\u0177\n>\f>\16>\u017a\13>\3?"+
		"\3?\3@\3@\3A\6A\u0181\nA\rA\16A\u0182\3A\3A\3B\3B\7B\u0189\nB\fB\16B\u018c"+
		"\13B\3B\3B\2\2C\3\3\5\2\7\2\t\2\13\2\r\2\17\2\21\2\23\2\25\2\27\2\31\2"+
		"\33\2\35\2\37\2!\2#\2%\2\'\2)\2+\2-\2/\2\61\2\63\2\65\2\67\29\2;\4=\5"+
		"?\6A\7C\bE\tG\nI\13K\fM\rO\16Q\17S\20U\21W\22Y\23[\24]\25_\26a\27c\30"+
		"e\31g\32i\33k\34m\35o\36q\37s u!w\"y#{$}\2\177\2\u0081%\u0083&\3\2\r\4"+
		"\2NNnn\3\2\63;\4\2ZZzz\5\2\62;CHch\3\2\629\4\2DDdd\3\2\62\63\6\2//C\\"+
		"aac|\7\2//\62;C\\aac|\5\2\13\f\16\17\"\"\4\2\f\f\17\17\2\u0191\2\3\3\2"+
		"\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2"+
		"\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S"+
		"\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2"+
		"\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2"+
		"\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y"+
		"\3\2\2\2\2{\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\3\u0089\3\2\2\2\5"+
		"\u008b\3\2\2\2\7\u008f\3\2\2\2\t\u0093\3\2\2\2\13\u0097\3\2\2\2\r\u009b"+
		"\3\2\2\2\17\u00a7\3\2\2\2\21\u00a9\3\2\2\2\23\u00b2\3\2\2\2\25\u00b4\3"+
		"\2\2\2\27\u00b7\3\2\2\2\31\u00bd\3\2\2\2\33\u00c0\3\2\2\2\35\u00c4\3\2"+
		"\2\2\37\u00c8\3\2\2\2!\u00cf\3\2\2\2#\u00d2\3\2\2\2%\u00d8\3\2\2\2\'\u00da"+
		"\3\2\2\2)\u00e0\3\2\2\2+\u00e7\3\2\2\2-\u00ea\3\2\2\2/\u00f0\3\2\2\2\61"+
		"\u00f2\3\2\2\2\63\u00f6\3\2\2\2\65\u00fd\3\2\2\2\67\u0100\3\2\2\29\u0106"+
		"\3\2\2\2;\u0108\3\2\2\2=\u010a\3\2\2\2?\u010c\3\2\2\2A\u010e\3\2\2\2C"+
		"\u0110\3\2\2\2E\u0112\3\2\2\2G\u0114\3\2\2\2I\u0117\3\2\2\2K\u011a\3\2"+
		"\2\2M\u011d\3\2\2\2O\u0120\3\2\2\2Q\u0123\3\2\2\2S\u0126\3\2\2\2U\u0129"+
		"\3\2\2\2W\u012c\3\2\2\2Y\u012f\3\2\2\2[\u0132\3\2\2\2]\u0135\3\2\2\2_"+
		"\u0139\3\2\2\2a\u013e\3\2\2\2c\u0143\3\2\2\2e\u0147\3\2\2\2g\u014b\3\2"+
		"\2\2i\u014f\3\2\2\2k\u0153\3\2\2\2m\u0157\3\2\2\2o\u015b\3\2\2\2q\u015f"+
		"\3\2\2\2s\u0163\3\2\2\2u\u0167\3\2\2\2w\u016b\3\2\2\2y\u016e\3\2\2\2{"+
		"\u0174\3\2\2\2}\u017b\3\2\2\2\177\u017d\3\2\2\2\u0081\u0180\3\2\2\2\u0083"+
		"\u0186\3\2\2\2\u0085\u008a\5\5\3\2\u0086\u008a\5\7\4\2\u0087\u008a\5\t"+
		"\5\2\u0088\u008a\5\13\6\2\u0089\u0085\3\2\2\2\u0089\u0086\3\2\2\2\u0089"+
		"\u0087\3\2\2\2\u0089\u0088\3\2\2\2\u008a\4\3\2\2\2\u008b\u008d\5\17\b"+
		"\2\u008c\u008e\5\r\7\2\u008d\u008c\3\2\2\2\u008d\u008e\3\2\2\2\u008e\6"+
		"\3\2\2\2\u008f\u0091\5\35\17\2\u0090\u0092\5\r\7\2\u0091\u0090\3\2\2\2"+
		"\u0091\u0092\3\2\2\2\u0092\b\3\2\2\2\u0093\u0095\5\'\24\2\u0094\u0096"+
		"\5\r\7\2\u0095\u0094\3\2\2\2\u0095\u0096\3\2\2\2\u0096\n\3\2\2\2\u0097"+
		"\u0099\5\61\31\2\u0098\u009a\5\r\7\2\u0099\u0098\3\2\2\2\u0099\u009a\3"+
		"\2\2\2\u009a\f\3\2\2\2\u009b\u009c\t\2\2\2\u009c\16\3\2\2\2\u009d\u00a8"+
		"\7\62\2\2\u009e\u00a5\5\25\13\2\u009f\u00a1\5\21\t\2\u00a0\u009f\3\2\2"+
		"\2\u00a0\u00a1\3\2\2\2\u00a1\u00a6\3\2\2\2\u00a2\u00a3\5\33\16\2\u00a3"+
		"\u00a4\5\21\t\2\u00a4\u00a6\3\2\2\2\u00a5\u00a0\3\2\2\2\u00a5\u00a2\3"+
		"\2\2\2\u00a6\u00a8\3\2\2\2\u00a7\u009d\3\2\2\2\u00a7\u009e\3\2\2\2\u00a8"+
		"\20\3\2\2\2\u00a9\u00ae\5\23\n\2\u00aa\u00ac\5\27\f\2\u00ab\u00aa\3\2"+
		"\2\2\u00ab\u00ac\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00af\5\23\n\2\u00ae"+
		"\u00ab\3\2\2\2\u00ae\u00af\3\2\2\2\u00af\22\3\2\2\2\u00b0\u00b3\7\62\2"+
		"\2\u00b1\u00b3\5\25\13\2\u00b2\u00b0\3\2\2\2\u00b2\u00b1\3\2\2\2\u00b3"+
		"\24\3\2\2\2\u00b4\u00b5\t\3\2\2\u00b5\26\3\2\2\2\u00b6\u00b8\5\31\r\2"+
		"\u00b7\u00b6\3\2\2\2\u00b8\u00b9\3\2\2\2\u00b9\u00b7\3\2\2\2\u00b9\u00ba"+
		"\3\2\2\2\u00ba\30\3\2\2\2\u00bb\u00be\5\23\n\2\u00bc\u00be\7a\2\2\u00bd"+
		"\u00bb\3\2\2\2\u00bd\u00bc\3\2\2\2\u00be\32\3\2\2\2\u00bf\u00c1\7a\2\2"+
		"\u00c0\u00bf\3\2\2\2\u00c1\u00c2\3\2\2\2\u00c2\u00c0\3\2\2\2\u00c2\u00c3"+
		"\3\2\2\2\u00c3\34\3\2\2\2\u00c4\u00c5\7\62\2\2\u00c5\u00c6\t\4\2\2\u00c6"+
		"\u00c7\5\37\20\2\u00c7\36\3\2\2\2\u00c8\u00cd\5!\21\2\u00c9\u00cb\5#\22"+
		"\2\u00ca\u00c9\3\2\2\2\u00ca\u00cb\3\2\2\2\u00cb\u00cc\3\2\2\2\u00cc\u00ce"+
		"\5!\21\2\u00cd\u00ca\3\2\2\2\u00cd\u00ce\3\2\2\2\u00ce \3\2\2\2\u00cf"+
		"\u00d0\t\5\2\2\u00d0\"\3\2\2\2\u00d1\u00d3\5%\23\2\u00d2\u00d1\3\2\2\2"+
		"\u00d3\u00d4\3\2\2\2\u00d4\u00d2\3\2\2\2\u00d4\u00d5\3\2\2\2\u00d5$\3"+
		"\2\2\2\u00d6\u00d9\5!\21\2\u00d7\u00d9\7a\2\2\u00d8\u00d6\3\2\2\2\u00d8"+
		"\u00d7\3\2\2\2\u00d9&\3\2\2\2\u00da\u00dc\7\62\2\2\u00db\u00dd\5\33\16"+
		"\2\u00dc\u00db\3\2\2\2\u00dc\u00dd\3\2\2\2\u00dd\u00de\3\2\2\2\u00de\u00df"+
		"\5)\25\2\u00df(\3\2\2\2\u00e0\u00e5\5+\26\2\u00e1\u00e3\5-\27\2\u00e2"+
		"\u00e1\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3\u00e4\3\2\2\2\u00e4\u00e6\5+"+
		"\26\2\u00e5\u00e2\3\2\2\2\u00e5\u00e6\3\2\2\2\u00e6*\3\2\2\2\u00e7\u00e8"+
		"\t\6\2\2\u00e8,\3\2\2\2\u00e9\u00eb\5/\30\2\u00ea\u00e9\3\2\2\2\u00eb"+
		"\u00ec\3\2\2\2\u00ec\u00ea\3\2\2\2\u00ec\u00ed\3\2\2\2\u00ed.\3\2\2\2"+
		"\u00ee\u00f1\5+\26\2\u00ef\u00f1\7a\2\2\u00f0\u00ee\3\2\2\2\u00f0\u00ef"+
		"\3\2\2\2\u00f1\60\3\2\2\2\u00f2\u00f3\7\62\2\2\u00f3\u00f4\t\7\2\2\u00f4"+
		"\u00f5\5\63\32\2\u00f5\62\3\2\2\2\u00f6\u00fb\5\65\33\2\u00f7\u00f9\5"+
		"\67\34\2\u00f8\u00f7\3\2\2\2\u00f8\u00f9\3\2\2\2\u00f9\u00fa\3\2\2\2\u00fa"+
		"\u00fc\5\65\33\2\u00fb\u00f8\3\2\2\2\u00fb\u00fc\3\2\2\2\u00fc\64\3\2"+
		"\2\2\u00fd\u00fe\t\b\2\2\u00fe\66\3\2\2\2\u00ff\u0101\59\35\2\u0100\u00ff"+
		"\3\2\2\2\u0101\u0102\3\2\2\2\u0102\u0100\3\2\2\2\u0102\u0103\3\2\2\2\u0103"+
		"8\3\2\2\2\u0104\u0107\5\65\33\2\u0105\u0107\7a\2\2\u0106\u0104\3\2\2\2"+
		"\u0106\u0105\3\2\2\2\u0107:\3\2\2\2\u0108\u0109\7]\2\2\u0109<\3\2\2\2"+
		"\u010a\u010b\7_\2\2\u010b>\3\2\2\2\u010c\u010d\7=\2\2\u010d@\3\2\2\2\u010e"+
		"\u010f\7<\2\2\u010fB\3\2\2\2\u0110\u0111\7.\2\2\u0111D\3\2\2\2\u0112\u0113"+
		"\7\f\2\2\u0113F\3\2\2\2\u0114\u0115\7t\2\2\u0115\u0116\7|\2\2\u0116H\3"+
		"\2\2\2\u0117\u0118\7t\2\2\u0118\u0119\7\62\2\2\u0119J\3\2\2\2\u011a\u011b"+
		"\7t\2\2\u011b\u011c\7\63\2\2\u011cL\3\2\2\2\u011d\u011e\7t\2\2\u011e\u011f"+
		"\7\64\2\2\u011fN\3\2\2\2\u0120\u0121\7t\2\2\u0121\u0122\7\65\2\2\u0122"+
		"P\3\2\2\2\u0123\u0124\7t\2\2\u0124\u0125\7\66\2\2\u0125R\3\2\2\2\u0126"+
		"\u0127\7h\2\2\u0127\u0128\7n\2\2\u0128T\3\2\2\2\u0129\u012a\7t\2\2\u012a"+
		"\u012b\7\67\2\2\u012bV\3\2\2\2\u012c\u012d\7r\2\2\u012d\u012e\7e\2\2\u012e"+
		"X\3\2\2\2\u012f\u0130\7t\2\2\u0130\u0131\79\2\2\u0131Z\3\2\2\2\u0132\u0133"+
		"\7g\2\2\u0133\u0134\7s\2\2\u0134\\\3\2\2\2\u0135\u0136\7o\2\2\u0136\u0137"+
		"\7q\2\2\u0137\u0138\7x\2\2\u0138^\3\2\2\2\u0139\u013a\7o\2\2\u013a\u013b"+
		"\7q\2\2\u013b\u013c\7x\2\2\u013c\u013d\7n\2\2\u013d`\3\2\2\2\u013e\u013f"+
		"\7u\2\2\u013f\u0140\7g\2\2\u0140\u0141\7v\2\2\u0141\u0142\7j\2\2\u0142"+
		"b\3\2\2\2\u0143\u0144\7u\2\2\u0144\u0145\7v\2\2\u0145\u0146\7t\2\2\u0146"+
		"d\3\2\2\2\u0147\u0148\7n\2\2\u0148\u0149\7f\2\2\u0149\u014a\7t\2\2\u014a"+
		"f\3\2\2\2\u014b\u014c\7c\2\2\u014c\u014d\7f\2\2\u014d\u014e\7f\2\2\u014e"+
		"h\3\2\2\2\u014f\u0150\7u\2\2\u0150\u0151\7w\2\2\u0151\u0152\7d\2\2\u0152"+
		"j\3\2\2\2\u0153\u0154\7c\2\2\u0154\u0155\7p\2\2\u0155\u0156\7f\2\2\u0156"+
		"l\3\2\2\2\u0157\u0158\7q\2\2\u0158\u0159\7t\2\2\u0159\u015a\7t\2\2\u015a"+
		"n\3\2\2\2\u015b\u015c\7l\2\2\u015c\u015d\7r\2\2\u015d\u015e\7t\2\2\u015e"+
		"p\3\2\2\2\u015f\u0160\7e\2\2\u0160\u0161\7o\2\2\u0161\u0162\7r\2\2\u0162"+
		"r\3\2\2\2\u0163\u0164\7p\2\2\u0164\u0165\7q\2\2\u0165\u0166\7r\2\2\u0166"+
		"t\3\2\2\2\u0167\u0168\7l\2\2\u0168\u0169\7r\2\2\u0169\u016a\7o\2\2\u016a"+
		"v\3\2\2\2\u016b\u016c\7l\2\2\u016c\u016d\7r\2\2\u016dx\3\2\2\2\u016e\u016f"+
		"\7\60\2\2\u016f\u0170\7y\2\2\u0170\u0171\7q\2\2\u0171\u0172\7t\2\2\u0172"+
		"\u0173\7f\2\2\u0173z\3\2\2\2\u0174\u0178\5}?\2\u0175\u0177\5\177@\2\u0176"+
		"\u0175\3\2\2\2\u0177\u017a\3\2\2\2\u0178\u0176\3\2\2\2\u0178\u0179\3\2"+
		"\2\2\u0179|\3\2\2\2\u017a\u0178\3\2\2\2\u017b\u017c\t\t\2\2\u017c~\3\2"+
		"\2\2\u017d\u017e\t\n\2\2\u017e\u0080\3\2\2\2\u017f\u0181\t\13\2\2\u0180"+
		"\u017f\3\2\2\2\u0181\u0182\3\2\2\2\u0182\u0180\3\2\2\2\u0182\u0183\3\2"+
		"\2\2\u0183\u0184\3\2\2\2\u0184\u0185\bA\2\2\u0185\u0082\3\2\2\2\u0186"+
		"\u018a\7=\2\2\u0187\u0189\n\f\2\2\u0188\u0187\3\2\2\2\u0189\u018c\3\2"+
		"\2\2\u018a\u0188\3\2\2\2\u018a\u018b\3\2\2\2\u018b\u018d\3\2\2\2\u018c"+
		"\u018a\3\2\2\2\u018d\u018e\bB\3\2\u018e\u0084\3\2\2\2!\2\u0089\u008d\u0091"+
		"\u0095\u0099\u00a0\u00a5\u00a7\u00ab\u00ae\u00b2\u00b9\u00bd\u00c2\u00ca"+
		"\u00cd\u00d4\u00d8\u00dc\u00e2\u00e5\u00ec\u00f0\u00f8\u00fb\u0102\u0106"+
		"\u0178\u0182\u018a\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}