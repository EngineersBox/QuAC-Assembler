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
		IntegerLiteral=1, LBRACK=2, RBRACK=3, SEMI=4, COLON=5, COMMA=6, RZ=7, 
		R0=8, R1=9, R2=10, R3=11, R4=12, FL=13, R5=14, PC=15, R7=16, EQ=17, MOV=18, 
		MOVL=19, SETH=20, STR=21, LDR=22, ADD=23, SUB=24, AND=25, ORR=26, JPR=27, 
		CMP=28, NOP=29, JPM=30, JP=31, WORD=32, Identifier=33, WS=34, LINE_COMMENT=35;
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
			"LBRACK", "RBRACK", "SEMI", "COLON", "COMMA", "RZ", "R0", "R1", "R2", 
			"R3", "R4", "FL", "R5", "PC", "R7", "EQ", "MOV", "MOVL", "SETH", "STR", 
			"LDR", "ADD", "SUB", "AND", "ORR", "JPR", "CMP", "NOP", "JPM", "JP", 
			"WORD", "Identifier", "Letter", "LetterOrDigit", "WS", "LINE_COMMENT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'['", "']'", "';'", "':'", "','", "'rz'", "'r0'", "'r1'", 
			"'r2'", "'r3'", "'r4'", "'fl'", "'r5'", "'pc'", "'r7'", "'eq'", "'mov'", 
			"'movl'", "'seth'", "'str'", "'ldr'", "'add'", "'sub'", "'and'", "'orr'", 
			"'jpr'", "'cmp'", "'nop'", "'jpm'", "'jp'", "'.word'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "IntegerLiteral", "LBRACK", "RBRACK", "SEMI", "COLON", "COMMA", 
			"RZ", "R0", "R1", "R2", "R3", "R4", "FL", "R5", "PC", "R7", "EQ", "MOV", 
			"MOVL", "SETH", "STR", "LDR", "ADD", "SUB", "AND", "ORR", "JPR", "CMP", 
			"NOP", "JPM", "JP", "WORD", "Identifier", "WS", "LINE_COMMENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2%\u018b\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\3\2\3\2\3\2\3\2\5\2\u0088\n\2\3\3\3\3\5\3\u008c"+
		"\n\3\3\4\3\4\5\4\u0090\n\4\3\5\3\5\5\5\u0094\n\5\3\6\3\6\5\6\u0098\n\6"+
		"\3\7\3\7\3\b\3\b\3\b\5\b\u009f\n\b\3\b\3\b\3\b\5\b\u00a4\n\b\5\b\u00a6"+
		"\n\b\3\t\3\t\5\t\u00aa\n\t\3\t\5\t\u00ad\n\t\3\n\3\n\5\n\u00b1\n\n\3\13"+
		"\3\13\3\f\6\f\u00b6\n\f\r\f\16\f\u00b7\3\r\3\r\5\r\u00bc\n\r\3\16\6\16"+
		"\u00bf\n\16\r\16\16\16\u00c0\3\17\3\17\3\17\3\17\3\20\3\20\5\20\u00c9"+
		"\n\20\3\20\5\20\u00cc\n\20\3\21\3\21\3\22\6\22\u00d1\n\22\r\22\16\22\u00d2"+
		"\3\23\3\23\5\23\u00d7\n\23\3\24\3\24\5\24\u00db\n\24\3\24\3\24\3\25\3"+
		"\25\5\25\u00e1\n\25\3\25\5\25\u00e4\n\25\3\26\3\26\3\27\6\27\u00e9\n\27"+
		"\r\27\16\27\u00ea\3\30\3\30\5\30\u00ef\n\30\3\31\3\31\3\31\3\31\3\32\3"+
		"\32\5\32\u00f7\n\32\3\32\5\32\u00fa\n\32\3\33\3\33\3\34\6\34\u00ff\n\34"+
		"\r\34\16\34\u0100\3\35\3\35\5\35\u0105\n\35\3\36\3\36\3\37\3\37\3 \3 "+
		"\3!\3!\3\"\3\"\3#\3#\3#\3$\3$\3$\3%\3%\3%\3&\3&\3&\3\'\3\'\3\'\3(\3(\3"+
		"(\3)\3)\3)\3*\3*\3*\3+\3+\3+\3,\3,\3,\3-\3-\3-\3.\3.\3.\3.\3/\3/\3/\3"+
		"/\3/\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3\62"+
		"\3\63\3\63\3\63\3\63\3\64\3\64\3\64\3\64\3\65\3\65\3\65\3\65\3\66\3\66"+
		"\3\66\3\66\3\67\3\67\3\67\3\67\38\38\38\38\39\39\39\39\3:\3:\3:\3:\3;"+
		"\3;\3;\3<\3<\3<\3<\3<\3<\3=\3=\7=\u0173\n=\f=\16=\u0176\13=\3>\3>\3?\3"+
		"?\3@\6@\u017d\n@\r@\16@\u017e\3@\3@\3A\3A\7A\u0185\nA\fA\16A\u0188\13"+
		"A\3A\3A\2\2B\3\3\5\2\7\2\t\2\13\2\r\2\17\2\21\2\23\2\25\2\27\2\31\2\33"+
		"\2\35\2\37\2!\2#\2%\2\'\2)\2+\2-\2/\2\61\2\63\2\65\2\67\29\2;\4=\5?\6"+
		"A\7C\bE\tG\nI\13K\fM\rO\16Q\17S\20U\21W\22Y\23[\24]\25_\26a\27c\30e\31"+
		"g\32i\33k\34m\35o\36q\37s u!w\"y#{\2}\2\177$\u0081%\3\2\r\4\2NNnn\3\2"+
		"\63;\4\2ZZzz\5\2\62;CHch\3\2\629\4\2DDdd\3\2\62\63\6\2//C\\aac|\7\2//"+
		"\62;C\\aac|\5\2\13\f\16\17\"\"\4\2\f\f\17\17\2\u018d\2\3\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2"+
		"\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2"+
		"U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3"+
		"\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2"+
		"\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2"+
		"\177\3\2\2\2\2\u0081\3\2\2\2\3\u0087\3\2\2\2\5\u0089\3\2\2\2\7\u008d\3"+
		"\2\2\2\t\u0091\3\2\2\2\13\u0095\3\2\2\2\r\u0099\3\2\2\2\17\u00a5\3\2\2"+
		"\2\21\u00a7\3\2\2\2\23\u00b0\3\2\2\2\25\u00b2\3\2\2\2\27\u00b5\3\2\2\2"+
		"\31\u00bb\3\2\2\2\33\u00be\3\2\2\2\35\u00c2\3\2\2\2\37\u00c6\3\2\2\2!"+
		"\u00cd\3\2\2\2#\u00d0\3\2\2\2%\u00d6\3\2\2\2\'\u00d8\3\2\2\2)\u00de\3"+
		"\2\2\2+\u00e5\3\2\2\2-\u00e8\3\2\2\2/\u00ee\3\2\2\2\61\u00f0\3\2\2\2\63"+
		"\u00f4\3\2\2\2\65\u00fb\3\2\2\2\67\u00fe\3\2\2\29\u0104\3\2\2\2;\u0106"+
		"\3\2\2\2=\u0108\3\2\2\2?\u010a\3\2\2\2A\u010c\3\2\2\2C\u010e\3\2\2\2E"+
		"\u0110\3\2\2\2G\u0113\3\2\2\2I\u0116\3\2\2\2K\u0119\3\2\2\2M\u011c\3\2"+
		"\2\2O\u011f\3\2\2\2Q\u0122\3\2\2\2S\u0125\3\2\2\2U\u0128\3\2\2\2W\u012b"+
		"\3\2\2\2Y\u012e\3\2\2\2[\u0131\3\2\2\2]\u0135\3\2\2\2_\u013a\3\2\2\2a"+
		"\u013f\3\2\2\2c\u0143\3\2\2\2e\u0147\3\2\2\2g\u014b\3\2\2\2i\u014f\3\2"+
		"\2\2k\u0153\3\2\2\2m\u0157\3\2\2\2o\u015b\3\2\2\2q\u015f\3\2\2\2s\u0163"+
		"\3\2\2\2u\u0167\3\2\2\2w\u016a\3\2\2\2y\u0170\3\2\2\2{\u0177\3\2\2\2}"+
		"\u0179\3\2\2\2\177\u017c\3\2\2\2\u0081\u0182\3\2\2\2\u0083\u0088\5\5\3"+
		"\2\u0084\u0088\5\7\4\2\u0085\u0088\5\t\5\2\u0086\u0088\5\13\6\2\u0087"+
		"\u0083\3\2\2\2\u0087\u0084\3\2\2\2\u0087\u0085\3\2\2\2\u0087\u0086\3\2"+
		"\2\2\u0088\4\3\2\2\2\u0089\u008b\5\17\b\2\u008a\u008c\5\r\7\2\u008b\u008a"+
		"\3\2\2\2\u008b\u008c\3\2\2\2\u008c\6\3\2\2\2\u008d\u008f\5\35\17\2\u008e"+
		"\u0090\5\r\7\2\u008f\u008e\3\2\2\2\u008f\u0090\3\2\2\2\u0090\b\3\2\2\2"+
		"\u0091\u0093\5\'\24\2\u0092\u0094\5\r\7\2\u0093\u0092\3\2\2\2\u0093\u0094"+
		"\3\2\2\2\u0094\n\3\2\2\2\u0095\u0097\5\61\31\2\u0096\u0098\5\r\7\2\u0097"+
		"\u0096\3\2\2\2\u0097\u0098\3\2\2\2\u0098\f\3\2\2\2\u0099\u009a\t\2\2\2"+
		"\u009a\16\3\2\2\2\u009b\u00a6\7\62\2\2\u009c\u00a3\5\25\13\2\u009d\u009f"+
		"\5\21\t\2\u009e\u009d\3\2\2\2\u009e\u009f\3\2\2\2\u009f\u00a4\3\2\2\2"+
		"\u00a0\u00a1\5\33\16\2\u00a1\u00a2\5\21\t\2\u00a2\u00a4\3\2\2\2\u00a3"+
		"\u009e\3\2\2\2\u00a3\u00a0\3\2\2\2\u00a4\u00a6\3\2\2\2\u00a5\u009b\3\2"+
		"\2\2\u00a5\u009c\3\2\2\2\u00a6\20\3\2\2\2\u00a7\u00ac\5\23\n\2\u00a8\u00aa"+
		"\5\27\f\2\u00a9\u00a8\3\2\2\2\u00a9\u00aa\3\2\2\2\u00aa\u00ab\3\2\2\2"+
		"\u00ab\u00ad\5\23\n\2\u00ac\u00a9\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\22"+
		"\3\2\2\2\u00ae\u00b1\7\62\2\2\u00af\u00b1\5\25\13\2\u00b0\u00ae\3\2\2"+
		"\2\u00b0\u00af\3\2\2\2\u00b1\24\3\2\2\2\u00b2\u00b3\t\3\2\2\u00b3\26\3"+
		"\2\2\2\u00b4\u00b6\5\31\r\2\u00b5\u00b4\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7"+
		"\u00b5\3\2\2\2\u00b7\u00b8\3\2\2\2\u00b8\30\3\2\2\2\u00b9\u00bc\5\23\n"+
		"\2\u00ba\u00bc\7a\2\2\u00bb\u00b9\3\2\2\2\u00bb\u00ba\3\2\2\2\u00bc\32"+
		"\3\2\2\2\u00bd\u00bf\7a\2\2\u00be\u00bd\3\2\2\2\u00bf\u00c0\3\2\2\2\u00c0"+
		"\u00be\3\2\2\2\u00c0\u00c1\3\2\2\2\u00c1\34\3\2\2\2\u00c2\u00c3\7\62\2"+
		"\2\u00c3\u00c4\t\4\2\2\u00c4\u00c5\5\37\20\2\u00c5\36\3\2\2\2\u00c6\u00cb"+
		"\5!\21\2\u00c7\u00c9\5#\22\2\u00c8\u00c7\3\2\2\2\u00c8\u00c9\3\2\2\2\u00c9"+
		"\u00ca\3\2\2\2\u00ca\u00cc\5!\21\2\u00cb\u00c8\3\2\2\2\u00cb\u00cc\3\2"+
		"\2\2\u00cc \3\2\2\2\u00cd\u00ce\t\5\2\2\u00ce\"\3\2\2\2\u00cf\u00d1\5"+
		"%\23\2\u00d0\u00cf\3\2\2\2\u00d1\u00d2\3\2\2\2\u00d2\u00d0\3\2\2\2\u00d2"+
		"\u00d3\3\2\2\2\u00d3$\3\2\2\2\u00d4\u00d7\5!\21\2\u00d5\u00d7\7a\2\2\u00d6"+
		"\u00d4\3\2\2\2\u00d6\u00d5\3\2\2\2\u00d7&\3\2\2\2\u00d8\u00da\7\62\2\2"+
		"\u00d9\u00db\5\33\16\2\u00da\u00d9\3\2\2\2\u00da\u00db\3\2\2\2\u00db\u00dc"+
		"\3\2\2\2\u00dc\u00dd\5)\25\2\u00dd(\3\2\2\2\u00de\u00e3\5+\26\2\u00df"+
		"\u00e1\5-\27\2\u00e0\u00df\3\2\2\2\u00e0\u00e1\3\2\2\2\u00e1\u00e2\3\2"+
		"\2\2\u00e2\u00e4\5+\26\2\u00e3\u00e0\3\2\2\2\u00e3\u00e4\3\2\2\2\u00e4"+
		"*\3\2\2\2\u00e5\u00e6\t\6\2\2\u00e6,\3\2\2\2\u00e7\u00e9\5/\30\2\u00e8"+
		"\u00e7\3\2\2\2\u00e9\u00ea\3\2\2\2\u00ea\u00e8\3\2\2\2\u00ea\u00eb\3\2"+
		"\2\2\u00eb.\3\2\2\2\u00ec\u00ef\5+\26\2\u00ed\u00ef\7a\2\2\u00ee\u00ec"+
		"\3\2\2\2\u00ee\u00ed\3\2\2\2\u00ef\60\3\2\2\2\u00f0\u00f1\7\62\2\2\u00f1"+
		"\u00f2\t\7\2\2\u00f2\u00f3\5\63\32\2\u00f3\62\3\2\2\2\u00f4\u00f9\5\65"+
		"\33\2\u00f5\u00f7\5\67\34\2\u00f6\u00f5\3\2\2\2\u00f6\u00f7\3\2\2\2\u00f7"+
		"\u00f8\3\2\2\2\u00f8\u00fa\5\65\33\2\u00f9\u00f6\3\2\2\2\u00f9\u00fa\3"+
		"\2\2\2\u00fa\64\3\2\2\2\u00fb\u00fc\t\b\2\2\u00fc\66\3\2\2\2\u00fd\u00ff"+
		"\59\35\2\u00fe\u00fd\3\2\2\2\u00ff\u0100\3\2\2\2\u0100\u00fe\3\2\2\2\u0100"+
		"\u0101\3\2\2\2\u01018\3\2\2\2\u0102\u0105\5\65\33\2\u0103\u0105\7a\2\2"+
		"\u0104\u0102\3\2\2\2\u0104\u0103\3\2\2\2\u0105:\3\2\2\2\u0106\u0107\7"+
		"]\2\2\u0107<\3\2\2\2\u0108\u0109\7_\2\2\u0109>\3\2\2\2\u010a\u010b\7="+
		"\2\2\u010b@\3\2\2\2\u010c\u010d\7<\2\2\u010dB\3\2\2\2\u010e\u010f\7.\2"+
		"\2\u010fD\3\2\2\2\u0110\u0111\7t\2\2\u0111\u0112\7|\2\2\u0112F\3\2\2\2"+
		"\u0113\u0114\7t\2\2\u0114\u0115\7\62\2\2\u0115H\3\2\2\2\u0116\u0117\7"+
		"t\2\2\u0117\u0118\7\63\2\2\u0118J\3\2\2\2\u0119\u011a\7t\2\2\u011a\u011b"+
		"\7\64\2\2\u011bL\3\2\2\2\u011c\u011d\7t\2\2\u011d\u011e\7\65\2\2\u011e"+
		"N\3\2\2\2\u011f\u0120\7t\2\2\u0120\u0121\7\66\2\2\u0121P\3\2\2\2\u0122"+
		"\u0123\7h\2\2\u0123\u0124\7n\2\2\u0124R\3\2\2\2\u0125\u0126\7t\2\2\u0126"+
		"\u0127\7\67\2\2\u0127T\3\2\2\2\u0128\u0129\7r\2\2\u0129\u012a\7e\2\2\u012a"+
		"V\3\2\2\2\u012b\u012c\7t\2\2\u012c\u012d\79\2\2\u012dX\3\2\2\2\u012e\u012f"+
		"\7g\2\2\u012f\u0130\7s\2\2\u0130Z\3\2\2\2\u0131\u0132\7o\2\2\u0132\u0133"+
		"\7q\2\2\u0133\u0134\7x\2\2\u0134\\\3\2\2\2\u0135\u0136\7o\2\2\u0136\u0137"+
		"\7q\2\2\u0137\u0138\7x\2\2\u0138\u0139\7n\2\2\u0139^\3\2\2\2\u013a\u013b"+
		"\7u\2\2\u013b\u013c\7g\2\2\u013c\u013d\7v\2\2\u013d\u013e\7j\2\2\u013e"+
		"`\3\2\2\2\u013f\u0140\7u\2\2\u0140\u0141\7v\2\2\u0141\u0142\7t\2\2\u0142"+
		"b\3\2\2\2\u0143\u0144\7n\2\2\u0144\u0145\7f\2\2\u0145\u0146\7t\2\2\u0146"+
		"d\3\2\2\2\u0147\u0148\7c\2\2\u0148\u0149\7f\2\2\u0149\u014a\7f\2\2\u014a"+
		"f\3\2\2\2\u014b\u014c\7u\2\2\u014c\u014d\7w\2\2\u014d\u014e\7d\2\2\u014e"+
		"h\3\2\2\2\u014f\u0150\7c\2\2\u0150\u0151\7p\2\2\u0151\u0152\7f\2\2\u0152"+
		"j\3\2\2\2\u0153\u0154\7q\2\2\u0154\u0155\7t\2\2\u0155\u0156\7t\2\2\u0156"+
		"l\3\2\2\2\u0157\u0158\7l\2\2\u0158\u0159\7r\2\2\u0159\u015a\7t\2\2\u015a"+
		"n\3\2\2\2\u015b\u015c\7e\2\2\u015c\u015d\7o\2\2\u015d\u015e\7r\2\2\u015e"+
		"p\3\2\2\2\u015f\u0160\7p\2\2\u0160\u0161\7q\2\2\u0161\u0162\7r\2\2\u0162"+
		"r\3\2\2\2\u0163\u0164\7l\2\2\u0164\u0165\7r\2\2\u0165\u0166\7o\2\2\u0166"+
		"t\3\2\2\2\u0167\u0168\7l\2\2\u0168\u0169\7r\2\2\u0169v\3\2\2\2\u016a\u016b"+
		"\7\60\2\2\u016b\u016c\7y\2\2\u016c\u016d\7q\2\2\u016d\u016e\7t\2\2\u016e"+
		"\u016f\7f\2\2\u016fx\3\2\2\2\u0170\u0174\5{>\2\u0171\u0173\5}?\2\u0172"+
		"\u0171\3\2\2\2\u0173\u0176\3\2\2\2\u0174\u0172\3\2\2\2\u0174\u0175\3\2"+
		"\2\2\u0175z\3\2\2\2\u0176\u0174\3\2\2\2\u0177\u0178\t\t\2\2\u0178|\3\2"+
		"\2\2\u0179\u017a\t\n\2\2\u017a~\3\2\2\2\u017b\u017d\t\13\2\2\u017c\u017b"+
		"\3\2\2\2\u017d\u017e\3\2\2\2\u017e\u017c\3\2\2\2\u017e\u017f\3\2\2\2\u017f"+
		"\u0180\3\2\2\2\u0180\u0181\b@\2\2\u0181\u0080\3\2\2\2\u0182\u0186\7=\2"+
		"\2\u0183\u0185\n\f\2\2\u0184\u0183\3\2\2\2\u0185\u0188\3\2\2\2\u0186\u0184"+
		"\3\2\2\2\u0186\u0187\3\2\2\2\u0187\u0189\3\2\2\2\u0188\u0186\3\2\2\2\u0189"+
		"\u018a\bA\3\2\u018a\u0082\3\2\2\2!\2\u0087\u008b\u008f\u0093\u0097\u009e"+
		"\u00a3\u00a5\u00a9\u00ac\u00b0\u00b7\u00bb\u00c0\u00c8\u00cb\u00d2\u00d6"+
		"\u00da\u00e0\u00e3\u00ea\u00ee\u00f6\u00f9\u0100\u0104\u0174\u017e\u0186"+
		"\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}