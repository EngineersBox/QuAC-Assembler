// Code generated from QuACLexer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr4

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type QuACLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var quaclexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func quaclexerLexerInit() {
	staticData := &quaclexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "'['", "']'", "';'", "':'", "','", "'\\n'", "'rz'", "'r0'",
		"'r1'", "'r2'", "'r3'", "'r4'", "'fl'", "'r5'", "'pc'", "'r7'", "",
		"", "", "", "", "", "", "", "", "", "", "'nop'", "", "", "'.word'",
		"'eq'",
	}
	staticData.symbolicNames = []string{
		"", "IntegerLiteral", "LBRACK", "RBRACK", "SEMI", "COLON", "COMMA",
		"NEWLINE", "RZ", "R0", "R1", "R2", "R3", "R4", "FL", "R5", "PC", "R7",
		"MOV", "MOVL", "SETH", "STR", "LDR", "ADD", "SUB", "AND", "ORR", "JPR",
		"CMP", "NOP", "JPM", "JP", "WORD", "EQ", "Identifier", "WS", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"IntegerLiteral", "DecimalIntegerLiteral", "HexIntegerLiteral", "OctalIntegerLiteral",
		"BinaryIntegerLiteral", "IntegerTypeSuffix", "DecimalNumeral", "Digits",
		"Digit", "NonZeroDigit", "DigitsAndUnderscores", "DigitOrUnderscore",
		"Underscores", "HexNumeral", "HexDigits", "HexDigit", "HexDigitsAndUnderscores",
		"HexDigitOrUnderscore", "OctalNumeral", "OctalDigits", "OctalDigit",
		"OctalDigitsAndUnderscores", "OctalDigitOrUnderscore", "BinaryNumeral",
		"BinaryDigits", "BinaryDigit", "BinaryDigitsAndUnderscores", "BinaryDigitOrUnderscore",
		"LBRACK", "RBRACK", "SEMI", "COLON", "COMMA", "NEWLINE", "RZ", "R0",
		"R1", "R2", "R3", "R4", "FL", "R5", "PC", "R7", "MOV", "MOVL", "SETH",
		"STR", "LDR", "ADD", "SUB", "AND", "ORR", "JPR", "CMP", "NOP", "JPM",
		"JP", "WORD", "EQ", "Identifier", "Letter", "LetterOrDigit", "WS", "LINE_COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 36, 436, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 136, 8, 0,
		1, 1, 1, 1, 3, 1, 140, 8, 1, 1, 2, 1, 2, 3, 2, 144, 8, 2, 1, 3, 1, 3, 3,
		3, 148, 8, 3, 1, 4, 1, 4, 3, 4, 152, 8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		3, 6, 159, 8, 6, 1, 6, 1, 6, 1, 6, 3, 6, 164, 8, 6, 3, 6, 166, 8, 6, 1,
		7, 1, 7, 3, 7, 170, 8, 7, 1, 7, 3, 7, 173, 8, 7, 1, 8, 1, 8, 3, 8, 177,
		8, 8, 1, 9, 1, 9, 1, 10, 4, 10, 182, 8, 10, 11, 10, 12, 10, 183, 1, 11,
		1, 11, 3, 11, 188, 8, 11, 1, 12, 4, 12, 191, 8, 12, 11, 12, 12, 12, 192,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 3, 14, 201, 8, 14, 1, 14, 3,
		14, 204, 8, 14, 1, 15, 1, 15, 1, 16, 4, 16, 209, 8, 16, 11, 16, 12, 16,
		210, 1, 17, 1, 17, 3, 17, 215, 8, 17, 1, 18, 1, 18, 3, 18, 219, 8, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 3, 19, 225, 8, 19, 1, 19, 3, 19, 228, 8, 19,
		1, 20, 1, 20, 1, 21, 4, 21, 233, 8, 21, 11, 21, 12, 21, 234, 1, 22, 1,
		22, 3, 22, 239, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 3, 24,
		247, 8, 24, 1, 24, 3, 24, 250, 8, 24, 1, 25, 1, 25, 1, 26, 4, 26, 255,
		8, 26, 11, 26, 12, 26, 256, 1, 27, 1, 27, 3, 27, 261, 8, 27, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40,
		1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 44, 3, 44, 310, 8, 44, 1, 45, 1, 45, 1, 45, 1, 45,
		1, 45, 1, 45, 3, 45, 318, 8, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1,
		46, 3, 46, 326, 8, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 333, 8,
		47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 3, 48, 340, 8, 48, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 3, 49, 347, 8, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		50, 3, 50, 354, 8, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 361, 8,
		51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52, 368, 8, 52, 1, 53, 1, 53,
		1, 53, 1, 53, 1, 53, 3, 53, 375, 8, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 3, 54, 382, 8, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 56, 3, 56, 393, 8, 56, 1, 57, 1, 57, 1, 57, 1, 57, 3, 57, 399,
		8, 57, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1,
		60, 1, 60, 5, 60, 412, 8, 60, 10, 60, 12, 60, 415, 9, 60, 1, 61, 1, 61,
		1, 62, 1, 62, 1, 63, 4, 63, 422, 8, 63, 11, 63, 12, 63, 423, 1, 63, 1,
		63, 1, 64, 1, 64, 5, 64, 430, 8, 64, 10, 64, 12, 64, 433, 9, 64, 1, 64,
		1, 64, 0, 0, 65, 1, 1, 3, 0, 5, 0, 7, 0, 9, 0, 11, 0, 13, 0, 15, 0, 17,
		0, 19, 0, 21, 0, 23, 0, 25, 0, 27, 0, 29, 0, 31, 0, 33, 0, 35, 0, 37, 0,
		39, 0, 41, 0, 43, 0, 45, 0, 47, 0, 49, 0, 51, 0, 53, 0, 55, 0, 57, 2, 59,
		3, 61, 4, 63, 5, 65, 6, 67, 7, 69, 8, 71, 9, 73, 10, 75, 11, 77, 12, 79,
		13, 81, 14, 83, 15, 85, 16, 87, 17, 89, 18, 91, 19, 93, 20, 95, 21, 97,
		22, 99, 23, 101, 24, 103, 25, 105, 26, 107, 27, 109, 28, 111, 29, 113,
		30, 115, 31, 117, 32, 119, 33, 121, 34, 123, 0, 125, 0, 127, 35, 129, 36,
		1, 0, 11, 2, 0, 76, 76, 108, 108, 1, 0, 49, 57, 2, 0, 88, 88, 120, 120,
		3, 0, 48, 57, 65, 70, 97, 102, 1, 0, 48, 55, 2, 0, 66, 66, 98, 98, 1, 0,
		48, 49, 4, 0, 45, 45, 65, 90, 95, 95, 97, 122, 5, 0, 45, 45, 48, 57, 65,
		90, 95, 95, 97, 122, 3, 0, 9, 10, 12, 13, 32, 32, 2, 0, 10, 10, 13, 13,
		451, 0, 1, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69,
		1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0,
		77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0,
		0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0,
		0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0,
		0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107,
		1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0,
		0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1,
		0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 1, 135, 1, 0, 0, 0, 3,
		137, 1, 0, 0, 0, 5, 141, 1, 0, 0, 0, 7, 145, 1, 0, 0, 0, 9, 149, 1, 0,
		0, 0, 11, 153, 1, 0, 0, 0, 13, 165, 1, 0, 0, 0, 15, 167, 1, 0, 0, 0, 17,
		176, 1, 0, 0, 0, 19, 178, 1, 0, 0, 0, 21, 181, 1, 0, 0, 0, 23, 187, 1,
		0, 0, 0, 25, 190, 1, 0, 0, 0, 27, 194, 1, 0, 0, 0, 29, 198, 1, 0, 0, 0,
		31, 205, 1, 0, 0, 0, 33, 208, 1, 0, 0, 0, 35, 214, 1, 0, 0, 0, 37, 216,
		1, 0, 0, 0, 39, 222, 1, 0, 0, 0, 41, 229, 1, 0, 0, 0, 43, 232, 1, 0, 0,
		0, 45, 238, 1, 0, 0, 0, 47, 240, 1, 0, 0, 0, 49, 244, 1, 0, 0, 0, 51, 251,
		1, 0, 0, 0, 53, 254, 1, 0, 0, 0, 55, 260, 1, 0, 0, 0, 57, 262, 1, 0, 0,
		0, 59, 264, 1, 0, 0, 0, 61, 266, 1, 0, 0, 0, 63, 268, 1, 0, 0, 0, 65, 270,
		1, 0, 0, 0, 67, 272, 1, 0, 0, 0, 69, 274, 1, 0, 0, 0, 71, 277, 1, 0, 0,
		0, 73, 280, 1, 0, 0, 0, 75, 283, 1, 0, 0, 0, 77, 286, 1, 0, 0, 0, 79, 289,
		1, 0, 0, 0, 81, 292, 1, 0, 0, 0, 83, 295, 1, 0, 0, 0, 85, 298, 1, 0, 0,
		0, 87, 301, 1, 0, 0, 0, 89, 304, 1, 0, 0, 0, 91, 311, 1, 0, 0, 0, 93, 319,
		1, 0, 0, 0, 95, 327, 1, 0, 0, 0, 97, 334, 1, 0, 0, 0, 99, 341, 1, 0, 0,
		0, 101, 348, 1, 0, 0, 0, 103, 355, 1, 0, 0, 0, 105, 362, 1, 0, 0, 0, 107,
		369, 1, 0, 0, 0, 109, 376, 1, 0, 0, 0, 111, 383, 1, 0, 0, 0, 113, 387,
		1, 0, 0, 0, 115, 394, 1, 0, 0, 0, 117, 400, 1, 0, 0, 0, 119, 406, 1, 0,
		0, 0, 121, 409, 1, 0, 0, 0, 123, 416, 1, 0, 0, 0, 125, 418, 1, 0, 0, 0,
		127, 421, 1, 0, 0, 0, 129, 427, 1, 0, 0, 0, 131, 136, 3, 3, 1, 0, 132,
		136, 3, 5, 2, 0, 133, 136, 3, 7, 3, 0, 134, 136, 3, 9, 4, 0, 135, 131,
		1, 0, 0, 0, 135, 132, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 134, 1, 0,
		0, 0, 136, 2, 1, 0, 0, 0, 137, 139, 3, 13, 6, 0, 138, 140, 3, 11, 5, 0,
		139, 138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 4, 1, 0, 0, 0, 141, 143,
		3, 27, 13, 0, 142, 144, 3, 11, 5, 0, 143, 142, 1, 0, 0, 0, 143, 144, 1,
		0, 0, 0, 144, 6, 1, 0, 0, 0, 145, 147, 3, 37, 18, 0, 146, 148, 3, 11, 5,
		0, 147, 146, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 8, 1, 0, 0, 0, 149,
		151, 3, 47, 23, 0, 150, 152, 3, 11, 5, 0, 151, 150, 1, 0, 0, 0, 151, 152,
		1, 0, 0, 0, 152, 10, 1, 0, 0, 0, 153, 154, 7, 0, 0, 0, 154, 12, 1, 0, 0,
		0, 155, 166, 5, 48, 0, 0, 156, 163, 3, 19, 9, 0, 157, 159, 3, 15, 7, 0,
		158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 164, 1, 0, 0, 0, 160,
		161, 3, 25, 12, 0, 161, 162, 3, 15, 7, 0, 162, 164, 1, 0, 0, 0, 163, 158,
		1, 0, 0, 0, 163, 160, 1, 0, 0, 0, 164, 166, 1, 0, 0, 0, 165, 155, 1, 0,
		0, 0, 165, 156, 1, 0, 0, 0, 166, 14, 1, 0, 0, 0, 167, 172, 3, 17, 8, 0,
		168, 170, 3, 21, 10, 0, 169, 168, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170,
		171, 1, 0, 0, 0, 171, 173, 3, 17, 8, 0, 172, 169, 1, 0, 0, 0, 172, 173,
		1, 0, 0, 0, 173, 16, 1, 0, 0, 0, 174, 177, 5, 48, 0, 0, 175, 177, 3, 19,
		9, 0, 176, 174, 1, 0, 0, 0, 176, 175, 1, 0, 0, 0, 177, 18, 1, 0, 0, 0,
		178, 179, 7, 1, 0, 0, 179, 20, 1, 0, 0, 0, 180, 182, 3, 23, 11, 0, 181,
		180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 184, 22, 1, 0, 0, 0, 185, 188, 3, 17, 8, 0, 186, 188, 5, 95,
		0, 0, 187, 185, 1, 0, 0, 0, 187, 186, 1, 0, 0, 0, 188, 24, 1, 0, 0, 0,
		189, 191, 5, 95, 0, 0, 190, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192,
		190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 26, 1, 0, 0, 0, 194, 195, 5,
		48, 0, 0, 195, 196, 7, 2, 0, 0, 196, 197, 3, 29, 14, 0, 197, 28, 1, 0,
		0, 0, 198, 203, 3, 31, 15, 0, 199, 201, 3, 33, 16, 0, 200, 199, 1, 0, 0,
		0, 200, 201, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 204, 3, 31, 15, 0,
		203, 200, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 30, 1, 0, 0, 0, 205, 206,
		7, 3, 0, 0, 206, 32, 1, 0, 0, 0, 207, 209, 3, 35, 17, 0, 208, 207, 1, 0,
		0, 0, 209, 210, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0,
		211, 34, 1, 0, 0, 0, 212, 215, 3, 31, 15, 0, 213, 215, 5, 95, 0, 0, 214,
		212, 1, 0, 0, 0, 214, 213, 1, 0, 0, 0, 215, 36, 1, 0, 0, 0, 216, 218, 5,
		48, 0, 0, 217, 219, 3, 25, 12, 0, 218, 217, 1, 0, 0, 0, 218, 219, 1, 0,
		0, 0, 219, 220, 1, 0, 0, 0, 220, 221, 3, 39, 19, 0, 221, 38, 1, 0, 0, 0,
		222, 227, 3, 41, 20, 0, 223, 225, 3, 43, 21, 0, 224, 223, 1, 0, 0, 0, 224,
		225, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 228, 3, 41, 20, 0, 227, 224,
		1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 40, 1, 0, 0, 0, 229, 230, 7, 4,
		0, 0, 230, 42, 1, 0, 0, 0, 231, 233, 3, 45, 22, 0, 232, 231, 1, 0, 0, 0,
		233, 234, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235,
		44, 1, 0, 0, 0, 236, 239, 3, 41, 20, 0, 237, 239, 5, 95, 0, 0, 238, 236,
		1, 0, 0, 0, 238, 237, 1, 0, 0, 0, 239, 46, 1, 0, 0, 0, 240, 241, 5, 48,
		0, 0, 241, 242, 7, 5, 0, 0, 242, 243, 3, 49, 24, 0, 243, 48, 1, 0, 0, 0,
		244, 249, 3, 51, 25, 0, 245, 247, 3, 53, 26, 0, 246, 245, 1, 0, 0, 0, 246,
		247, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 3, 51, 25, 0, 249, 246,
		1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 50, 1, 0, 0, 0, 251, 252, 7, 6,
		0, 0, 252, 52, 1, 0, 0, 0, 253, 255, 3, 55, 27, 0, 254, 253, 1, 0, 0, 0,
		255, 256, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257,
		54, 1, 0, 0, 0, 258, 261, 3, 51, 25, 0, 259, 261, 5, 95, 0, 0, 260, 258,
		1, 0, 0, 0, 260, 259, 1, 0, 0, 0, 261, 56, 1, 0, 0, 0, 262, 263, 5, 91,
		0, 0, 263, 58, 1, 0, 0, 0, 264, 265, 5, 93, 0, 0, 265, 60, 1, 0, 0, 0,
		266, 267, 5, 59, 0, 0, 267, 62, 1, 0, 0, 0, 268, 269, 5, 58, 0, 0, 269,
		64, 1, 0, 0, 0, 270, 271, 5, 44, 0, 0, 271, 66, 1, 0, 0, 0, 272, 273, 5,
		10, 0, 0, 273, 68, 1, 0, 0, 0, 274, 275, 5, 114, 0, 0, 275, 276, 5, 122,
		0, 0, 276, 70, 1, 0, 0, 0, 277, 278, 5, 114, 0, 0, 278, 279, 5, 48, 0,
		0, 279, 72, 1, 0, 0, 0, 280, 281, 5, 114, 0, 0, 281, 282, 5, 49, 0, 0,
		282, 74, 1, 0, 0, 0, 283, 284, 5, 114, 0, 0, 284, 285, 5, 50, 0, 0, 285,
		76, 1, 0, 0, 0, 286, 287, 5, 114, 0, 0, 287, 288, 5, 51, 0, 0, 288, 78,
		1, 0, 0, 0, 289, 290, 5, 114, 0, 0, 290, 291, 5, 52, 0, 0, 291, 80, 1,
		0, 0, 0, 292, 293, 5, 102, 0, 0, 293, 294, 5, 108, 0, 0, 294, 82, 1, 0,
		0, 0, 295, 296, 5, 114, 0, 0, 296, 297, 5, 53, 0, 0, 297, 84, 1, 0, 0,
		0, 298, 299, 5, 112, 0, 0, 299, 300, 5, 99, 0, 0, 300, 86, 1, 0, 0, 0,
		301, 302, 5, 114, 0, 0, 302, 303, 5, 55, 0, 0, 303, 88, 1, 0, 0, 0, 304,
		305, 5, 109, 0, 0, 305, 306, 5, 111, 0, 0, 306, 307, 5, 118, 0, 0, 307,
		309, 1, 0, 0, 0, 308, 310, 3, 119, 59, 0, 309, 308, 1, 0, 0, 0, 309, 310,
		1, 0, 0, 0, 310, 90, 1, 0, 0, 0, 311, 312, 5, 109, 0, 0, 312, 313, 5, 111,
		0, 0, 313, 314, 5, 118, 0, 0, 314, 315, 5, 108, 0, 0, 315, 317, 1, 0, 0,
		0, 316, 318, 3, 119, 59, 0, 317, 316, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0,
		318, 92, 1, 0, 0, 0, 319, 320, 5, 115, 0, 0, 320, 321, 5, 101, 0, 0, 321,
		322, 5, 116, 0, 0, 322, 323, 5, 104, 0, 0, 323, 325, 1, 0, 0, 0, 324, 326,
		3, 119, 59, 0, 325, 324, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 94, 1,
		0, 0, 0, 327, 328, 5, 115, 0, 0, 328, 329, 5, 116, 0, 0, 329, 330, 5, 114,
		0, 0, 330, 332, 1, 0, 0, 0, 331, 333, 3, 119, 59, 0, 332, 331, 1, 0, 0,
		0, 332, 333, 1, 0, 0, 0, 333, 96, 1, 0, 0, 0, 334, 335, 5, 108, 0, 0, 335,
		336, 5, 100, 0, 0, 336, 337, 5, 114, 0, 0, 337, 339, 1, 0, 0, 0, 338, 340,
		3, 119, 59, 0, 339, 338, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 98, 1,
		0, 0, 0, 341, 342, 5, 97, 0, 0, 342, 343, 5, 100, 0, 0, 343, 344, 5, 100,
		0, 0, 344, 346, 1, 0, 0, 0, 345, 347, 3, 119, 59, 0, 346, 345, 1, 0, 0,
		0, 346, 347, 1, 0, 0, 0, 347, 100, 1, 0, 0, 0, 348, 349, 5, 115, 0, 0,
		349, 350, 5, 117, 0, 0, 350, 351, 5, 98, 0, 0, 351, 353, 1, 0, 0, 0, 352,
		354, 3, 119, 59, 0, 353, 352, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 102,
		1, 0, 0, 0, 355, 356, 5, 97, 0, 0, 356, 357, 5, 110, 0, 0, 357, 358, 5,
		100, 0, 0, 358, 360, 1, 0, 0, 0, 359, 361, 3, 119, 59, 0, 360, 359, 1,
		0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 104, 1, 0, 0, 0, 362, 363, 5, 111,
		0, 0, 363, 364, 5, 114, 0, 0, 364, 365, 5, 114, 0, 0, 365, 367, 1, 0, 0,
		0, 366, 368, 3, 119, 59, 0, 367, 366, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0,
		368, 106, 1, 0, 0, 0, 369, 370, 5, 106, 0, 0, 370, 371, 5, 112, 0, 0, 371,
		372, 5, 114, 0, 0, 372, 374, 1, 0, 0, 0, 373, 375, 3, 119, 59, 0, 374,
		373, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 108, 1, 0, 0, 0, 376, 377,
		5, 99, 0, 0, 377, 378, 5, 109, 0, 0, 378, 379, 5, 112, 0, 0, 379, 381,
		1, 0, 0, 0, 380, 382, 3, 119, 59, 0, 381, 380, 1, 0, 0, 0, 381, 382, 1,
		0, 0, 0, 382, 110, 1, 0, 0, 0, 383, 384, 5, 110, 0, 0, 384, 385, 5, 111,
		0, 0, 385, 386, 5, 112, 0, 0, 386, 112, 1, 0, 0, 0, 387, 388, 5, 106, 0,
		0, 388, 389, 5, 112, 0, 0, 389, 390, 5, 109, 0, 0, 390, 392, 1, 0, 0, 0,
		391, 393, 3, 119, 59, 0, 392, 391, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393,
		114, 1, 0, 0, 0, 394, 395, 5, 106, 0, 0, 395, 396, 5, 112, 0, 0, 396, 398,
		1, 0, 0, 0, 397, 399, 3, 119, 59, 0, 398, 397, 1, 0, 0, 0, 398, 399, 1,
		0, 0, 0, 399, 116, 1, 0, 0, 0, 400, 401, 5, 46, 0, 0, 401, 402, 5, 119,
		0, 0, 402, 403, 5, 111, 0, 0, 403, 404, 5, 114, 0, 0, 404, 405, 5, 100,
		0, 0, 405, 118, 1, 0, 0, 0, 406, 407, 5, 101, 0, 0, 407, 408, 5, 113, 0,
		0, 408, 120, 1, 0, 0, 0, 409, 413, 3, 123, 61, 0, 410, 412, 3, 125, 62,
		0, 411, 410, 1, 0, 0, 0, 412, 415, 1, 0, 0, 0, 413, 411, 1, 0, 0, 0, 413,
		414, 1, 0, 0, 0, 414, 122, 1, 0, 0, 0, 415, 413, 1, 0, 0, 0, 416, 417,
		7, 7, 0, 0, 417, 124, 1, 0, 0, 0, 418, 419, 7, 8, 0, 0, 419, 126, 1, 0,
		0, 0, 420, 422, 7, 9, 0, 0, 421, 420, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0,
		423, 421, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425,
		426, 6, 63, 0, 0, 426, 128, 1, 0, 0, 0, 427, 431, 5, 59, 0, 0, 428, 430,
		8, 10, 0, 0, 429, 428, 1, 0, 0, 0, 430, 433, 1, 0, 0, 0, 431, 429, 1, 0,
		0, 0, 431, 432, 1, 0, 0, 0, 432, 434, 1, 0, 0, 0, 433, 431, 1, 0, 0, 0,
		434, 435, 6, 64, 1, 0, 435, 130, 1, 0, 0, 0, 44, 0, 135, 139, 143, 147,
		151, 158, 163, 165, 169, 172, 176, 183, 187, 192, 200, 203, 210, 214, 218,
		224, 227, 234, 238, 246, 249, 256, 260, 309, 317, 325, 332, 339, 346, 353,
		360, 367, 374, 381, 392, 398, 413, 423, 431, 2, 6, 0, 0, 0, 1, 0,
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

// QuACLexerInit initializes any static state used to implement QuACLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewQuACLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func QuACLexerInit() {
	staticData := &quaclexerLexerStaticData
	staticData.once.Do(quaclexerLexerInit)
}

// NewQuACLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewQuACLexer(input antlr.CharStream) *QuACLexer {
	QuACLexerInit()
	l := new(QuACLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &quaclexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "QuACLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QuACLexer tokens.
const (
	QuACLexerIntegerLiteral = 1
	QuACLexerLBRACK         = 2
	QuACLexerRBRACK         = 3
	QuACLexerSEMI           = 4
	QuACLexerCOLON          = 5
	QuACLexerCOMMA          = 6
	QuACLexerNEWLINE        = 7
	QuACLexerRZ             = 8
	QuACLexerR0             = 9
	QuACLexerR1             = 10
	QuACLexerR2             = 11
	QuACLexerR3             = 12
	QuACLexerR4             = 13
	QuACLexerFL             = 14
	QuACLexerR5             = 15
	QuACLexerPC             = 16
	QuACLexerR7             = 17
	QuACLexerMOV            = 18
	QuACLexerMOVL           = 19
	QuACLexerSETH           = 20
	QuACLexerSTR            = 21
	QuACLexerLDR            = 22
	QuACLexerADD            = 23
	QuACLexerSUB            = 24
	QuACLexerAND            = 25
	QuACLexerORR            = 26
	QuACLexerJPR            = 27
	QuACLexerCMP            = 28
	QuACLexerNOP            = 29
	QuACLexerJPM            = 30
	QuACLexerJP             = 31
	QuACLexerWORD           = 32
	QuACLexerEQ             = 33
	QuACLexerIdentifier     = 34
	QuACLexerWS             = 35
	QuACLexerLINE_COMMENT   = 36
)
