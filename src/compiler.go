package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/EngineersBox/QuAC-Compiler/antlr4"
	"github.com/EngineersBox/QuAC-Compiler/src/insn"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var commentRegex *regexp.Regexp = regexp.MustCompile(";.*")

func removeComments(line string) string {
	return commentRegex.ReplaceAllString(line, "")
}

func compile(scanner *bufio.Scanner) ([]byte, error) {
	var compiled []byte
	var lineNumber uint64 = 0
	for scanner.Scan() {
		var toParse string = scanner.Text()
		toParse = strings.TrimSpace(removeComments(toParse))

		if len(toParse) == 0 {
			lineNumber++
			continue
		}

		insn, err := insn.ParseInsn(toParse)
		if err != nil {
			return nil, fmt.Errorf("[Line %d] %s\n", lineNumber, err)
		}
		// fmt.Printf("%s => %016b\n", toParse, insn)
		insnBytes := make([]byte, 2)
		binary.BigEndian.PutUint16(insnBytes, insn)

		compiled = append(compiled, insnBytes...)
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return compiled, nil
}

// func main() {
// 	var args []string = os.Args[1:]
// 	if len(args) != 2 {
// 		panic("Usage: compiler <source assembly> <destination binary>")
// 	}
// 	file, err := os.Open(args[0])
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)

// 	compiled, err := compile(scanner)
// 	if err != nil {
// 		panic(err)
// 	}

// 	outFile, err := os.Create(args[1])
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer outFile.Close()

// 	_, err = outFile.Write(compiled)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	var args []string = os.Args[1:]
	if len(args) != 2 {
		panic("Usage: compiler <source assembly> <destination binary>")
	}

	input, _ := antlr.NewFileStream(args[0])
	lexer := antlr4.NewQuACLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := antlr4.NewQuACParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Parse()

	var visitor insn.InsnVisitor = insn.InsnVisitor{}
	fmt.Println("Calling visitor")
	var result []uint16 = visitor.Visit(tree).([]uint16)
	var bytesResult []byte = make([]byte, len(result)*2)
	for _, asmCommand := range result {
		bytesResult = append(bytesResult, byte(asmCommand&0x00FF), byte((asmCommand&0xFF00)>>8))
	}

	outFile, err := os.Create(args[1])
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, err = outFile.Write(bytesResult)
	if err != nil {
		panic(err)
	}
}
