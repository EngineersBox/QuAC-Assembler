parser grammar QuACParser;

options {
    tokenVocab = QuACLexer;
}

parse: statement* EOF;
statement: iFormat # iFormatStatement
    | rMemFormat # rMemFormatStatement
    | rALUFormat # rALUFormatStatement
    | nop # nopStatement
    | pseudo2Param # pseudo2ParamStatement
    | jpr # jprStatement
    | jpm # jpmStatement
    | jp # jpStatement
	| WORD IntegerLiteral # wordStatement
	| Identifier COLON # labelStatement;

// Core Insns
iFormat: (MOVL | SETH) EQ? register COMMA IntegerLiteral;
rMemFormat: (STR | LDR) EQ? register COMMA LBRACK register RBRACK;
rALUFormat: (ADD | SUB | AND | ORR) EQ? register COMMA register COMMA register;

// Pseudo Insns
nop: NOP;
pseudo2Param: (MOV | CMP) EQ? register COMMA register;
jpr: JPR EQ? register;
jpm: JPM EQ? LBRACK register RBRACK;
jp: JP EQ? (IntegerLiteral | Identifier);

// Registers
register: RZ
    | R0
    | R1
    | R2
    | R3
    | R4
    | FL
    | R5
    | PC
    | R7;

