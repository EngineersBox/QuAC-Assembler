lexer grammar QuACLexer;

IntegerLiteral:
	DecimalIntegerLiteral
	| HexIntegerLiteral
	| OctalIntegerLiteral
	| BinaryIntegerLiteral;

fragment DecimalIntegerLiteral: DecimalNumeral IntegerTypeSuffix?;
fragment HexIntegerLiteral: HexNumeral IntegerTypeSuffix?;
fragment OctalIntegerLiteral: OctalNumeral IntegerTypeSuffix?;
fragment BinaryIntegerLiteral: BinaryNumeral IntegerTypeSuffix?;
fragment IntegerTypeSuffix: [lL];
fragment DecimalNumeral: '0' | NonZeroDigit (Digits? | Underscores Digits);
fragment Digits: Digit (DigitsAndUnderscores? Digit)?;
fragment Digit: '0' | NonZeroDigit;
fragment NonZeroDigit: [1-9];
fragment DigitsAndUnderscores: DigitOrUnderscore+;
fragment DigitOrUnderscore: Digit | '_';
fragment Underscores: '_'+;
fragment HexNumeral: '0' [xX] HexDigits;
fragment HexDigits: HexDigit (HexDigitsAndUnderscores? HexDigit)?;
fragment HexDigit: [0-9a-fA-F];
fragment HexDigitsAndUnderscores: HexDigitOrUnderscore+;
fragment HexDigitOrUnderscore: HexDigit | '_';
fragment OctalNumeral: '0' Underscores? OctalDigits;
fragment OctalDigits: OctalDigit (OctalDigitsAndUnderscores? OctalDigit)?;
fragment OctalDigit: [0-7];
fragment OctalDigitsAndUnderscores: OctalDigitOrUnderscore+;
fragment OctalDigitOrUnderscore: OctalDigit | '_';
fragment BinaryNumeral: '0' [bB] BinaryDigits;
fragment BinaryDigits: BinaryDigit (BinaryDigitsAndUnderscores? BinaryDigit)?;
fragment BinaryDigit: [01];
fragment BinaryDigitsAndUnderscores: BinaryDigitOrUnderscore+;
fragment BinaryDigitOrUnderscore: BinaryDigit | '_';

// Separators
LBRACK: '[';
RBRACK: ']';
SEMI: ';';
COLON: ':';
COMMA: ',';
NEWLINE: '\n';

// Registers
RZ: 'rz';
R0: 'r0';
R1: 'r1';
R2: 'r2';
R3: 'r3';
R4: 'r4';
FL: 'fl';
R5: 'r5';
PC: 'pc';
R7: 'r7';

// Insns
MOV: 'mov' EQ?;
MOVL: 'movl' EQ?;
SETH: 'seth' EQ?;
STR: 'str' EQ?;
LDR: 'ldr' EQ?;
ADD: 'add' EQ?;
SUB: 'sub' EQ?;
AND: 'and' EQ?;
ORR: 'orr' EQ?;
JPR: 'jpr' EQ?;
CMP: 'cmp' EQ?;
NOP: 'nop';
JPM: 'jpm' EQ?;
JP: 'jp' EQ?;
WORD: '.word';

// Conditions
fragment EQ: 'eq';

// Identifiers (must appear after all keywords in the grammar)
Identifier: Letter LetterOrDigit*;
fragment Letter: [a-zA-Z\-_];
fragment LetterOrDigit: [a-zA-Z0-9\-_];

// Whitespace and comments
WS: [ \t\r\n\u000C]+ -> skip;
LINE_COMMENT: ';' ~[\r\n]* -> channel(HIDDEN);