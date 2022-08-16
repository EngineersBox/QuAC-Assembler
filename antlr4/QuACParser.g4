parser grammar QuACParser;

options {
    tokenVocab = QuACLexer;
}

parse: (statement)* EOF;
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
iFormat: (MOVL | SETH) register COMMA IntegerLiteral;
rMemFormat: (STR | LDR) register COMMA LBRACK register RBRACK;
rALUFormat: (ADD | SUB | AND | ORR) register COMMA register COMMA register;

// Pseudo Insns
nop: NOP;
pseudo2Param: (MOV | CMP) register COMMA register;
jpr: JPR register;
jpm: JPM LBRACK register RBRACK;
jp: JP (IntegerLiteral | Identifier);

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
